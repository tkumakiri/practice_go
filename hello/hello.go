package main

import (
	"fmt"
	"github.com/tkumakiri/practice_go/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}