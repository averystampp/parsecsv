package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Item struct {
	upc   string
	name  string
	price string
}

func main() {

}

func parsefile(filename string) []Item {
	b, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer b.Close()

	content, err := csv.NewReader(b).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	output := []Item{}

	for _, item := range content[1:] {

		newitem := Item{
			upc:   item[0],
			name:  item[1],
			price: item[2],
		}
		output = append(output, newitem)
	}

	return output
}

func compareitems(list1, list2 []Item) {
	for _, newitem := range list1 {
		for _, testitem := range list2 {
			if newitem.upc == testitem.upc {
				fmt.Println(newitem, testitem)
			}
		}
	}
}
