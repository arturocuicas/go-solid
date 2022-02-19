package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var movementCount = 0

type Balance struct {
	movement []string
}

func (b *Balance) AddMovement(text string) int {
	movementCount++
	fmt.Sprintf("%d: %s", movementCount, text)
	b.movement = append(b.movement, text)
	return movementCount
}

func (b *Balance) String() string {
	return strings.Join(b.movement, "\n")
}

func (b *Balance) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(b.String()), 644)
}

var LineSeparator = "\n"

func SaveToFile(b *Balance, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(b.movement, LineSeparator)), 644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(b *Balance, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(b.movement, p.lineSeparator)), 644)
}

func main() {
	b := Balance{}
	b.AddMovement("Buy Groceries")
	b.AddMovement("Buy ")
	fmt.Println(b.String())

	SaveToFile(&b, "balance.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&b, "balance_2.txt")
}
