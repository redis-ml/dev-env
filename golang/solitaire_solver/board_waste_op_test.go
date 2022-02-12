package solitaire_solver_test

import (
  "testing"
)

func TestWasteToStack(t *testing.T) {
  b := GetBaseBoard(t)

  defaultBoard = []string {
    "stock,s11,h8,d10,h5,d1,h6,c11,c10,h3,d12,h2,c2,h9,h1,c3,d9,d4",
    "waste,d2,c7,c12,s12,d13,d6,s10,s1",
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
  originalWasteLen := len(b.Waste)

  move, undo, done := b.WasteToStack()
  b.Print()
  expectedStatus := "♠A,x,x,x,;♠J♥8♦10♥5♦A♥6♣J♣10♥3♦Q♥2♣2♥9♥A♣3♦9♦4;♦2♣7♣Q♠Q♦K♦6♠10;♠6;x♥J;xx♣K;xxx♦J;x♠8♥4♥Q♣A;xxxxx♠K;xxxxxx♠2;"
  if b.String() != expectedStatus {
    t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, expectedStatus)
  }
  if !done {
    t.Fatalf("done is 'false'")
  }

  if len(b.Waste) != originalWasteLen - 1 {
    t.Fatalf("waste not updated: %s", b.Waste)
  }
  expectedMove := "waste ♠A -> stack"
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

func TestWasteToPile(t *testing.T) {
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
  originalWasteLen := len(b.Waste)

  // Should not move.
  move, undo, done := b.WasteToPile(0)
  b.Print()
  if done {
    t.Fatal("done = true, expecting 'false'")
  }
  if b.String() != originalStatus {
    t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", originalStatus, b)
  }

  // Should move.
  move, undo, done = b.WasteToPile(1)
  b.Print()
  expectedStatus := "x,x,x,x,;♠J♥8♦10♥5♦A♥6♣J♣10♥3♦Q♥2♣2♥9♥A♣3♦9♦4;♦2♣7♣Q♠Q♦K♦6;♠6;x♥J♠10;xx♣K;xxx♦J;x♠8♥4♥Q♣A;xxxxx♠K;xxxxxx♠2;"
  if b.String() != expectedStatus {
    t.Fatalf("Comparing new status:\nobtained:\n%s\nexpected:\n%s", b, expectedStatus)
  }
  if !done {
    t.Fatalf("done is 'false'")
  }

  if len(b.Waste) != originalWasteLen - 1 {
    t.Fatalf("waste not updated: %s", b.Waste)
  }
  expectedMove := "waste ♠10 -> (1, 2)"
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

