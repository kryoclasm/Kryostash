package main

import (
	"fmt"
	"time"
)

//
/* g00,r1,b13,r36,b24,r3,b15,r34,b22,
r5,b17,r32,b20,r7,b11,r30,b26,r9,b28,
g0,b2,r14,b35,r23,b4,r16,b33,r21,b6,
r18,b31,r19,b8,r12,b29,r25,b10,r27 */

func main() {

	const spins = 2000 // Number of spins to run.
	var black int      // Roulette Black.
	var red int        // Roulette Red.
	var doubleZero int // Roulette Doubble Zero.
	var zero int       // Roulette Zero.
	var green int      // Roulette Green.
	var dist [38]int   // Keep track of each number that was hit.
	wheel := [38]string{
		"g00", "r1", "b13", "r36", "b24", "r3", "b15", "r34", "b22",
		"r5", "b17", "r32", "b20", "r7", "b11", "r30", "b26", "r9", "b28",
		"g0", "b2", "r14", "b35", "r23", "b4", "r16", "b33", "r21", "b6",
		"r18", "b31", "r19", "b8", "r12", "b29", "r25", "b10", "r27"} // The Roulette Wheel

	var i, j int // Used in for loops.

	fmt.Printf("Calculating the results of %v games of Roulette!\n", spins)
	fmt.Printf("Please be patient...\n\n")

	for i <= (spins - 1) {
		newPosition := getRandomNumber()                                   // Get a random number seed based on the Unix time.
		time.Sleep(1 * time.Millisecond)                                   // Wait 1 Millisecond to better help with the random number generation spread.
		dist[newPosition]++                                                // Increment each number from each spin. Shows how many times a number on the wheel was hit.
		calcResults(&zero, &doubleZero, &green, &black, &red, newPosition) // Keep a record of the results to be printed after last spin.
		i++
	}

	/* I could implement a sort to represent
	   each number from least hit to most hit. */

	j = 0
	for j <= 37 {
		fmt.Printf("(Number %v = %v)\n", wheel[j], dist[j]) // Prints the total of each of the 38 numbers on the Roulette wheel.
		j++
	}
	printResults(zero, doubleZero, green, black, red, spins) // Prints the total of each color and the zeros.
}
