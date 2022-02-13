package solitaire_solver

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (b *Board) InitBoardFromInput() {
	for {
		line := b.UpdateCardFromInput()
		if line == "init_end" {
			b.Print()
			return
		}
		b.Print()
	}
}

func (b *Board) GetInput(prompt string) string {
	fmt.Printf(">>> %s\n>", prompt)
	scanner := b.Scanner
	if !scanner.Scan() {
		panic("EOF")
	}
	line := strings.TrimSpace(scanner.Text())
	fmt.Printf("got input: %s.\n", line)
	return line
}

func (b *Board) GetCardFromInput(prompt string) (Card, bool) {
	line := b.GetInput(prompt)
	return CardFromString(line)
}

func (b *Board) UpdateCardFromInput() string {
	line := b.GetInput("next")
	b.UpdateCardByString(line)
	return line
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

	x, y, card, ok := ParseCardAndPosition(l)
	if !ok {
		fmt.Printf("[WARN] invalid input %s\n", line)
		return
	}

	b.SetPileCard(x, y, card, y == x)
}

func ParseCardAndPositionByString(line string) (x, y int, card Card, ok bool) {
	return ParseCardAndPosition(strings.Split(line, ","))
}

func ParseCardAndPosition(l []string) (x, y int, card Card, ok bool) {
	var err error
	x, err = strconv.Atoi(strings.TrimSpace(l[0]))
	if err != nil {
		return
	}
	y, err = strconv.Atoi(strings.TrimSpace(l[1]))
	if err != nil {
		return
	}

	card, ok = CardFromString(l[2])
	return
}

func (b *Board) HasPendingCard() bool {
	for i, pile := range b.Piles {
		l := len(pile)
		if l == 0 {
			continue
		}
		if pile[l-1].Card == nil {
			fmt.Printf("(%d, %d) should have been revealed !\n", i, l-1)
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
