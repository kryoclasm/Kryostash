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
		fmt.Print("Grab and GO!")
	} else if isHeistOn && openedVault < 70 {
		isHeistOn = false
		fmt.Print("The vault cannot be opened.")
	}

}
