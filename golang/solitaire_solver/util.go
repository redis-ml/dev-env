package solitaire_solver

func Must(x int, e error) int {
  if e != nil {
    panic(e)
  }
  return x
}

