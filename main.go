package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(a string) Profile {
	var profile Profile

	if a == "Goku" {
		profile.Name = a
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	}

	if a == "Sasuke" {
		profile.Name = a
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	}

	if a == "Naruto" {
		profile.Name = a
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}

	return profile
}

func PowerUp(a Profile, b int) Profile {
	var powerUp Profile

	powerUp.Name = a.Name
	powerUp.Health = a.Health * b
	powerUp.Power = a.Power * b
	powerUp.Exp = a.Exp * b

	return powerUp
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 5)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}
