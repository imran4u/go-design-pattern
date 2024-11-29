package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}
type Tv struct {
	brand string
}

func (t *Tv) TurnOn() {
	fmt.Println("Tv ON")
}

func (t *Tv) TurnOff() {
	fmt.Println("Tv OFF")
}

type Radio struct {
	brand string
	year  int
}

func (r *Radio) TurnOn() {
	fmt.Println("Radio ON")
}

func (r *Radio) TurnOff() {
	fmt.Println("Radio OFF")
}

// Bridge implmentation

type DeviceController struct {
	d Device
}

func (dc *DeviceController) DeviceOn() {
	dc.d.TurnOn()
}
func (dc *DeviceController) DeviceOff() {
	dc.d.TurnOff()
}

func main() {

	tv := Tv{"sanmsung"}
	radio := Radio{"phlips", 1998}

	tvController := DeviceController{&tv}
	radioController := DeviceController{&radio}

	// tv control
	tvController.DeviceOn()  //Tv ON
	tvController.DeviceOff() //Tv OFF

	radioController.DeviceOn()  //Radio On
	radioController.DeviceOff() //Radio off

}
