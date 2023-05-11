package main

import "fmt"

type rockBand struct {
	name string
	vocal string
	bassist string
	drummer string
	bestSong string
}

func main() {
	
	bringMeTheHorizon := rockBand{
		name: "Bring me the horizon",
		vocal:  "Oliver Sykes",
		bassist: "Matt Kaen",
		drummer: "Matt Nicholls",
		bestSong:  "MANTRA",
	}

	updateBandName(&bringMeTheHorizon)

	fmt.Printf("%+v", bringMeTheHorizon)
		// |-> {name:Bring me the horizon vocal:Oliver Sykes bassist:Matt Kaen drummer:Matt Nicholls bestSong:MANTRA}

	myFavoriteBands := []string{"CPM-22", "Paramore", "BMTH", "É o thcan!"}

	updateBandsNameArray(myFavoriteBands)
		// |-> [UPDATED Paramore BMTH É o thcan!]

	fmt.Printf("%+v", myFavoriteBands)

}

func updateBandName(band *rockBand) {
	band.name = band.name + " UPDATED"
}

func updateBandsNameArray(bandsArray []string) {
	bandsArray[0] = bandsArray[0] + " UPDATED"
}