package output

import (
	"github.com/gondo/deliveroo-interview/internal/parser"
	"testing"
)

func TestTable(t *testing.T) {
	testCases := map[string]struct {
		results  []parser.Result
		expected string
	}{
		"Normal": {
			results: []parser.Result{
				{
					Label: "minute",
					Items: []int{0, 15, 30, 45},
				},
				{
					Label: "hour",
					Items: []int{0},
				},
				{
					Label: "day of month",
					Items: []int{1, 15},
				},
				{
					Label: "month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				},
				{
					Label: "day of week",
					Items: []int{1, 2, 3, 4, 5},
				},
			},
			expected: `minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5`,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tab := Table(testCase.results)

			if tab != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, tab)
			}
		})
	}
}

func TestSliceToStr(t *testing.T) {
	testCases := map[string]struct {
		slice    []int
		sep      string
		expected string
	}{
		"Normal": {
			slice:    []int{1, 2, 3},
			sep:      ", ",
			expected: "1, 2, 3",
		},
		"Empty slice": {
			slice:    []int{},
			sep:      ", ",
			expected: "",
		},
		"Empty separator": {
			slice:    []int{1, 2, 3},
			sep:      "",
			expected: "123",
		},
		"Big separator": {
			slice:    []int{1, 2, 3},
			sep:      "-_-_-",
			expected: "1-_-_-2-_-_-3",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			s := SliceToStr(testCase.slice, testCase.sep)

			if s != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, s)
			}
		})
	}
}

func TestRow(t *testing.T) {
	testCases := map[string]struct {
		label    string
		value    string
		expected string
	}{
		"Normal": {
			label:    "Row Label",
			value:    "Row Value",
			expected: "Row Label     Row Value",
		},
		"Long label": {
			label:    "abc abc abc abc abc abc",
			value:    "value",
			expected: "abc abc abc abc abc abc value",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			r := Row(testCase.label, testCase.value)

			if r != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, r)
			}
		})
	}
}
