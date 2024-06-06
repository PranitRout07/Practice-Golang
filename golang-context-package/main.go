package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx,cancel := context.WithCancel(ctx)

	go func(){
		time.Sleep(3*time.Second)
		cancel()
	}()
	
	select {
	case <-time.After(2*time.Second):
			fmt.Println("hello")
		
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}

}