package main

import "fmt"

func main() {
	p := &Proxy{}
	p.Request()

}

// Real subject
type RealSubject struct {
}

func (r *RealSubject) Request() {
	fmt.Println("Requesting on Real object ....")
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() {
	fmt.Println("Requesting on Proxy  object here we can add control ....")
	if p.realSubject == nil {
		p.realSubject = &RealSubject{} // lazy initialization
	}
	fmt.Println("Calling real request from Proxy ....")
	p.realSubject.Request()

}
