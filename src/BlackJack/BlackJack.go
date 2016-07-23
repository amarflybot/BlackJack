package BlackJack


type BlackJack struct {
	deck Deck
	dealer Dealer
	player Player
}

type Dealer struct {
	score int
}

type Player struct {
	score int
}

type Card struct {
	value string
}

type Deck struct {
	cards []Card
}
