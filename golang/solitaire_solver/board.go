package solitaire_solver

import (
  "fmt"
)

type Board struct {
  Stack map[CardType]Card
  Stock []Card
  Waste []Card
  Piles [][]GameCard
}

func NewBoard() *Board {
  piles := make([][]GameCard, 7)
  for i := 0; i < 7; i++ {
    piles[i] = make([]GameCard, 0, i+1)
    for j := 0; j < i + 1; j++ {
      piles[i] = append(piles[i], GameCard{})
    }
  }

  ret := &Board {
    Stack: map[CardType]Card{},
    Stock: []Card{},
    Waste: nil,
    Piles: piles,
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
}

func (b *Board) SavePileCardToStack(x int, y int) bool {
   gameCard := b.Piles[x][y]
   if gameCard.Card == nil {
     return false
   }
   card := *gameCard.Card

   if !b.CanSaveToStack(card) {
     return false
   }
   // action
   return b.SaveToStack(card)
}

func (b *Board) BorrowFromStack(cardType CardType) (Card, bool) {
   card, ok := b.Stack[cardType]
   return card, ok
}

func (b *Board) SaveToStack(card Card) bool {
   b.Stack[card.Type] = card
   return true
}

func (b *Board) CanSaveToStack(card Card) bool {
    topCard, ok := b.Stack[card.Type]
    if !ok {
      return card.Number == 0
    } else {
      return card.Number == topCard.Number + 1
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

func (b *Board) CanMovePiles(x int, y int, toPile int) bool {
    gameCard := b.Piles[x][y]
    if gameCard.Card == nil {
      return false
    }

    if len(b.Piles[toPile]) == 0 {
      // case 1
      return gameCard.Card.IsK()
    } else {
      pile := b.Piles[toPile]
      dest := pile[len(pile) - 1]
      return gameCard.Card.Number + 1 == dest.Card.Number &&
        gameCard.Card.Color() != dest.Card.Color()
    }
}

