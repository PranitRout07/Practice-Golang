package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next func(w http.ResponseWriter,r *http.Request)) http.Handler{
	limiter := rate.NewLimiter(2,4)
	return  http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		if !limiter.Allow(){
			msg := Msg{
				Status: "Request Failed",
				Body: "The API is limit , try later!",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&msg)
			return
		}else{
			next(w,r)
		}
	})
}