package BlackJack

import "testing"

func setup()  {

}


func TestPlayerWinsForScore21(t *testing.T) {
	player := new(Player)
	dealer := new(Dealer)
	deck := new(Deck)
	blackJack := BlackJack{deck:deck, player:player, dealer:dealer}
	result := blackJack.CheckWin()

	if result != 4 {
		t.Errorf("Expected area of square to be 4, got %d", result)
	}
}
