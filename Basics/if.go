package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h := t.Hour()
	if h < 12 {
		fmt.Println("Good morning!")
	} else if h < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
}
