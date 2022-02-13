package solitaire_solver

import (
  "fmt"
)

func (b *Board) PileToStack(pileIdx int) (string, UndoFunc, bool) {
  srcI := pileIdx
  if len(b.Piles[pileIdx]) == 0 {
    return "", func(){}, false
  }
  gameCard, srcJ := b.GetPileTailCard(srcI)
  card := NewCard(gameCard.Card.Type, gameCard.Card.Number)

  if b.SavePileCardToStack(srcI) {
    move := fmt.Sprintf("pile card %s (%d, %d) -> stack", card, srcI, srcJ)

    // Check if a new card is revealed.
    isNewCardRevealed := false
    if srcJ > 0 {
      prevGameCard, _ := b.GetPileTailCard(srcI)
      if !prevGameCard.IsRevealed() {
        isNewCardRevealed = true
        prevGameCard.Reveal()
        b.Piles[srcI][srcJ-1] = prevGameCard
      }
    }

    undo := func() {
      fmt.Printf("UNDO >>>> %s\n", move)
      if isNewCardRevealed {
        // unreveal
        prevGameCard, _ := b.GetPileTailCard(srcI)
        prevGameCard.Unreveal()
        b.Piles[srcI][srcJ-1] = prevGameCard
      }
      b.PopStack(card.Type)
      b.Piles[srcI] = append(b.Piles[srcI], gameCard)
    }
    return move, undo, true
  }
  return "", func(){}, false
}

func (b *Board) PileToPile(x int, y int, dst int) (string, UndoFunc, bool) {
  srcI, srcJ := x, y
  i := dst

  if i == srcI {
    return "", func(){}, false
  }

  gameCard := b.GetPileCard(srcI, srcJ)
  if !gameCard.IsRevealed() {
    return "", func(){}, false
  }
  card := NewCard(gameCard.Card.Type, gameCard.Card.Number)

  pile := b.Piles[i]
  j := len(pile)
  if b.CanMoveToPile(card, i) {
    // Check if a new card is revealed.
    isNewCardRevealed := false

    if srcJ > 0 {
      prevGameCard := b.GetPileCard(srcI, srcJ-1)
      // Check if such move is necessary.
      if prevGameCard.IsRevealed() && !b.CanSaveToStack(*prevGameCard.Card) {
        return "", func(){}, false
      }

      if !prevGameCard.IsRevealed() {
        isNewCardRevealed = true
      }
    }

    // [Make Changes] Actually update the piles.
    b.MovePiles(srcI, srcJ, i)

    // Now let's get input of the new card.
    if isNewCardRevealed {
      prevGameCard, _ := b.GetPileTailCard(srcI)
      prevGameCard.Reveal()
      b.Piles[srcI][srcJ-1] = prevGameCard
    }

    move := fmt.Sprintf("pile card %s (%d, %d) -> (%d, %d)", card, srcI, srcJ, i, j)
    undo := func() {
      if isNewCardRevealed {
        // unreveal
        prevGameCard, _ := b.GetPileTailCard(srcI)
        prevGameCard.Unreveal()
        b.Piles[srcI][srcJ-1] = prevGameCard
      }
      b.Piles[srcI] = append(b.Piles[srcI], b.Piles[i][j:]...)
      b.Piles[i] = b.Piles[i][0:j]
    }
    return move, undo, true
  }
  return "", func(){}, false
}
