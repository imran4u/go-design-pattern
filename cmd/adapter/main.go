package main

import (
	"fmt"
)

type OldPrinterInterface interface {
	OldPrint()
}

type OldPrinter struct{}

func (o OldPrinter) OldPrint() {
	fmt.Println("Old printer ")
}

type NewPrinterInterface interface {
	NewPrint()
}

type NewPrinter struct{}

func (n NewPrinter) NewPrint() {
	fmt.Println("New Printer style")
}

// Adapter
type PrinterAdapter struct {
	old OldPrinter
}

func (pa PrinterAdapter) NewPrint() {
	pa.old.OldPrint()
}

func main() {
	newPrinter := NewPrinter{}
	oldPrinter := OldPrinter{}

	newPrinter.NewPrint()
	oldPrinter.OldPrint() // it has no NewPrint.

	fmt.Println("After using printer adapter. ")
	pa := PrinterAdapter{old: oldPrinter}

	newPrinter.NewPrint()
	pa.NewPrint() // here same method NewPrint() but it will print older printer.

	fmt.Println("************From print job ***************. ")
	printJobs := []NewPrinterInterface{newPrinter, pa}

	for _, v := range printJobs {
		v.NewPrint()
	}

}
