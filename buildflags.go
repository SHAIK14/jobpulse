package main

import (
	"flag"
	"fmt"
)

func buildflags() {

	wordptr := flag.String("word", "asif", "a name")
	numptr := flag.Int("age", 25, " a integer of age")
	adultptr := flag.Bool("adult", true, " above 18")

	flag.Parse()

	fmt.Println("word:", *wordptr)
	fmt.Println("age:", *numptr)
	fmt.Println("adult:", *adultptr)

}
