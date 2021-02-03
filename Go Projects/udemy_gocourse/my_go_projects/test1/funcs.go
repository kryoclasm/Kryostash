package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
