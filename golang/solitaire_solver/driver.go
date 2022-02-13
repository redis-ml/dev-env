package solitaire_solver

import (
  "fmt"
  "bufio"
  "strings"
)

type Driver struct {
  Scanner *bufio.Scanner
  Board *Board
  debug bool
  counter int
}

func (d *Driver) Solve(
  moves []string,
  visited map[string]bool,
) bool {
  b := d.Board

  b.CurrentMoves = moves

  d.counter++
  debug := d.debug || d.counter % 100 == 0
  if debug {
    fmt.Println("#################")
    printMoves(moves)
    fmt.Println("#####")
    b.Print()
    fmt.Println("#################")
  }

  status := b.String()
  if visited[status] {
    if debug {
      fmt.Printf("duplicated status, return false..")
    }
    return false
  }
  visited[status] = true

  // Termination condition
  if b.Done() {
    printMoves(moves)
    return true
  }

  // 1. Pile -> Stack
  for i := range b.Piles {
    if d.SolverWrapper(func() (string, UndoFunc, bool) {
      return b.PileToStack(i)
    }, moves, visited) {
      return true
    }
  }

  // 2. Waste -> Stack
  if d.SolverWrapper(func() (string, UndoFunc, bool) {
    return b.WasteToStack()
  }, moves, visited) {
    return true
  }

  // 3. Pile -> Pile
  if d.TryPileToPileMoves(moves, visited) {
    return true
  }

  // 4. Waste <-> Stock
  if d.SolverWrapper(func() (string, UndoFunc, bool) {
    return b.CycleBetweenStockAndWaste()
  }, moves, visited) {
    return true
  }

  // 5. Waste -> Pile
  for i := range b.Piles {
    if d.SolverWrapper(func() (string, UndoFunc, bool) {
      return b.WasteToPile(i)
    }, moves, visited) {
      return true
    }
  }

  // 6. Stack -> Pile
  for _, cardType := range AllCardTypes {
    if d.SolverWrapper(func() (string, UndoFunc, bool) {
      return b.StackToPile(cardType)
    }, moves, visited) {
      return true
    }
  }

  return false
}

type OpFunc func()(string, UndoFunc, bool)

func (d *Driver) SolverWrapper(opFunc OpFunc, moves []string, visited map[string]bool) bool {
  if move, undo, done := opFunc(); done {
    if d.Solve(append(moves, move), visited) {
      return true
    }
    undo()
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
        if d.SolverWrapper(func() (string, UndoFunc, bool) {
          return b.PileToPile(srcI, srcJ, i)
        }, moves, visited) {
          return true
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
