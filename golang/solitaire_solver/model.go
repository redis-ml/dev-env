package solitaire_solver

import (
  "strings"
  "strconv"
)

type CardType string

const (
  CardType_Spade CardType = "♠"
  CardType_Heart CardType = "♥"
  CardType_Diamond CardType = "♦"
  CardType_Club CardType = "♣"
)

var (
  AllCardTypes []CardType = []CardType {
    CardType_Spade,
    CardType_Heart,
    CardType_Diamond,
    CardType_Club,
  }

  CardFaces = []string {
    "A",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "10",
    "J",
    "Q",
    "K",
  }
)

type Card struct {
  Type CardType
  Number int
}

func (c Card) String() string {
  return string(c.Type) + CardFaces[c.Number]
}

func (c *Card) IsK() bool {
  return c.Number == len(CardFaces) - 1
}

func (c *Card) Color() string {
  switch c.Type {
  case CardType_Spade:
    fallthrough
  case CardType_Club:
    return "black"
  case CardType_Heart:
    fallthrough
  case CardType_Diamond:
    fallthrough
  default:
    return "red"
  }
}

type GameCard struct {
  Card *Card
}

func (gc GameCard) String() string {
  if gc.Card == nil {
    return "-X-"
  }
  return gc.Card.String()
}

func NewCard(cardType CardType, number int) Card {
  return Card {
    Type: cardType,
    Number: number,
  }
}

func GetFullDeck() []Card {
  ret := make([]Card, 0, len(CardFaces) * len(AllCardTypes))
  for _, t := range AllCardTypes {
    for number := 0; number < len(CardFaces); number ++ {
      ret = append(ret, NewCard(t, number))
    }
  }
  return ret
}

var cardTypeMap map[string]CardType = map[string]CardType {
  "c": CardType_Club,
  "s": CardType_Spade,
  "h": CardType_Heart,
  "d": CardType_Diamond,
}

func CardTypeFromString(line string) CardType {
  cardType := strings.ToLower(strings.ToLower(line))
  // TODO: this could be smarter.
  return cardTypeMap[cardType]
}

func CardFromString(line string) Card {
  card := strings.TrimSpace(line)

  face := strings.ToUpper(strings.TrimSpace(card[1:]))
  n := -1
  for i, f := range CardFaces {
    if f == face {
      n = i
      break
    }
  }
  if n == -1 {
    n = Must(strconv.Atoi(card[1:])) - 1
  }

  cardTypeString := card[0:1]
  return NewCard(CardTypeFromString(cardTypeString), n)
}
