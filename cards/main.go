
package main

import "fmt"

func main(){
	//var card string = "Ace of Spades"
	card := "Ace of Spades" // initlizaion statment 
	card = "Five of Diamonds"

	arr_card := []string{"six of spade", newCard()}
	arr_card = append(arr_card,card)

	for i, card := range arr_card{
		fmt.Println(i,card)
	}

	fmt.Println(card , arr_card)
}

func newCard() string{
	return "Five of Diamonds"
}