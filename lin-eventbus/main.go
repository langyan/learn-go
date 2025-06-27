package main

import (
	"fmt"
	"time"
	"github.com/langyan/lin-eventbus/eventbus"

)
func main() {
	eventBus := eventbus.NewEventBus()
  
	// Subscribe to the "post" topic event
	subscribe := eventBus.Subscribe("post")
  
	go func() {
	  for event := range subscribe {
		fmt.Println(event.Payload)
	  }
	}()
  
	eventBus.Publish("post", eventbus.Event{Payload: map[string]any{
	  "postId": 1,
	  "title":  "Welcome to Leapcell",
	  "author": "Leapcell",
	}})
	// Topic with no subscribers
	eventBus.Publish("pay", eventbus.Event{Payload: "pay"})
  
	time.Sleep(time.Second * 2)
	// Unsubscribe from the "post" topic event
	eventBus.Unsubscribe("post", subscribe)
  }