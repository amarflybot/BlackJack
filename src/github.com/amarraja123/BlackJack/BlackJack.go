package BlackJack


type BlackJack struct {
	deck *Deck
	dealer *Dealer
	player *Player
}

type Dealer struct {
	score int
}

type Player struct {
	score int
}

type Winner interface {

}

type Deck struct {
	cards []string
}

func (blackjack BlackJack) CheckWin() Winner {
	if blackjack.player.score == 21{
		return blackjack.player
	}
	return blackjack.dealer
}

func (blackjack BlackJack) NewDeck() *Deck {
	deck := new(Deck)
	deck.cards = []string{"A","A","A","A",
	"2","2","2","2",
	"3","3","3","3",
	"4","4","4","4",
	"10","10","10","10",
	"10","10","10","10",
	"10","10","10","10",
	}
	return deck
}
