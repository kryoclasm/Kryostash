package main

import "fmt"

func main() {
	publisher, writer, artist, title := "", "", "", ""
	year, pageNumber := 0, 0
	grade := 0.0

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println("Title:", title+"\n"+"Written by:", writer+"\n"+"Drawn by:", artist)
	fmt.Println("Publisher:", publisher+"\n"+"Year Published:", year, "\n"+"Number of Pages:", pageNumber)
	fmt.Println("Condition Grade:", grade)
	fmt.Println("")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println("Title:", title+"\n"+"Written by:", writer+"\n"+"Drawn by the Spiegeler-winning:", artist)
	fmt.Println("Publisher:", publisher+"\n"+"Year Published:", year, "\n"+"Number of Pages:", pageNumber)
	fmt.Println("Condition Grade:", grade)

}
