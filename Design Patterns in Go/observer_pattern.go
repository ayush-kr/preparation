package main

import (
	"log"
)

//Event is the event to be published
type Event struct {
	val int
}

//Clienter is interface implemented by the clients
type Clienter interface {
	NotifyCallback(Event)
}

//Subjecter is the interface implemented by the Subject
type Subjecter interface {
	AddClient(Clienter)
	Notify(Event)
}

type subject struct {
	val     int
	clients map[Clienter]interface{}
}

type client struct {
	id string
}

//AddClient adds a new client in the subject map
func (s *subject) AddClient(client Clienter) {
	s.clients[client] = struct{}{}
}

//Notify sends the event to all the clients
func (s *subject) Notify(event Event) {
	for key := range s.clients {
		key.(Clienter).NotifyCallback(event)
	}
}

//NotifyCallback receives the callback from the subject about the event
func (c *client) NotifyCallback(event Event) {
	log.Println(event.val)
}

func main() {

	c1 := client{
		id: "client1",
	}

	c2 := client{
		id: "client2",
	}

	s := subject{
		val:     0,
		clients: make(map[Clienter]interface{}),
	}

	s.AddClient(&c1)
	s.AddClient(&c2)

	ev := Event{
		val: 0,
	}
	for i := 0; i < 20; i++ {
		ev.val = i
		s.Notify(ev)
	}
}
