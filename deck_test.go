package main

import "testing"


func TestNewDeck(t *testing.T){// for the purpose of testing the newDeck function from the deck.go file.
	d:=newDeck()
	if len(d) !=48{
		t.Errorf("Expected Deck Length of New Deck  as 48, But got %v",len(d))//%v for the purpose of using Errorf as it required an argument which we want to use.
	}
}