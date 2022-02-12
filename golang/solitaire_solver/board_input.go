package solitaire_solver

import (
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func (b *Board) UpdateCardFromInput() {
  scanner := b.Scanner
  if !scanner.Scan() {
    panic("EOF")
  }
  line := scanner.Text()
  fmt.Printf("got input: %s.\n", line)
  b.UpdateCardByString(line)
}

func (b *Board) UpdateCardByString(line string) {
  l := strings.Split(strings.TrimSpace(line), ",")
  if len(l) < 2 {
    fmt.Printf("[WARN] invalid input %s\n", line)
    return
  }

  switch l[0] {
  case "stock":
    b.Stock = CardsFromStringArray(l[1:len(l)])
    return
  case "waste":
    b.Waste = CardsFromStringArray(l[1:len(l)])
    return
  }

  // Otherwise, update a pile card.
  x := Must(strconv.Atoi(strings.TrimSpace(l[0])))
  y := Must(strconv.Atoi(strings.TrimSpace(l[1])))

  b.SetPileCard(x, y, CardFromString(l[2]), true)
}

func (b *Board) HasPendingCard() bool {
  for i, pile := range b.Piles {
    l := len(pile)
    if l == 0 {
      continue
    }
    if pile[l - 1].Card == nil {
      fmt.Printf("(%d, %d) should have been revealed !\n", i, l - 1)
      return true
    }
  }
  return false
}

func (b *Board) SetPileCard(x int, y int, card Card, revealed bool) bool {
  gameCard := b.Piles[x][y]
  if gameCard.Card != nil {
    fmt.Printf("[ERROR], (%d, %d) is already assigned to %s, While you're trying to assign to %s\n", x, y, gameCard.Card, card)
    return false
  }
  gameCard.Card = &card

  if revealed {
    gameCard.Reveal()
  } else {
    gameCard.Unreveal()
  }

  fmt.Printf("(%d, %d) -> %s, revealed: %v, %v\n", x, y, card, revealed, gameCard.IsRevealed())

  b.Piles[x][y] = gameCard
  return true
}

func (b *Board) HandleCommand(scanner *bufio.Scanner) {
  // TODO: fix this.
  fmt.Println("Please type in the command:")
  if !scanner.Scan() {
    panic("EOF!!")
  }
  line := scanner.Text()
  fmt.Printf("got command: %s.\n", line)
}
