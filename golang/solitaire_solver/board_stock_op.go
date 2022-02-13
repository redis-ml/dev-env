package solitaire_solver

import (
  "fmt"
)

func (b *Board) CycleBetweenStockAndWaste() (move string, undo UndoFunc, done bool) {
  move, undo, done = b.StockToWaste()
  if done {
    return
  }
  return b.RestartStockFromWaste()
}

func (b *Board) RestartStockFromWaste() (string, UndoFunc, bool) {
  if len(b.Stock) == 0 && len(b.Waste) > 0 {
    currWaste := b.Waste
    b.Stock = b.Waste
    b.Waste = nil

    // Recursion.
    move := fmt.Sprintf("refresh stock <- waste (%s)", currWaste)

    // Reset
    undo := func() {
      b.Waste = currWaste
      b.Stock = nil
    }
    return move, undo, true
  }
  return "", func(){}, false
}

func (b *Board) StockToWaste() (string, UndoFunc, bool) {
  if len(b.Stock) > 0 {
    var tmp []Card
    if len(b.Stock) > 3 {
      tmp = b.Stock[0:3]
      b.Stock = b.Stock[3:]
    } else {
      tmp = b.Stock
      b.Stock = nil
    }
    b.Waste = append(b.Waste, tmp...)

    move := fmt.Sprintf("stock (%s) -> waste", tmp)
    undo := func() {
      b.Waste = b.Waste[0:len(b.Waste)-len(tmp)]
      b.Stock = append(tmp, b.Stock...) 
    }
    return move, undo, true
  }
  return "", func(){}, false
}
