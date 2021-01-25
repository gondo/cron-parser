package parser

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected []int
	}{
		"Nothing removed": {
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		"One removed": {
			input:    []int{1, 2, 1},
			expected: []int{1, 2},
		},
		"Only duplicates": {
			input:    []int{1, 1, 1},
			expected: []int{1},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res := removeDuplicates(testCase.input)

			if !reflect.DeepEqual(res, testCase.expected) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, res)
			}
		})
	}
}

func TestCleanInput(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"Nothing to clean": {
			input:    "Value abc",
			expected: "Value abc",
		},
		"Clean": {
			input:    " Value abc  ",
			expected: "Value abc",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			out := cleanInput(testCase.input)

			if out != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, out)
			}
		})
	}
}

func TestSeparateCommand(t *testing.T) {
	testCases := map[string]struct {
		sections         []string
		number           int
		expectedSections []string
		expectedCommand  string
	}{
		"Normal": {
			sections:         []string{"a", "b", "c"},
			number:           3,
			expectedSections: []string{"a", "b"},
			expectedCommand:  "c",
		},
		"More sections": {
			sections:         []string{"a", "b", "c"},
			number:           2,
			expectedSections: []string{"a"},
			expectedCommand:  "b",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			sections, command := separateCommand(testCase.sections, testCase.number)

			if !reflect.DeepEqual(sections, testCase.expectedSections) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedSections, sections)
			}

			if command != testCase.expectedCommand {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedCommand, command)
			}
		})
	}
}

func TestIsRange(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected bool
	}{
		"Normal": {
			input:    "1-9",
			expected: true,
		},
		"More ranges": {
			input:    "1-9-a",
			expected: true,
		},
		"Just range": {
			input:    "-",
			expected: true,
		},
		"No range": {
			input:    "123",
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			is := isRange(testCase.input)

			if is != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, is)
			}
		})
	}
}
