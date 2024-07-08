package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct { // custom error
	When time.Time
	What string
}

func (e *MyError) Error() string { // hata mesajını simgeleyen fonksiyon. error handlingteki error mesajı gibi düşünebiliriz.
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42")
	if err != nil { // err == nil ise hata yok demektir
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}
