package main

import "fmt"

// opne-close: open for extension close for modification

type Color int

const (
	red Color = iota
	gree
	blue
)

var colorMaps = map[Color]string{
	red:  "Red",
	gree: "Green",
	blue: "Blue",
}

type Size int

const (
	small Size = iota + 1
	medium
	large
)

type product struct {
	name  string
	color Color
	size  Size
}

func (p product) string() string {
	return fmt.Sprintf("Name: %s, color = %s, size = %v", p.name, colorMaps[p.color], p.size)
}

type Filter struct {
}

// it is breaking as we are incresing the responsiblity of products
func (f Filter) filterBySize(products []product, size Size) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.size == size {
			resut = append(resut, &p)
		}
	}
	return resut

}

// it is breaking as we are incresing the responsiblity of products
func (f Filter) filterByColor(products []product, color Color) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.color == color {
			resut = append(resut, &p)
		}
	}
	return resut

}

// it is breaking as we are incresing the responsiblity of products
func (f Filter) filterByColorAndSize(products []product, color Color, size Size) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.color == color && size == p.size {
			resut = append(resut, &p)
		}
	}
	return resut

}

// ################################ implementation of open for extension and closed for modification ######

type Specification interface {
	isSatisfied(p *product) bool
}

type ColorSpecification struct {
	Color
}

func (c ColorSpecification) isSatisfied(p *product) bool {
	return p.color == c.Color
}

type SizeSpecification struct {
	Size
}

func (s SizeSpecification) isSatisfied(p *product) bool {
	return s.Size == p.size
}

type AndSpecification struct {
	color Color
	size  Size
}

func (a AndSpecification) isSatisfied(p *product) bool {
	return a.size == p.size && a.color == p.color
}

type BetterFilter struct {
}

func (b BetterFilter) filter(products []product, s Specification) []*product {
	resut := make([]*product, 0)
	for _, p := range products {
		if s.isSatisfied(&p) {
			resut = append(resut, &p)
		}
	}
	return resut
}

func main() {

	apple := product{"Apple", red, small}
	orange := product{"Orange", gree, small}
	elephant := product{"Elephant", red, large}

	products := []product{
		apple,
		orange,
		elephant,
	}

	fmt.Println(" old way of filter, breaking the open-close principle")

	f := Filter{}
	fmt.Println(" filter by size small ")
	for _, p := range f.filterBySize(products, small) {
		fmt.Println(p.string())
	}
	fmt.Println(" filter by color  red ")
	for _, p := range f.filterByColor(products, red) {
		fmt.Println(p.string())
	}

	fmt.Println(" filter by color  gree and size small ")
	for _, p := range f.filterByColorAndSize(products, gree, small) {
		fmt.Println(p.string())
	}

	fmt.Println(" using better filter :  ")
	bf := BetterFilter{}
	cs := ColorSpecification{gree}
	ss := SizeSpecification{small}
	as := AndSpecification{gree, small}
	fmt.Println(" filter by color  gree  ")
	for _, p := range bf.filter(products, cs) {
		fmt.Println(p.string())
	}

	fmt.Println(" filter by size small ")
	for _, p := range bf.filter(products, ss) {
		fmt.Println(p.string())
	}

	fmt.Println(" filter by color  gree and size small ")
	for _, p := range bf.filter(products, as) {
		fmt.Println(p.string())
	}

}
