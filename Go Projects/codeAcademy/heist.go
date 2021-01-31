package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var isHeistOn bool
	var eludedGuards int
	var openedVault int
	eludedGuards = rand.Intn(100)
	openedVault = rand.Intn(100)
	isHeistOn = false

	fmt.Println(isHeistOn, eludedGuards, openedVault)

	if eludedGuards >= 50 {
		fmt.Print("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
		isHeistOn = true
	} else if !isHeistOn {
		fmt.Print("Plan a better disguise next time?")
	}
	fmt.Println("")
	fmt.Println(isHeistOn, eludedGuards, openedVault)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault < 70 {
		isHeistOn = false
		fmt.Print("The vault cannot be opened.")
	}

	var leftSafely int
	leftSafely = rand.Intn(5)

	if isHeistOn {
		switch leftSafely {

		case 0:
			isHeistOn = false
			fmt.Print("The bank doors auto locked you inside!")

		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")

		case 2:
			isHeistOn = false
			fmt.Print("Lucky guard guard got the drop on you!")

		case 3:
			isHeistOn = false
			fmt.Print("The getaway car got towed!")

		case 4:
			isHeistOn = true
			fmt.Println("Start the getaway car, we are outta here!")
		}
	}

	if isHeistOn {
		var amtStolen int
		amtStolen = 10000 + rand.Intn(1000000)
		fmt.Println("We stole $", amtStolen)
	}
}
