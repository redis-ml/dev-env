package solitaire_solver_test

import (
  "testing"
)

func TestPileToStack(t *testing.T) {
  b := GetBaseBoard(t)

  defaultBoard = []string {
    "stock,s11,h8,d10,h5,d1,h6,c11,c10,h3,d12,h2,c2,h9,h1,c3,d9,d4",
    "waste,d2,c7,c12,s12,d13,d6,s10",
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
  for _, l := range defaultBoard {
    b.UpdateCardByString(l)
  }
  b.Print()
  originalStatus := b.String()

  nonMovingPiles := []int{0, 1, 2, 3, 5, 6}
  for _, i := range nonMovingPiles {
    move, undo, done := b.PileToStack(i)
    b.Print()
    if b.String() != originalStatus {
      t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, originalStatus)
    }
    if done {
      t.Fatalf("done='true', expecting 'false'")
    }
    if move != "" {
      t.Fatalf("move='%s', expecting ''", move)
    }
    if undo == nil {
      t.Fatal("'undo' function is nil")
    }
  }

  movingPiles := map[int][]string{
    4: []string{
      "pile card ♣A (4, 4) -> stack",
      "x,x,x,♣A,;♠J♥8♦10♥5♦A♥6♣J♣10♥3♦Q♥2♣2♥9♥A♣3♦9♦4;♦2♣7♣Q♠Q♦K♦6♠10;♠6;x♥J;xx♣K;xxx♦J;x♠8♥4♥Q;xxxxx♠K;xxxxxx♠2;",
    },
  }
  for pileIdx, expected := range movingPiles {
    expectedMove := expected[0]
    expectedStatus := expected[1]

    originalPileLen := len(b.Piles[pileIdx])
    move, undo, done := b.PileToStack(pileIdx)
    b.Print()
    if b.String() != expectedStatus {
      t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, expectedStatus)
    }
    if !done {
      t.Fatalf("done is 'false'")
    }

    if len(b.Piles[pileIdx]) != originalPileLen - 1 {
      t.Fatalf("pile[%d] not updated: %s", pileIdx, b.Piles[pileIdx])
    }
    if move != expectedMove {
      t.Fatalf("move = %s, expecting '%s'", move, expectedMove)
    }

    // Test Undo.
    undo()
    b.Print()
    status := b.String()
    if status != originalStatus {
      t.FailNow()
    }
  }
}

func TestPileToPile(t *testing.T) {
  b := GetBaseBoard(t)

  defaultBoard = []string {
    "stock,s11,h8,d10,h5,d1,h6,c11,c10,h3,d12,h2,c2,h9,h1,c3,d4",
    "waste,c7,c12,s12,d13,d6",
    "0,0,s6",
    "1,1,h11",
    "2,2,c13",
    "3,3,s13",
    "4,4,c1",
    "5,3,d11",
    "5,4,s10",
    "5,5,d9",
    "6,6,d2",
    "4,3,h12",
    "4,2,h4",
    "4,1,s8",
  }
  for _, l := range defaultBoard {
    b.UpdateCardByString(l)
  }
  tmp := b.Piles[4][3]
  tmp.Unreveal()
  b.Piles[4][3] = tmp
  b.Piles[4][2].Unreveal()
  b.Piles[4][1].Unreveal()

  b.Print()
  originalStatus := b.String()

  nonMovingPiles := [][]int{
    []int{5, 4, 1},
  }
  for _, testCase := range nonMovingPiles {
    move, undo, done := b.PileToPile(testCase[0], testCase[1], testCase[2])
    b.Print()
    if done {
      t.Fatalf("done='true', expecting 'false'")
    }
    if b.String() != originalStatus {
      t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, originalStatus)
    }
    if move != "" {
      t.Fatalf("move='%s', expecting ''", move)
    }
    if undo == nil {
      t.Fatal("'undo' function is nil")
    }
  }

  type goodCase struct {
    x, y int
    dst int
    expectedStatus string
    expectedMove string
  }
  movingPiles := []goodCase {
    goodCase {
      x: 4,
      y: 4,
      dst: 6,
      expectedMove: "pile card ♣A (4, 4) -> (6, 7)",
      expectedStatus: "x,x,x,x,;♠J♥8♦10♥5♦A♥6♣J♣10♥3♦Q♥2♣2♥9♥A♣3♦4;♣7♣Q♠Q♦K♦6;♠6;x♥J;xx♣K;xxx♠K;xxx♥Q;xxx♦J♠10♦9;xxxxxx♦2♣A;",
    },
  }
  for _, testCase := range movingPiles {
    expectedMove := testCase.expectedMove
    expectedStatus := testCase.expectedStatus

    move, undo, done := b.PileToPile(testCase.x, testCase.y, testCase.dst)
    b.Print()
    if !done {
      t.Fatalf("done is 'false'")
    }
    if b.String() != expectedStatus {
      t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, expectedStatus)
    }

    if move != expectedMove {
      t.Fatalf("move = %s, expecting '%s'", move, expectedMove)
    }

    // Test Undo.
    undo()
    b.Print()
    status := b.String()
    if status != originalStatus {
      t.Fatal("status doesn't match")
      t.FailNow()
    }
  }
}
