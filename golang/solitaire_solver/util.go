package solitaire_solver

import (
  "fmt"
)

func Must(x int, e error) int {
  if e != nil {
    panic(e)
  }
  return x
}

func CardsFromStringArray(l []string) []Card {
    var cards []Card
    for _, s := range l {
      card, ok := CardFromString(s)
      if !ok {
        panic(fmt.Sprintf("invalid format for card: %s", s))
      }
      cards = append(cards, card)
    }
    return cards
}

