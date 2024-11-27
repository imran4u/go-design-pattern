package main

import "fmt"

type Car struct {
	Make     string
	Model    string
	Year     int
	Color    string
	IsPetrol bool
}

type CarBuilder struct {
	car Car
}

func NewBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) setMake(make string) *CarBuilder {
	b.car.Make = make
	return b
}
func (b *CarBuilder) setModel(model string) *CarBuilder {
	b.car.Model = model
	return b
}
func (b *CarBuilder) setYear(year int) *CarBuilder {
	b.car.Year = year
	return b
}
func (b *CarBuilder) setColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) setIsPertorl(IsPetrol bool) *CarBuilder {
	b.car.IsPetrol = IsPetrol
	return b
}

func (b *CarBuilder) build() Car {
	return b.car
}

func main() {
	builder := NewBuilder()
	car := builder.setMake("Honda").
		setColor("Blue").
		setIsPertorl(true).
		setYear(1990).
		build()

	fmt.Printf("Car : %+v \n", car)
}
