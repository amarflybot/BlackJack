package main

func main() {

	player := new(Player)
	dealer := new(Dealer)
	blackJack := BlackJack{player:player, dealer:dealer}
	blackJack.NewDeck()
	blackJack.PlayGame(true)

}
