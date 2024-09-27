package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemTable struct {
	mu    sync.RWMutex
	table map[string]string
}

func NewMemTable() *MemTable {
	return &MemTable{
		table: make(map[string]string),
	}
}

func (m *MemTable) Put(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.table[key] = value
}

func (m *MemTable) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.table[key]
	return val, exists
}

func (m *MemTable) FlushToFile(filename string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Sort keys and prepare data for writing
	keys := make([]string, 0, len(m.table))
	for k := range m.table {
		keys = append(keys, k)
	}
	fmt.Println(keys, "[keys]BEFORE")
	sort.Strings(keys)
	fmt.Println(keys, "[keys]AFTER")
	var kvs []KeyValue
	for _, k := range keys {
		kvs = append(kvs, KeyValue{Key: k, Value: m.table[k]})
	}

	data, err := json.Marshal(kvs)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

type LSMTree struct {
	memTable   *MemTable
	sstables   []string
	flushLimit int
}

func NewLSMTree(flushLimit int) *LSMTree {
	return &LSMTree{
		memTable:   NewMemTable(),
		flushLimit: flushLimit,
	}
}

var tempTable *MemTable


func (l *LSMTree) FlushToSingleFile() error{
	keys := make([]string, 0, len(l.memTable.table))
	for k := range l.memTable.table {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var kvs []KeyValue
	for _, k := range keys {
		kvs = append(kvs, KeyValue{Key: k, Value: l.memTable.table[k]})
	}

	data, err := json.Marshal(kvs)
	if err != nil {
		return err
	}
	
	return os.WriteFile(singleFile, data, 0644)
	
}


func (l *LSMTree) Put(key, value string) {
	l.memTable.Put(key, value)
	if len(l.memTable.table) >= l.flushLimit {
		if global_length == l.flushLimit {
			//handle
			fmt.Println("HELLOOOOOO")
			l.FlushToSingleFile()
			l.sstables = append(l.sstables, singleFile)
			global_length += 1
			l.memTable = NewMemTable()

		} else {
			// fmt.Println("FLUSHHHHH")
			l.Flush()
		} // 0 1 2 3         limit 5  g==5

	} else if global_length < l.flushLimit {

		// Sort keys and prepare data for writing              //HANDLING limit 10000, data entry:10 -> solve 10 data in one file
		//problem:- flush first full when getting full , memtable is not getting cleared
		fmt.Println(global_length, l.flushLimit) //effect:- data duplicate

		//globallength ==5 ? handle, len(memtable)==5 => execute => (1-5)json
		l.FlushToSingleFile()
		global_length += 1
	}

}

func (l *LSMTree) Get(key string) (string, bool) {
	if val, exists := l.memTable.Get(key); exists {
		return val, true
	}

	//check if cur json file?
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("FILES", len(files))
	for _, entry := range files {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {

			l.sstables = append(l.sstables,entry.Name())
		}
	}

	//if present load to program if not skip

	// Check SSTables
	for _, file := range l.sstables {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		var kvs []KeyValue
		if err := json.Unmarshal(data, &kvs); err == nil {
			for _, kv := range kvs {
				if kv.Key == key {
					return kv.Value, true
				}
			}
		}
	}
	return "", false
}

// global length of data
var global_length int = 1

// singlefile
var singleFile = fmt.Sprintf("sstable_%d.json", 0)

func (l *LSMTree) Flush() {

	filename := fmt.Sprintf("sstable_%d.json", len(l.sstables)+1)
	if err := l.memTable.FlushToFile(filename); err == nil {
		l.sstables = append(l.sstables, filename)
		global_length = global_length + len(l.memTable.table)
		l.memTable = NewMemTable() // Reset the MemTable
	}
}

func main() {
	lsm := NewLSMTree(3)

	// Insert some key-value pairs

	now := time.Now()
	for i := 1; i <= 9; i++ {
		if i%10 != 0 {
			lsm.Put(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i))
		}

	}
	if val, exists := lsm.Get("key-9"); exists {
		fmt.Println("Retrieved key2:", val)
	} else {
		fmt.Println("key2 not found")
	}
	// if val, exists := lsm.Get("key-7899"); exists {
	// 	fmt.Println("Retrieved key7899:", val)
	// } else {
	// 	fmt.Println("key99 not found")
	// }

	// if val, exists := lsm.Get("key-9999"); exists {
	// 	fmt.Println("Retrieved key9999:", val)
	// } else {
	// 	fmt.Println("key49 not found")
	// }
	fmt.Println("time taken", time.Since(now))

	// Retrieve values
	// if val, exists := lsm.Get("key-2"); exists {
	// 	fmt.Println("Retrieved key2:", val)
	// } else {
	// 	fmt.Println("key2 not found")
	// }

	// if val, exists := lsm.Get("key-9999"); exists {
	// 	fmt.Println("Retrieved key9999:", val)
	// } else {
	// 	fmt.Println("key49 not found")
	// }

}
