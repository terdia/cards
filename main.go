package main

func main() {
	cards := newDeck()
	//cards.saveToFile("mycards")
	//deck := readDeckFromFile("mycards")
	cards.shuffle()
	cards.list()
}
