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
    pileLen := len(b.Piles[pileIdx])
    if b.AppendCardToPile(card, pileIdx) {
      b.Waste = b.Waste[0:l-1]

      move := fmt.Sprintf("waste %s -> (%d, %d)", card, pileIdx, pileLen)
      undo := func() {
        b.Waste = append(b.Waste, card)
        b.Piles[pileIdx] = b.Piles[pileIdx][0:pileLen]
      }
      return move, undo, true
    }
  }
  return "", func(){}, false
}
