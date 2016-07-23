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

	generateCardInDeck(blackjack.deck.cards, 1, 11, "A");
	generateCardInDeck(blackjack.deck.cards, 2, 2, "2");
	generateCardInDeck(blackjack.deck.cards, 3, 3, "3");
	generateCardInDeck(blackjack.deck.cards, 4, 4, "4");
	generateCardInDeck(blackjack.deck.cards, 5, 5, "5");
	generateCardInDeck(blackjack.deck.cards, 6, 6, "6");
	generateCardInDeck(blackjack.deck.cards, 7, 7, "7");
	generateCardInDeck(blackjack.deck.cards, 8, 8, "8");
	generateCardInDeck(blackjack.deck.cards, 9, 9, "9");
	generateCardInDeck(blackjack.deck.cards, 10, 10, "10");
	generateCardInDeck(blackjack.deck.cards, 10, 10, "K");
	generateCardInDeck(blackjack.deck.cards, 10, 10, "Q");
	generateCardInDeck(blackjack.deck.cards, 10, 10, "J");


	return blackjack.deck
}
func generateCardInDeck(cards []Card, value int, altValue int, name string){
	card := new(Card)
	card.value = value;
	card.altValue = altValue;
	card.name = name
	cards = append(cards,card)
}

