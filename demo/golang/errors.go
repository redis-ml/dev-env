package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	val string
}

func (e MyErr) Error() string {
	return e.val
}

func main() {
	var err error
	err = &MyErr{}
	var err2 error
	err2 = MyErr{}

	err3 := fmt.Errorf("abc %w", err)

	fmt.Printf("err is Err2: %v\n", errors.As(err, &err2))
	fmt.Printf("err is Err: %v\n", errors.As(err, &err))
	fmt.Printf("err2 is Err: %v\n", errors.As(err2, &err))
	fmt.Printf("err2 is Err2: %v\n", errors.As(err2, &err2))
	fmt.Printf("err3 is Err: %v\n", errors.As(err3, &err))
	fmt.Printf("err3 is Err2: %v\n", errors.As(err3, &err2))
	fmt.Println("vim-go")
}
