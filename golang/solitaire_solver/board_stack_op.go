package solitaire_solver

import (
  "fmt"
)

func (b *Board) StackToPile(cardType CardType) (string, UndoFunc, bool) {
  card, ok := b.Stack[cardType]
  if ok && card.Number > 1 {
    for i, pile := range b.Piles {
      j := len(pile)
      if b.AppendCardToPile(card, i) {
        b.PopStack(cardType)

        // Recursion.
        move := fmt.Sprintf("stack %s -> (%d, %d)", card, i, j)

        undo := func() {
          b.Stack[cardType] = card
          b.Piles[i] = b.Piles[i][0:j]
        }
        return move, undo, true
      }
    }
  }
  return "", func(){}, false
}
