package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}

func (d deck) printDeck() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func deal(d deck, h int) (deck, deck){
	return d[:h], d[h:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(f string) error {
	return ioutil.WriteFile(f, []byte(d.toString()), 0666)
}

func newDeckFromFile(fn string) deck {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		np := r.Intn(len(d) -1)

		d[i], d[np] = d[np], d[i]
	}

}