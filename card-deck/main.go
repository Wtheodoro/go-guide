package main

// just for learn, we can import any function from deck.go and run like this go run main.go deck.go

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
}
