package main

import (
	"fmt"

	tp "github.com/imran4u/go-design-pattern/cmd/decorator/third_party"
)

func main() {
	// Directly using third party
	soldier := tp.BasicSoldier{}
	tp.Print(soldier)

	// using after decoration
	fmt.Println("\nAttack will increase and defence will decrease")
	swordSoldier := SwordSoldier{soldier}
	tp.Print(swordSoldier)

	// More attach and more defence
	fmt.Println("\nMore attach and more defence")
	airSolder := AirSoldier{soldier}
	tp.Print(airSolder)

}

type SwordSoldier struct {
	soldier tp.SoldierInterface
}

func (s SwordSoldier) Attack() int {
	return s.soldier.Attack() + 5 // increase attach by 5.
}

func (s SwordSoldier) Defance() int {
	return s.soldier.Defance() - 3 // Decresed Defence by 3
}

type AirSoldier struct {
	soldier tp.SoldierInterface
}

func (a AirSoldier) Attack() int {
	return a.soldier.Attack() + 15 // increase attach by 15.
}

func (a AirSoldier) Defance() int {
	return a.soldier.Defance() + 6 // Decresed Defence by 6
}
