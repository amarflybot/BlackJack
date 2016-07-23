package main

import (
	"testing"
)

func TestPlayerWinsForScore21(t *testing.T) {
	player := new(Player)
	player.score = 21
	dealer := new(Dealer)
	blackJack := BlackJack{player:player, dealer:dealer}
	blackJack.NewDeck()
	result := blackJack.CheckWin()
	switch v := result.(type) {
	case *Player:
		t.Logf("Expected type Player and got %+v\n", result)
	default:
		t.Errorf("Expected type player got %v", v)
	}
}

func TestPlayerWinsForDealerScoreIsLess(t *testing.T) {
	player := new(Player)
	dealer := new(Dealer)
	blackJack := BlackJack{player:player, dealer:dealer}
	blackJack.NewDeck()
	blackJack.PlayGame(false)
	result := blackJack.CheckWin()
	switch v := result.(type) {
	case *Player:
		t.Log("Expected type Player got Player")
	default:
		t.Errorf("Expected type player got %d", v)
	}
}
