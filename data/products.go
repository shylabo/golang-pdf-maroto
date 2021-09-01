package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func genereateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	froot := []string{}
	froot = append(froot, f.Name)
	froot = append(froot, f.Description)
	froot = append(froot, fmt.Sprintf("%.2f", f.Price))

	return froot
}

func FruitList(length int) [][]string {
	var fruits [][]string

	for i := 0; i < length; i++ {
		onefruit := genereateFruit()
		fruits = append(fruits, onefruit)
	}

	return fruits
}
