package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

//Event Struct of a basic event
type Event struct {
	eventID    int
	eventValue string
}

func workers(eventChan chan Event, id int) {
	defer wg.Done()
	for event := range eventChan {
		log.Println("worker id = ", id, event.eventID)
	}
}

func main() {
	events := make(chan Event)

	for i := 0; i < 10; i++ {
		go wg.Add(1)
		go workers(events, i)
	}

	for i := 0; i < 10; i++ {
		event := Event{i, "hell"}
		events <- event
	}
	close(events)
	wg.Wait()
	log.Println("Done")
}
