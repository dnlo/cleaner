Package for cleaning strings

[![Go Report Card](https://goreportcard.com/badge/github.com/dnlo/cleaner)](https://goreportcard.com/report/github.com/dnlo/cleaner)

Docs: https://godoc.org/github.com/dnlo/cleaner

Example:

```
package main

import (
	"fmt"
	"github.com/dnlo/cleaner"
)

var sellers = []string{
	"Seller: 3*3whiteangel3*3(1,175) 100%",
	"Seller: 80vann(545) 100%",
	"Seller: accessman4(1,010) 97.3%",
	"Seller: a_clothingdipity(1,403) 100%  View seller's store: a_clothingdipity",
	"Seller: andrewandania(11,405) 100%  View seller's store: andrewandania",
}

type seller struct {
	Name     string
	Sales    string
	Feedback string
}

func extract(s string) seller {
	getName := cleaner.Clean(
		cleaner.LastAfter("Seller: "),
		cleaner.FirstBefore(`(`))
	getSales := cleaner.Clean(
		cleaner.LastAfter("("),
		cleaner.FirstBefore(")"),
		cleaner.Delete(","))
	getFeedback := cleaner.Clean(
		cleaner.LastAfter(") "),
		cleaner.Delete("%"),
		cleaner.Extract(`[0-9\.]*`))

	return seller{
		Name:     getName(s),
		Sales:    getSales(s),
		Feedback: getFeedback(s),
	}
}

func main() {
	for _, v := range sellers {
		fmt.Println(extract(v))
	}
}
```