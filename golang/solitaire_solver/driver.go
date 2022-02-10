package solitaire_solver

import (
  "fmt"
  "bufio"
)

type Driver struct {
  Scanner *bufio.Scanner
  Board *Board
}

func (d *Driver) Solve(
  moves []string,
  visited map[string]bool,
) bool {
  // d.Board.Print()
  // fmt.Println("#################")
  // fmt.Printf("%s\n", moves)
  // fmt.Println("#################")

  status := d.Board.String()
  if visited[status] {
    fmt.Printf("duplicated status, return false..")
    return false
  }
  visited[status] = true


  if d.Board.Done() {
    fmt.Println("#################")
    fmt.Printf("%s\n", moves)
    return true
  }

  // Hand operations
  if len(d.Board.Hand) > 0 {
    // 1. Hand -> Stack
    card := d.Board.Hand[len(d.Board.Hand) - 1]
    stackCard, ok := d.Board.Stack[card.Type]
    if ok && stackCard.Number + 1 == card.Number {
      // Make the move.
      d.Board.Hand = d.Board.Hand[0:len(d.Board.Hand)-1]
      d.Board.Stack[card.Type] = card
      // recursion
      newMove := append(moves, fmt.Sprintf("hand %s -> stack", card))
      ret := d.Solve(newMove, visited)
      if ret {
        return true
      }

      // Reset
      d.Board.Hand = append(d.Board.Hand, card)
      d.Board.Stack[card.Type] = stackCard
    }

    // 2. Hand -> Pile
    for i, pile := range d.Board.Piles {
      j := len(pile)
      if d.Board.CanMoveToPile(card, i) {
        newCard := GameCard {
          Card: &card,
        }
        newCard.Reveal()
        d.Board.Piles[i] = append(d.Board.Piles[i], newCard)
        d.Board.Hand = d.Board.Hand[0:len(d.Board.Hand)-1]

        // Recursion.
        newMove := append(moves,
          fmt.Sprintf("hand %s -> (%d, %d)", card, i, j))
        ret := d.Solve(newMove, visited)
        if ret {
          return true
        }

        // Reset
        d.Board.Hand = append(d.Board.Hand, card)
        d.Board.Piles[i] = d.Board.Piles[i][0:j]
      }
    }
  }
  // 
  {
    // 3. Hand -> Waste
    currHand := d.Board.Hand
    d.Board.Waste = append(d.Board.Waste, d.Board.Hand...)
    if len(d.Board.Stock) > 3 {
      d.Board.Hand = d.Board.Stock[0:3]
      d.Board.Stock = d.Board.Stock[3:]
    } else {
      d.Board.Hand = d.Board.Stock
      d.Board.Stock = nil
    }
    // Recursion.
    newMove := append(moves,
      fmt.Sprintf("refresh hand (%s) -> waste", currHand))
    ret := d.Solve(newMove, visited)
    if ret {
      return true
    }

    // Reset
    d.Board.Waste = d.Board.Waste[0:len(d.Board.Waste)-len(currHand)]
    d.Board.Stock = append(d.Board.Hand, d.Board.Stock...) 
    d.Board.Hand = currHand

  }

  // 4. Stack -> Pile
  for _, cardType := range AllCardTypes {
    card, ok := d.Board.Stack[cardType]
    if !ok || card.Number < 2 {
      continue
    }

    for i, pile := range d.Board.Piles {
      j := len(pile)
      if d.Board.CanMoveToPile(card, i) {
        newCard := GameCard {
          Card: &card,
        }
        newCard.Reveal()
        d.Board.Piles[i] = append(d.Board.Piles[i], newCard)
        d.Board.PopStack(cardType)

        // Recursion.
        newMove := append(moves,
          fmt.Sprintf("stack %s -> (%d, %d)", card, i, j))
        ret := d.Solve(newMove, visited)
        if ret {
          return true
        }

        // Reset
        d.Board.Stack[cardType] = card
        d.Board.Piles[i] = d.Board.Piles[i][0:j]
      }
    }
  }

  // 5. Pile -> Pile
  for srcI, srcPile := range d.Board.Piles {
    for srcJ := len(srcPile) - 1; srcJ >= 0; srcJ-- {
      gameCard := d.Board.GetPileCard(srcI, srcJ)
      if !gameCard.IsRevealed() {
        break
      }
      card := NewCard(gameCard.Card.Type, gameCard.Card.Number)

      for i, pile := range d.Board.Piles {
        if i == srcI {
          continue
        }
        j := len(pile)
        if d.Board.CanMoveToPile(card, i) {
          d.Board.MovePiles(srcI, srcJ, i)

          // Update moves
          newMove := append(moves,
            fmt.Sprintf("pile card %s (%d, %d) -> (%d, %d)", card, srcI, srcJ, i, j))

          // Check if a new card is revealed.
          isNewCardRevealed := false
          if srcJ > 0 {
            fmt.Printf("%s\n", newMove)
            prevGameCard, _ := d.Board.GetPileTailCard(srcI)
            if !prevGameCard.IsRevealed() {
              isNewCardRevealed = true
              prevGameCard.Reveal()
              d.Board.Piles[srcI][srcJ - 1] = prevGameCard
            }
          }

          // Recursion.
          ret := d.Solve(newMove, visited)
          if ret {
            return true
          }

          // Reset
          if isNewCardRevealed {
            // unreveal
            prevGameCard, _ := d.Board.GetPileTailCard(srcI)
            prevGameCard.Unreveal()
            d.Board.Piles[srcI][srcJ - 1] = prevGameCard
          }
          d.Board.Piles[srcI] = append(d.Board.Piles[srcI], d.Board.Piles[i][j:]...)
          d.Board.Piles[i] = d.Board.Piles[i][0:j]
        }
      }
    }
  }
  // 6. Pile -> Stack
  for srcI, _ := range d.Board.Piles {
      gameCard, srcJ := d.Board.GetPileTailCard(srcI)
      card := NewCard(gameCard.Card.Type, gameCard.Card.Number)

        if d.Board.SavePileCardToStack(srcI) {
          newMove := append(moves,
            fmt.Sprintf("pile card %s (%d, %d) -> stack", card, srcI, srcJ))

          // Check if a new card is revealed.
          isNewCardRevealed := false
          if srcJ > 0 {
            fmt.Printf("%s\n", newMove)
            prevGameCard, _ := d.Board.GetPileTailCard(srcI)
            if !prevGameCard.IsRevealed() {
              isNewCardRevealed = true
              prevGameCard.Reveal()
              d.Board.Piles[srcI][srcJ - 1] = prevGameCard
            }
          }

          // Recursion.
          ret := d.Solve(newMove, visited)
          if ret {
            return true
          }

          // Reset
          if isNewCardRevealed {
            // unreveal
            prevGameCard, _ := d.Board.GetPileTailCard(srcI)
            prevGameCard.Unreveal()
            d.Board.Piles[srcI][srcJ - 1] = prevGameCard
          }
          d.Board.PopStack(card.Type)
          d.Board.Piles[srcI] = append(d.Board.Piles[srcI], gameCard)
        }
  }

  // 7. Refresh Stock
  if len(d.Board.Stock) == 0 && len(d.Board.Hand) == 0 && len(d.Board.Waste) > 0 {
    currWaste := d.Board.Waste
    d.Board.Stock = d.Board.Waste
    d.Board.Waste = nil

    // Recursion.
    newMove := append(moves,
      fmt.Sprintf("refresh waste (%s) -> stock", currWaste))
    ret := d.Solve(newMove, visited)
    if ret {
      return true
    }

    // Reset
    d.Board.Waste = currWaste
    d.Board.Stock = nil
  }
  // 8. Next Hand

  return false
}



