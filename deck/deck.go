package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	Spade CardSuit = iota
	Diamond
	Club
	Heart
	Color
	Gray
)

const (
	Joker CardValue = iota
	Ace 
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var cardSuits = []CardSuit{
	Spade,
	Diamond,
	Club,
	Heart,
}

var cardValues = []CardValue{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

type CardSuit int

func (cs CardSuit) String() string {
	switch cs {
	case Color:
		return "Color"
	case Gray:
		return "Gray"
	case Spade:
		return "Spade"
	case Diamond:
		return "Diamond"
	case Club:
		return "Club"
	case Heart:
		return "Heart"
	default:
		return ""
	}
}

type CardValue int

func (cv CardValue) String() string {
	switch cv {
	case Joker:
		return "Joker"
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}

type Card struct {
	Suit CardSuit
	Value CardValue
}

func New() *[]Card {
	var deck []Card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}

	return &deck
}

func Shuffle(deck *[]Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*deck), func(i, j int) {
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	})
}

func AddJokersToDeck(deck *[]Card, totalNumber int, numberOfColors int) {
	if totalNumber < 0 {
		fmt.Println("Total number of jokers cant be less than 0")
		return
	}
	if numberOfColors < 0 {
		fmt.Println("Number of colored jokers cant be less than 0")
		return
	}
	if numberOfColors > totalNumber {
		fmt.Println("Number of colored jokers cant be more than total number of jokers")
		return
	}

	for i := 0; i < totalNumber; i++ {
		if numberOfColors > 0 {
			*deck = append(*deck, Card{Suit: Color, Value: Joker})
			numberOfColors--
			continue
		}
		*deck = append(*deck, Card{Suit: Gray, Value: Joker})
	}
}

func FilterCards(deck *[]Card, values... CardValue) {
	deckLen := len(*deck)
	i := 0
	Loop:
	for i < deckLen {
		for _, value := range values {
			if (*deck)[i].Value == value {
				*deck = append((*deck)[:i], (*deck)[i+1:]...)
				deckLen = len(*deck)
				continue Loop
			}
		}
		i++
	}
}

func SortDeckFunc(deck *[]Card, sortFunc func(i, j int) bool) {
	sort.Slice(*deck, sortFunc)
}

func SortDeck(deck *[]Card) {
	SortDeckFunc(deck, func(i, j int) bool {
		if (*deck)[i].Suit == (*deck)[j].Suit {
			return (*deck)[i].Value < (*deck)[j].Value
		} else {
			return (*deck)[i].Suit > (*deck)[j].Suit
		}
	})
}

func GetCard(deck *[]Card) Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}