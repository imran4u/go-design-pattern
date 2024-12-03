package main

import "log"

func main() {
	p := NewPublisher()

	s1 := Subscriber{id: "1234"}
	s2 := Subscriber{id: "abcd"}
	p.Add(&s1)
	p.Add(&s2)

	p.Publish("Hello s1 & s2")

}

type PublisherInter interface {
	Add(sub SubscriberInter)
	Remove(id string) // id of subscriber
	Publish(smg string)
}

type SubscriberInter interface {
	getId() string
	handleMessage(msg string)
}

// Publisher -> implementation
type Publisher struct {
	registry map[string]SubscriberInter
}

func NewPublisher() PublisherInter {
	return &Publisher{registry: make(map[string]SubscriberInter)}
}

func (p *Publisher) Add(sub SubscriberInter) {
	p.registry[sub.getId()] = sub
}
func (p *Publisher) Remove(id string) {
	delete(p.registry, id)
}

func (p *Publisher) Publish(msg string) {
	for _, s := range p.registry {
		s.handleMessage(msg)
	}
}

//Subscriber -> Implementation

type Subscriber struct {
	id string
}

func (s *Subscriber) getId() string {
	return s.id
}

func (s *Subscriber) handleMessage(msg string) {
	log.Printf("handleMessage on Subscriber = %s , Message = %s", s.getId(), msg)
}
