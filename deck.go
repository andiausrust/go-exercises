package main

import "fmt"

type deck []string

func (d deck) printDeck() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}
