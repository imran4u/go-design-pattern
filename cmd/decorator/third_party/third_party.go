package thirdparty

import "fmt"

//It is third party code so you can't change it.

type SoldierInterface interface {
	Attack() int
	Defance() int
}

type BasicSoldier struct{}

func (b BasicSoldier) Attack() int {
	return 5
}

func (b BasicSoldier) Defance() int {
	return 5
}

func Print(b SoldierInterface) {
	fmt.Printf("Attack = %d, Defance = %d \n", b.Attack(), b.Defance())
}
