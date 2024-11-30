package main

import (
	"fmt"
	"strings"
)

type Component interface {
	Display() string
}

type Leaf struct {
	name string
}

func (l *Leaf) Display() string {
	return l.name
}

// Composite
type Composite struct {
	name     string
	children []Component // type of interface so that composite and Leaf both can fit here
}

func (c *Composite) AddChild(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) Display() string {
	var result strings.Builder
	result.WriteString("{\n")
	result.WriteString(c.name)
	result.WriteString(":\n")
	for _, child := range c.children {
		result.WriteString(child.Display()) // here is recursive call.
		result.WriteString(",\n")
	}

	result.WriteString("}")

	// Note: you can adjust display as per your desire
	return result.String()
}

func main() {
	l1 := Leaf{name: "Leaf one"}
	l2 := Leaf{name: "Leaf Two"}
	l3 := Leaf{name: "Leaf Three"}
	fmt.Println(l1.Display())

	c1 := Composite{name: "Composite One"}
	c2 := Composite{name: "Composite Two"}

	c1.AddChild(&l1)
	c1.AddChild(&l2)

	c2.AddChild(&l3)

	c2.AddChild(&c1)

	fmt.Println(c2.Display())
}
