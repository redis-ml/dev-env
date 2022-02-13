package solitaire_solver

import (
	"bufio"
	"fmt"
	"strings"
)

type Board struct {
	Stack   map[CardType]Card
	Stock   []Card
	Waste   []Card
	Piles   [][]GameCard
	Scanner *bufio.Scanner

	PreviousMoves []string
	CurrentMoves  []string
}

type UndoFunc func()

func EmptyUndoFunc() {}

func writeString(cards []Card, sb *strings.Builder) {
	for _, c := range cards {
		sb.WriteString(c.String())
	}
	sb.WriteString(";")
}

func (b Board) String() string {
	sb := new(strings.Builder)

	// Write Stack ("Foundation")
	for _, c := range AllCardTypes {
		top, ok := b.Stack[c]
		if ok {
			sb.WriteString(top.String())
		} else {
			sb.WriteString("x")
		}
		sb.WriteString(",")
	}
	sb.WriteString(";")

	writeString(b.Stock, sb)
	writeString(b.Waste, sb)

	// Write Piles ("Tableau")
	for _, pile := range b.Piles {
		for _, c := range pile {
			if c.IsRevealed() {
				sb.WriteString(c.Card.String())
			} else {
				sb.WriteString("x")
			}
		}
		sb.WriteString(";")
	}

	return sb.String()

}

func NewBoard(scanner *bufio.Scanner) *Board {
	piles := make([][]GameCard, 7)
	for i := 0; i < 7; i++ {
		piles[i] = make([]GameCard, 0, i+1)
		for j := 0; j < i+1; j++ {
			piles[i] = append(piles[i], GameCard{})
		}
	}

	ret := &Board{
		Stack:   map[CardType]Card{},
		Stock:   []Card{},
		Waste:   nil,
		Piles:   piles,
		Scanner: scanner,
	}
	return ret
}

func (b *Board) Print() {
	fmt.Printf("Stock: %s\n", b.Stock)
	fmt.Printf("Waste: %s\n", b.Waste)
	fmt.Printf("Stack: %s\n", b.Stack)
	for i, pile := range b.Piles {
		fmt.Printf("%2d ", i)
		fmt.Printf("%s\n", pile)
	}
	fmt.Printf("%s\n", b.String())
}

func (b *Board) PopStack(cardType CardType) *Card {
	card, ok := b.Stack[cardType]
	if !ok {
		return nil
	}
	if card.Number == 0 {
		delete(b.Stack, cardType)
	} else {
		b.Stack[cardType] = NewCard(cardType, card.Number-1)
	}

	return &card
}

func (b *Board) SavePileCardToStack(x int) bool {
	gameCard, y := b.GetPileTailCard(x)
	if gameCard.Card == nil {
		return false
	}
	card := *gameCard.Card

	if !b.CanSaveToStack(card) {
		return false
	}
	// action
	b.Piles[x] = b.Piles[x][0:y]
	b.Stack[card.Type] = card
	return true
}

// func (b *Board) BorrowFromStack(cardType CardType) (Card, bool) {
//    card, ok := b.Stack[cardType]
//    return card, ok
// }

func (b *Board) CanSaveToStack(card Card) bool {
	topCard, ok := b.Stack[card.Type]
	if !ok {
		return card.Number == 0
	} else {
		return card.Number == topCard.Number+1
	}
}

func (b *Board) MovePiles(x int, y int, toPile int) bool {
	if !b.CanMovePiles(x, y, toPile) {
		return false
	}
	// Move the cards
	b.Piles[toPile] = append(b.Piles[toPile], b.Piles[x][y:]...)
	b.Piles[x] = b.Piles[x][0:y]
	return true
}

func (b *Board) CanMoveToPile(src Card, toPile int) bool {
	if len(b.Piles[toPile]) == 0 {
		return src.IsK()
	} else {
		dest, _ := b.GetPileTailCard(toPile)
		return src.Number+1 == dest.Card.Number &&
			src.Color() != dest.Card.Color()
	}
}

func (b *Board) CanMovePiles(x int, y int, toPile int) bool {
	gameCard := b.GetPileCard(x, y)
	if gameCard.Card == nil {
		return false
	}
	return b.CanMoveToPile(*gameCard.Card, toPile)
}

func (b *Board) PrintComparativeMoves() {
	n := len(b.PreviousMoves)
	n2 := len(b.CurrentMoves)
	if n < n2 {
		n = n2
	}
	same := true
	for i := 0; i < n; i++ {
		prev := "<-->"
		if i < len(b.PreviousMoves) {
			prev = b.PreviousMoves[i]
		}
		curr := "<++>"
		if i < len(b.CurrentMoves) {
			curr = b.CurrentMoves[i]
		}
		if same && prev == curr {
			fmt.Printf(" %06d both: %s\n", i, prev)
		} else {
			same = false
			fmt.Printf(" %06d prev: %s\n", i, prev)
			fmt.Printf(" %06d curr: %s\n", i, curr)
		}
	}
	b.PreviousMoves = b.CurrentMoves
	b.Print()
}

func (b *Board) GetPileTailCard(x int) (GameCard, int) {
	for {
		pile := b.Piles[x]
		y := len(pile) - 1
		card := b.GetPileCard(x, y)
		// Validate card.
		if card.Card == nil {
			b.PrintComparativeMoves()
			c, ok := b.GetCardFromInput(
				fmt.Sprintf("Need to know card at position (%d, %d)", x, y),
			)
			if !ok {
				fmt.Printf("invalid input, try again!")
				continue
			}
			card.Card = &c
		}

		return card, y
	}
}

func (b *Board) AppendCardToPile(card Card, i int) bool {
	if b.CanMoveToPile(card, i) {
		tmp := NewCard(card.Type, card.Number)
		newCard := GameCard{
			Card: &tmp,
		}
		newCard.Reveal()
		b.Piles[i] = append(b.Piles[i], newCard)
		return true
	}
	return false
}

func (b *Board) GetPileCard(x int, y int) GameCard {
	gameCard := b.Piles[x][y]
	// TODO: check and ask for input if the card is unknown.
	return gameCard
}

func (b *Board) Done() bool {
	for _, cardType := range AllCardTypes {
		card, ok := b.Stack[cardType]
		if !ok || !card.IsK() {
			return false
		}
	}
	return true

}
