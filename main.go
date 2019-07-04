package main

func main() {

	cards := deck{"ace of diamonds", newCard()}
	cards.printDeck()
}

func newCard() string {
	return "five diamonds"

}
