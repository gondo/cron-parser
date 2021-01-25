package output

import (
	"fmt"
	"github.com/gondo/deliveroo-interview/internal/parser"
	"strings"
)

func Table(results []parser.Result) string {
	var rows []string

	for i := range results {
		res := results[i]
		rows = append(rows, Row(res.Label, SliceToStr(res.Items, " ")))
	}

	return strings.Join(rows, "\n")
}

func SliceToStr(slice []int, sep string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), sep), "[]")
}

func Row(label, value string) string {
	// Hardcoded number of characters for a label
	// 13 + extra space = 14 columns as per assignment
	return fmt.Sprintf("%-13s %s", label, value)
}
