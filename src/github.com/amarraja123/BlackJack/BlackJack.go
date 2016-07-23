package main

import (
	"math/rand"
	"fmt"
	"bufio"
	"os"
)

type BlackJack struct {
	deck *Deck
	dealer *Dealer
	player *Player
}

type Winner interface {

}

func (blackjack *BlackJack) CheckWin() Winner {
	if blackjack.player.score == 21{
		return blackjack.player
	} else if blackjack.player.score >= 21{
		return blackjack.dealer
	}
	return nil
}

func (blackJack *BlackJack)PlayGame(playerChance bool)  {
	gameOver := false;
	for !gameOver{
		if playerChance{
			//Select Card from the deck
			cardValue := blackJack.GetCard();
			fmt.Println("Card Value of the user %s ",cardValue.name)
			//Get Input from user
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter HIT or STAND: ")
			text, _ := reader.ReadString('\n')
			if text == "HIT\n" {
				blackJack.player.score += cardValue.value;
				winner := blackJack.CheckWin()
				if winner == nil {
					continue
				} else {
					gameOver = true
					fmt.Printf("Winner : %v ", winner)
				}

			} else if text == "STAND\n"{
				playerChance = false
				continue
			}
		} else {
			//Select Card from the deck
			cardValue := blackJack.GetCard();
			fmt.Println("Card Value of the user %s ",cardValue.name)
				blackJack.dealer.score += cardValue.value;
				winner := blackJack.CheckWin()
				if winner == nil {
					continue
				} else {
					gameOver = true
					fmt.Printf("Winner : %v " , winner)
				}
		}
	}
}


func (blackjack *BlackJack) GetCard() Card {
	sliceIndex := rand.Intn(len(blackjack.deck.cards))
	randomCard := blackjack.deck.cards[sliceIndex]
	//remove card from slice after selection
	blackjack.deck.cards = append(blackjack.deck.cards[:sliceIndex], blackjack.deck.cards[sliceIndex+1:]...)
	return randomCard
}
