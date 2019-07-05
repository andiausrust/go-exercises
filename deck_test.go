package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("expected first card of Ace of Spaces but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("decktesting")

	if len(ld) != 16 {
		t.Errorf("expected 16 cards, got %v", len(ld))
	}

	os.Remove("_decktesting")
}