package solitaire_solver

import (
  "fmt"
)

func (b *Board) WasteToStack() (string, UndoFunc, bool) {
  l := len(b.Waste) 
  if l > 0 {
    card := b.Waste[l-1]
    if b.CanSaveToStack(card) {
      b.Waste = b.Waste[0:l-1]
      b.Stack[card.Type] = card

      undo := func() {
        b.Waste = append(b.Waste, card)
        b.PopStack(card.Type)
      }
      cmd := fmt.Sprintf("waste %s -> stack", card)
      return cmd, undo, true
    }
  }
  return "", func(){}, false
}

func (b *Board) WasteToPile(pileIdx int) (string, UndoFunc, bool) {
  l := len(b.Waste) 
  if l > 0 {
    card := b.Waste[l-1]
    if b.CanMoveToPile(card, pileIdx) {
      originalPileLen := len(b.Piles[pileIdx])
      tmp := NewCard(card.Type, card.Number)
      newCard := GameCard {
        Card: &tmp,
      }
      newCard.Reveal()

      b.Piles[pileIdx] = append(b.Piles[pileIdx], newCard)

      b.Waste = b.Waste[0:l-1]


      move := fmt.Sprintf("waste %s -> (%d, %d)", card, pileIdx, originalPileLen)
      undo := func() {
        b.Waste = append(b.Waste, card)
        b.Piles[pileIdx] = b.Piles[pileIdx][0:originalPileLen]
      }
      return move, undo, true
    }
  }
  return "", func(){}, false
}
