package main

import (
	"fmt"
)

type Card struct {
	Name string
}

func main() {
	var aotd = Card{"Archfiend of the Dross\n"}
	fmt.Print(aotd.cardName())
}

func (card Card) cardName() string { return card.Name }
