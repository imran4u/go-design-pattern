package main

import "fmt"

func main() {
	open := OpenLaptop{}
	ppb := PressPowerButton{}
	ep := EnterPassword{}
	// adding chain
	open.nextHandler(&ppb)
	ppb.nextHandler(&ep)

	// triger the execution
	open.execute()

}

type Handler interface {
	execute()
	nextHandler(h Handler)
}

type OpenLaptop struct {
	next Handler
}

func (ol *OpenLaptop) execute() {
	fmt.Println("Open laptop .....")
	if ol.next != nil {
		ol.next.execute()
	}
}

func (ol *OpenLaptop) nextHandler(h Handler) {
	ol.next = h
}

type PressPowerButton struct {
	next Handler
}

func (ppb *PressPowerButton) execute() {
	fmt.Println("Press PowerButton.....")
	if ppb.next != nil {
		ppb.next.execute()
	}
}

func (ppb *PressPowerButton) nextHandler(h Handler) {
	ppb.next = h
}

type EnterPassword struct {
	next Handler
}

func (ep *EnterPassword) execute() {
	fmt.Println("Enter password.....")
	fmt.Println("Yo ho laptp is open .....")
	if ep.next != nil {
		ep.next.execute()
	}
}

func (ep *EnterPassword) nextHandler(h Handler) {
	ep.next = h
}
