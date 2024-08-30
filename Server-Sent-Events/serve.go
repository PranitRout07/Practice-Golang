package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handleapi() {
	http.HandleFunc("/", health)
	http.HandleFunc("/event", eventHandler)
	http.ListenAndServe(":3000", nil)
}
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Visited endpoint", r.URL.RequestURI())
}
func eventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Visited endpoint", r.URL.RequestURI())

	flusher, _ := w.(http.Flusher)
	doneChan := make(chan bool)
	eventChan := make(chan Event, 1024)

	go func() {
		for {
			eventChan <- Event{
				Attempts: rand.Intn(100),
			}
		}
	}()

	go func() {
		fmt.Println(r.Context().Done())
		<-r.Context().Done()
		fmt.Println("CONTEXT", r.Context())
		doneChan <- true
	}()

	for {
		event := <-eventChan
		select {
		case <-doneChan:
			fmt.Println("client disconnected")
			return
		default:
			jsonbytes, _ := json.Marshal(event)
			// res := string(jsonbytes)
			// fmt.Fprintf(w, "Event data: %s\n\n", res)
			w.Write(jsonbytes)
			
			time.Sleep(2 * time.Second)
			flusher.Flush()

		}

	}
}

type Event struct {
	Attempts int
}
