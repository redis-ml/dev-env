package main

import (
  "bufio"
  "fmt"
  "os"

  "github.com/redisliu/dev-env/golang/solitaire_solver"
)

func main() {
  b := solitaire_solver.NewBoard()
  fmt.Println("Hello Worldle!")
  scanner := bufio.NewScanner(os.Stdin)

  b.Print()
  for {
    for b.HasPendingCard() {
      b.UpdateCardFromInput(scanner)
      b.Print()
    }
    b.HandleCommand(scanner)
  }
}
