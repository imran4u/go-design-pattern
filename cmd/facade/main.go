package main

import "fmt"

func main() {
	m := MediaPlayerInstance()
	m.play()
	m.stop()

}

type Payer interface {
	play()
	stop()
}
type VideoPalyer struct{}

func (v *VideoPalyer) play() {
	fmt.Println("Video play ....")
}

func (v *VideoPalyer) stop() {
	fmt.Println("Video stop ....")
}

type AudioPalyer struct{}

func (a *AudioPalyer) play() {
	fmt.Println("Audio play ....")
}

func (a *AudioPalyer) stop() {

	fmt.Println("Audio stop ....")
}

// Faced
type MediaPlayerFacade struct {
	a *AudioPalyer
	v *VideoPalyer
}

func MediaPlayerInstance() *MediaPlayerFacade {
	return &MediaPlayerFacade{
		a: &AudioPalyer{},
		v: &VideoPalyer{},
	}
}

func (m *MediaPlayerFacade) play() {
	m.a.play()
	m.v.play()
}

func (m *MediaPlayerFacade) stop() {
	m.a.stop()
	m.v.stop()
}
