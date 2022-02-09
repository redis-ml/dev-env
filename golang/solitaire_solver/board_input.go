package solitaire_solver

import (
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func (b *Board) UpdateCardFromInput(scanner *bufio.Scanner) {
  scanner.Scan()
  line := scanner.Text()
  fmt.Printf("got input: %s.\n", line)

  l := strings.Split(line, ",")
  if len(l) < 2 {
    fmt.Printf("[WARN] invalid input %s\n", line)
    return
  }

  if l[0] == "stock" {
    // Special logic for updating the Stock
    var stock []Card
    for i, s := range l {
      if i == 0 {
        continue
      }
      stock = append(stock, CardFromString(s))
    }
    b.Waste = nil
    b.Stock = stock

    return
  }

  // Otherwise, update a pile card.
  x := Must(strconv.Atoi(strings.TrimSpace(l[0])))
  y := Must(strconv.Atoi(strings.TrimSpace(l[1])))

  b.SetPileCard(x, y, CardFromString(l[2]))
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

func (b *Board) SetPileCard(x int, y int, card Card) bool {
  gameCard := b.Piles[x][y]
  if gameCard.Card != nil {
    fmt.Printf("[ERROR], (%d, %d) is already assigned to %s, While you're trying to assign to %s\n", x, y, gameCard.Card, card)
    return false
  }
  gameCard.Card = &card
  fmt.Printf("(%d, %d) -> %s\n", x, y, card)
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
