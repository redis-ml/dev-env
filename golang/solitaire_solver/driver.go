package solitaire_solver

import (
  "fmt"
  "bufio"
  "strings"
)

type Driver struct {
  Scanner *bufio.Scanner
  Board *Board
}

func (d *Driver) Solve(
  moves []string,
  visited map[string]bool,
) bool {
  b := d.Board

  fmt.Println("#################")
  if len(moves) > 0 {
    fmt.Printf("%s\n", moves[len(moves)-1])
  }
  fmt.Println("#####")
  b.Print()
  fmt.Println("#################")

  status := b.String()
  if visited[status] {
    fmt.Printf("duplicated status, return false..")
    return false
  }
  visited[status] = true

  // Termination condition
  if b.Done() {
    printMoves(moves)
    return true
  }

  // Waste operations
  if d.TryWasteOps(moves, visited) {
    return true
  }

  // 3. Stock -> Waste
  if d.TryStockToWaste(moves, visited) {
    return true
  }

  // 4. Pile -> Stack
  if d.TryPileToStack(moves, visited) {
    return true
  }

  // 5. Pile -> Pile
  if d.TryPileToPileMoves(moves, visited) {
    return true
  }

  // 6. Stack -> Pile
  if d.TryStackToPile(moves, visited) {
    return true
  }

  // 7. Waste -> Stock
  if len(b.Stock) == 0 && len(b.Waste) > 0 {
    currWaste := b.Waste
    b.Stock = b.Waste
    b.Waste = nil

    // Recursion.
    newMove := append(moves,
      fmt.Sprintf("refresh waste (%s) -> stock", currWaste))
    if d.Solve(newMove, visited) {
      return true
    }

    // Reset
    b.Waste = currWaste
    b.Stock = nil
  }

  return false
}

func (d *Driver) TryStockToWaste(moves []string, visited map[string]bool) bool {
    b := d.Board

    if len(b.Stock) == 0 {
      return false
    }

    var tmp []Card
    if len(b.Stock) > 3 {
      tmp = b.Stock[0:3]
      b.Stock = b.Stock[3:]
    } else {
      tmp = b.Stock
      b.Stock = nil
    }
    b.Waste = append(b.Waste, tmp...)

    // Recursion.
    newMove := append(moves,
      fmt.Sprintf("stock (%s) -> waste", tmp))
    ret := d.Solve(newMove, visited)
    if ret {
      return true
    }

    // Reset
    b.Waste = b.Waste[0:len(b.Waste)-len(tmp)]
    b.Stock = append(tmp, b.Stock...) 

    return false
}

func (d *Driver) TryWasteOps(moves []string, visited map[string]bool) bool {
  b := d.Board
  if len(b.Waste) == 0 {
    return false
  }

  // 1. Waste -> Stack
  if move, undo, done := b.WasteToStack(); done {
    // recursion
    if d.Solve(append(moves, move), visited) {
      return true
    }
    undo()
  }

  // 2. Waste -> Pile
  for i := range b.Piles {
    if move, undo, done := b.WasteToPile(i); done {
      // Recursion.
      if d.Solve(append(moves, move), visited) {
        return true
      }
      undo()
    }
  }

  return false
}

func (d *Driver) TryStackToPile(moves []string, visited map[string]bool) bool {
  for _, cardType := range AllCardTypes {
    card, ok := d.Board.Stack[cardType]
    if !ok || card.Number < 2 {
      continue
    }

    for i, pile := range d.Board.Piles {
      j := len(pile)
      if d.Board.CanMoveToPile(card, i) {
        tmp := NewCard(card.Type, card.Number)
        newCard := GameCard {
          Card: &tmp,
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
  return false
}

func (d *Driver) TryPileToStack(moves []string, visited map[string]bool) bool {
  b := d.Board
  for i := range b.Piles {
    if move, undo, done := b.PileToStack(i); done {
      // Recursion.
      ret := d.Solve(append(moves, move), visited)
      if ret {
        return true
      }

      undo()
    }
  }
  return false
}

func (d *Driver) TryPileToPileMoves(moves []string, visited map[string]bool) bool {
  b := d.Board
  for srcI, srcPile := range b.Piles {
    for srcJ := len(srcPile)-1; srcJ >= 0; srcJ-- {
      gameCard := b.GetPileCard(srcI, srcJ)
      if !gameCard.IsRevealed() {
        break
      }
      for i := range b.Piles {
        if move, undo, done := b.PileToPile(srcI, srcJ, i); done {
          if d.Solve(append(moves, move), visited) {
            return true
          }
          undo()
        }
      }
    }
  }
  return false
}


func printMoves(s []string) {
  fmt.Println("#################")
  fmt.Printf("%s\n", strings.Join(s, "\n  "))
  fmt.Println("#################")
}
