
package main
// import "fmt"



func main(){
	// made the executable My_cards File
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("My_cards")

	// cards:=newDeckFromFile("my_cards")
	// cards.print()

	cards:= newDeck()
	cards.shuffle()
	cards.print()
}
