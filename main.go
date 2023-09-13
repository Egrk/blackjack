package main

import (
	"blackjack/deck"
	"fmt"
)

type BJPlayer struct {
	Hand []deck.Card
	HaveAce bool
	Turn bool
	Points int
}

func (d *BJPlayer) AddCardToHand(card deck.Card) {
	d.Hand = append(d.Hand, card)
	if card.Value == deck.Ace {
		d.HaveAce = true
	}
	d.Points += getValue(&card)
}

func (d *BJPlayer) ShowHand() {
	if d.Turn {
		for i, item := range d.Hand {
			fmt.Printf("%d. %s %s\n", i + 1, item.Suit, item.Value)
		}
		fmt.Printf("Total: %d\n", d.Points)
	} else {
		for i, item := range d.Hand {
			if i == 0 {
				fmt.Printf("%d. %s %s\n", i + 1, item.Suit, item.Value)
			} else {
				fmt.Printf("%d. %s\n", i + 1, "Back of a card")
			}
		}
		fmt.Printf("Total: %d\n", getValue(&d.Hand[0]))
	}
}

func getValue(card *deck.Card) int {
	switch card.Value {
	case deck.Jack, deck.King, deck.Queen:
		return 10
	case deck.Ace:
		return 11
	default:
		return int(card.Value)
	}
}

func main() {
	fmt.Println("Welcome to BlackJack game!")
	bjDeck := deck.New()
	*bjDeck = append(*bjDeck, *deck.New()...)
	deck.Shuffle(bjDeck)
	var player BJPlayer
	var dealer BJPlayer
	player.AddCardToHand(deck.GetCard(bjDeck))
	dealer.AddCardToHand(deck.GetCard(bjDeck))
	player.AddCardToHand(deck.GetCard(bjDeck))
	dealer.AddCardToHand(deck.GetCard(bjDeck))
	dealer.Turn = false
	player.Turn = true
	for {
		fmt.Printf("\x1bc")
		fmt.Println("Dealer Hand:")
		dealer.ShowHand()
		fmt.Println("Player Hand:")
		player.ShowHand()
		if player.Points == 21 {
			fmt.Println("You win!")
			break
		} else if player.Points > 21 {
			if player.HaveAce {
				player.Points -= 10
				player.HaveAce = false
				continue
			} else {
				fmt.Println("You lose...")
				break
			}
		} else if dealer.Points > 21 {
			if dealer.HaveAce {
				dealer.Points -= 10
				dealer.HaveAce = false
				continue
			} else {
				fmt.Println("You win!")
				break
			}
		}
		if !dealer.Turn {
			fmt.Print("What would you do?:\n1. Get more\n2. Done\n")
			var decision int
			fmt.Scan(&decision)
			switch decision {
			case 1:
				player.AddCardToHand(deck.GetCard(bjDeck))
			case 2:
				dealer.Turn = true
			}
		} else {
			fmt.Println("Dealer turn")
			if dealer.Points <= 16 {
				dealer.AddCardToHand(deck.GetCard(bjDeck))
			} else {
				if dealer.Points > player.Points {
					fmt.Println("Casino win. He always win...")
					break
				} else if dealer.Points < player.Points {
					fmt.Println("You win, wow!!!")
					break
				} else {
					fmt.Println("Draw!")
					break
				}
			}
		}
	}
}