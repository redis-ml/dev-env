package solitaire_solver_test

import (
  "testing"
  "bufio"
  "strings"

  "github.com/redisliu/dev-env/golang/solitaire_solver"
)

var (
  defaultBoard = []string {
    "stock,d2,c7,c12,s12,d13,d6,s10,s11,h8,d10,h5,d1,h6,c11,c10,h3,d12,h2,c2,h9,h1,c3,d9,d4",
    "0,0,s6",
    "1,1,h11",
    "2,2,c13",
    "3,3,d11",
    "4,4,c1",
    "5,5,s13",
    "6,6,s2",
    "4,3,h12",
    "4,2,h4",
    "4,1,s8",
  }
)

func GetBaseBoard(t *testing.T) *solitaire_solver.Board {
  scanner := bufio.NewScanner(strings.NewReader(""))
  board := solitaire_solver.NewBoard(scanner)
  return board
}

func TestBaseBoard(t *testing.T) {
  b := GetBaseBoard(t)
  for _, l := range defaultBoard {
    b.UpdateCardByString(l)
  }
  expected := "x,x,x,x,;♦2♣7♣Q♠Q♦K♦6♠10♠J♥8♦10♥5♦A♥6♣J♣10♥3♦Q♥2♣2♥9♥A♣3♦9♦4;;♠6;x♥J;xx♣K;xxx♦J;x♠8♥4♥Q♣A;xxxxx♠K;xxxxxx♠2;"
  if expected != b.String() {
    t.Fatalf("Comparing board status:\nobtained:\n%s\nexpected:\n%s", b, expected)
  }
}
