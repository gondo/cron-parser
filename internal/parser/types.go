package parser

import (
	"sort"
	"strings"
)

type Slot struct {
	Label           string
	Min             int
	Max             int
	ValidCharacters string
	Replacer        *strings.Replacer
}

type Result struct {
	Label string
	Items []int
}

func (c *Result) AddItems(items []int) {
	items = removeDuplicates(items)
	sort.Ints(items)
	c.Items = items
}
