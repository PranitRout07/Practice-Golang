package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// resp, _ := http.Get("http://example.com/")
	client := &http.Client{
		CheckRedirect: nil,
	}

	// resp, _ := client.Get("http://example.com")
	client.CloseIdleConnections()  //closes previously connections

	req, _ := http.NewRequest("GET", "http://example.com", nil)

	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	// resp,_ := client.Do(req)
	resp,_ := client.Head("http://example.com")
	defer resp.Body.Close()
	log.Println(resp.Header)

	out, _ := io.ReadAll(resp.Body)
	log.Println(string(out))
}
// package main

// import (
// 	"fmt"
// 	"os"
// 	"net/http"
// )

// func memInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	data, err := os.ReadFile("/proc/meminfo")
// 	if err != nil {
// 		http.Error(w, "Could not read /proc/meminfo", http.StatusInternalServerError)
// 		return
// 	}

	
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write(data)
// }

// func main() {
// 	http.HandleFunc("/meminfo", memInfoHandler)
// 	fmt.Println("Server started at http://localhost:8080/meminfo")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }
