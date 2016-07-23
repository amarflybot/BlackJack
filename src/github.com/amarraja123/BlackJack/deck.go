package main

type Deck struct {
	cards []Card
}

type Card struct {
	value int
	altValue int
	name string
}



func (blackjack *BlackJack) NewDeck() *Deck {
	blackjack.deck = new(Deck)

	card := new(Card)
	card.value = 1;
	card.altValue = 11;
	card.name = "A"
	blackjack.deck.cards = append(blackjack.deck.cards,card)

	return blackjack.deck
}
