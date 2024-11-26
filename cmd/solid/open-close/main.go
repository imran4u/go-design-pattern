package main

import "fmt"

// opne-close: open for extension close for modification

type Color int

const (
	red Color = iota
	gree
	blue
)

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
	return fmt.Sprintf("Name: %s, color =%v, size = %v", p.name, p.color, p.size)
}

type Filter struct {
}

func (f Filter) filterBySize(products []product, size Size) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.size == size {
			resut = append(resut, &p)
		}
	}
	return resut

}
func (f Filter) filterByColor(products []product, color Color) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.color == color {
			resut = append(resut, &p)
		}
	}
	return resut

}

func (f Filter) filterByColorAndSize(products []product, color Color, size Size) []*product {
	resut := make([]*product, 0)

	for _, p := range products {
		if p.color == color && size == p.size {
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
}
