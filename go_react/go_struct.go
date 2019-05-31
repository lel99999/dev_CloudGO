package main

import {
	"fmt"
}

type initpower struct {
	attack int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

type nonPlayerCharacter struct {
	name string
	speed int
	hp int
	power initpower
	loc location
}

func main(){
	fmt.Println("Structs ... ")

	demon: nonPlayerCharacter{
		name: "Alfred",
		speed: 21,
		hp: 1000,
		power: initpower{attack: 75,defense: 50},
		loc: location{x: 1075.123,y:521.123,z:211.231},
	}

	fmt.Println(demon)

	anotherDemon:= nonPlayerCharacter{
		name: "Beezlebub",
		speed: 30,
		hp: 5000,
		power: power{attack: 10, defense: 10},
		loc: location{x:32.03,y:72.45,z:65.231},
	}

	fmt.Println(anotherDemon)
}
