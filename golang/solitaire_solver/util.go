package solitaire_solver

func Must(x int, e error) int {
  if e != nil {
    panic(e)
  }
  return x
}

func CardsFromStringArray(l []string) []Card {
    var cards []Card
    for _, s := range l {
      cards = append(cards, CardFromString(s))
    }
    return cards
}

