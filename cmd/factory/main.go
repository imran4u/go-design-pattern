package main

import "fmt"

type Animal interface {
	speak() string
}

type Cat struct{}

func (c Cat) speak() string {
	return "Meow"
}

type Dog struct{}

func (d Dog) speak() string {
	return "Bark/ Bhow"
}

func AnimalFactory(name string) Animal {
	switch name {
	case "cat":
		return &Cat{}
	case "dog":
		return &Dog{}
	default:
		return nil
	}
}

func main() {
	dog := AnimalFactory("dog")
	cat := AnimalFactory("cat")
	fmt.Println(dog.speak())
	fmt.Println(cat.speak())

}
