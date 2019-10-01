	package main

	import "fmt"
	func main() {
// Various Functions of Deck have been Called here and added to comment as a flow, just remove the comments to use it.

		// cards:=deck{"kai mantos",newcard(), "bye"}
		// cards=append(cards,"b.bye")


		// cards := readFromFile("demo")
		// cards.print()
		 fmt.Println("File has been red sucessfully")

cards:= newDeck()
cards.shuffle()
cards.print()

		// fmt.Println(cards.toString())
		// hand, remainingCards := deal(cards, 5)
		// hand.print()
		// remainingCards.print()
	}

	// func newcard() string{
	// return "Hello there"
	// }
