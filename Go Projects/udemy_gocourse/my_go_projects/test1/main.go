package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* g00,r1,b13,r36,b24,r3,b15,r34,b22,
r5,b17,r32,b20,r7,b11,r30,b26,r9,b28,
g0,b2,r14,b35,r23,b4,r16,b33,r21,b6,
r18,b31,r19,b8,r12,b29,r25,b10,r27 */

func main() {

	const spins = 10000 // Number of spins to run.
	var black int       // Roulette Black.
	var red int         // Roulette Red.
	var doubleZero int  // Roulette Doubble Zero.
	var zero int        // Roulette Zero.
	var green int       // Roulette Green.
	var dist [38]int    // Keep track of each number that was hit.
	wheel := [38]string{
		"g00", "r1", "b13", "r36", "b24", "r3", "b15", "r34", "b22",
		"r5", "b17", "r32", "b20", "r7", "b11", "r30", "b26", "r9", "b28",
		"g0", "b2", "r14", "b35", "r23", "b4", "r16", "b33", "r21", "b6",
		"r18", "b31", "r19", "b8", "r12", "b29", "r25", "b10", "r27"} // The Roulette Wheel

	var i, j int // Used in for loops.

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

/*=============================================*/
/* Generate a random number to be used in      */
/* Roulette to simulate a spin and the ball.   */
/* Returns the random number.                  */
/*=============================================*/
func getRandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	newPosition := r.Intn(39 - 1)
	return newPosition
}

/*=============================================*/
/* Calculate the total results of the          */
/* Roulette colors and zeros.                  */
/* Inputs pointers to the values to be tracked.*/
/* Indirectaly modified.                       */
/* Inputs a new position variable to           */
/* determine what color or zone was hit on     */
/* the spin.                                   */
/*=============================================*/
func calcResults(z, dz, g, b, r *int, np int) {
	var test int // Used to determine the colors and zeros on the Roulette wheel by sending the result through a sorting sieve.

	if (np % 2) == 0 { // Even and Black. Uses modulo remainder to dertrmine if this is an even number.
		test = 0
	} else {
		test = 1 // Red or Green. Odd numbers
	}

	if np == 0 { // Double Zero and Green. Spot 0 on the Roulette wheel is Double zero.
		test = 3
	}

	if np == 19 { // Single Zero and Green. Spot 19 on the Roulette wheel is Single zero.
		test = 4
	}

	switch test {
	case 0: // Increment Black.
		*b++
	case 1: // Increment Red.
		*r++
	case 3: // Increment Double Zero and Green.
		*dz++
		*g++
	case 4: // Increment Zero and Green.
		*z++
		*g++
	default: // No default.
	}
}

/*=============================================*/
/* Prints the calculated results of            */
/* the total spins                             */
/*=============================================*/
func printResults(z, dz, g, b, r, spinNum int) {
	fmt.Printf("\nTotal spins: %v\n", spinNum)
	fmt.Printf("\nDouble Zero = %v\n", dz)
	fmt.Printf("Zero = %v\n", z)
	fmt.Printf("Green = %v\n", g)
	fmt.Printf("Black = %v\n", b)
	fmt.Printf("Red = %v\n", r)
}
