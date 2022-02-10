package main

import (
  "bufio"
  "fmt"
  "os"

  "github.com/redisliu/dev-env/golang/solitaire_solver"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  b := solitaire_solver.NewBoard(scanner)
  d := &solitaire_solver.Driver {
    Scanner: scanner,
    Board: b,
  }
  fmt.Println("Hello Worldle!")

  b.Print()
  for b.HasPendingCard() {
    b.UpdateCardFromInput()
    b.Print()
  }
  ret := d.Solve(nil, map[string]bool{})
  d.Board.Print()
  fmt.Printf("result: %v\n", ret)
}
