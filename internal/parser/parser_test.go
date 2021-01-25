package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := map[string]struct {
		input           string
		expectedResults []Result
		expectedCommand string
		expectedErr     string
	}{
		"Assigment": {
			input: `*/15 0 1,15 * 1-5 /usr/bin/find`,
			expectedResults: []Result{
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
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Default": {
			input: `0 0 1 1 0 /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{0},
				},
				{
					Label: "hour",
					Items: []int{0},
				},
				{
					Label: "day of month",
					Items: []int{1},
				},
				{
					Label: "month",
					Items: []int{1},
				},
				{
					Label: "day of week",
					Items: []int{0},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Every second": {
			input: `* * * * * /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
				},
				{
					Label: "hour",
					Items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				},
				{
					Label: "day of month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				},
				{
					Label: "month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				},
				{
					Label: "day of week",
					Items: []int{0, 1, 2, 3, 4, 5, 6},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Every second question mark": {
			input: `* * ? * ? /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
				},
				{
					Label: "hour",
					Items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				},
				{
					Label: "day of month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				},
				{
					Label: "month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				},
				{
					Label: "day of week",
					Items: []int{0, 1, 2, 3, 4, 5, 6},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Steps": {
			input: `*/2 5/3 1/1 1/5 2/9 /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58},
				},
				{
					Label: "hour",
					Items: []int{5, 8, 11, 14, 17, 20, 23},
				},
				{
					Label: "day of month",
					Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				},
				{
					Label: "month",
					Items: []int{1, 6, 11},
				},
				{
					Label: "day of week",
					Items: []int{2},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Ranges": {
			input: `1-2 0-23 1-10/2 11-12 0-6 /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{1, 2},
				},
				{
					Label: "hour",
					Items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
				},
				{
					Label: "day of month",
					Items: []int{1, 3, 5, 7, 9},
				},
				{
					Label: "month",
					Items: []int{11, 12},
				},
				{
					Label: "day of week",
					Items: []int{0, 1, 2, 3, 4, 5, 6},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Joins": {
			input: `1,2 1-3,5-7 1-3,2-4 1,2,1 *,5 /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{1, 2},
				},
				{
					Label: "hour",
					Items: []int{1, 2, 3, 5, 6, 7},
				},
				{
					Label: "day of month",
					Items: []int{1, 2, 3, 4},
				},
				{
					Label: "month",
					Items: []int{1, 2},
				},
				{
					Label: "day of week",
					Items: []int{0, 1, 2, 3, 4, 5, 6},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},
		"Strings": {
			input: `0 0 1 FeB tUe /usr/bin/find`,
			expectedResults: []Result{
				{
					Label: "minute",
					Items: []int{0},
				},
				{
					Label: "hour",
					Items: []int{0},
				},
				{
					Label: "day of month",
					Items: []int{1},
				},
				{
					Label: "month",
					Items: []int{2},
				},
				{
					Label: "day of week",
					Items: []int{2},
				},
			},
			expectedCommand: `/usr/bin/find`,
			expectedErr:     "",
		},

		// Errors

		"Short input": {
			input:       `0 0 /usr/bin/find`,
			expectedErr: "invalid number of sections",
		},
		"Invalid input": {
			input:       `a b * * * /usr/bin/find`,
			expectedErr: "`a` does not match expected pattern `^[\\d|\\*|\\-|,|/]+$` in `minute`",
		},
		// TODO: validation for each slot
		"Out of range end": {
			input:       `0-99 * * * * /usr/bin/find`,
			expectedErr: "invalid range end in `minute`",
		},
		"Out of range start": {
			input:       `* * 0-1 * * /usr/bin/find`,
			expectedErr: "invalid range start in `day of month`",
		},
		"Out of range single": {
			input:       `* * 0 * * /usr/bin/find`,
			expectedErr: "item `0` out of range in `day of month`",
		},
		"Invalid range": {
			input:       `2-1 * * * * /usr/bin/find`,
			expectedErr: "invalid range, start `2` > end `1` in `minute`",
		},
		"Invalid step": {
			input:       `1/0 * * * * /usr/bin/find`,
			expectedErr: "`invalid step` in `minute`",
		},

		// TODO: many other cases...
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			results, command, err := Parse(testCase.input, Slots)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if command != testCase.expectedCommand {
				t.Errorf("expected command: %v\nbut got: %v", testCase.expectedCommand, command)
				return
			}

			if !reflect.DeepEqual(results, testCase.expectedResults) {
				t.Errorf("expected results: %v\nbut got: %v", testCase.expectedResults, results)
				return
			}
		})
	}
}

func TestNormalizeValues(t *testing.T) {
	testCases := map[string]struct {
		section  string
		slot     Slot
		expected string
	}{
		"No replacer": {
			section:  "123",
			slot:     Slot{},
			expected: "123",
		},
		"Replacer": {
			section: "old string",
			slot: Slot{
				Replacer: strings.NewReplacer("old", "new"),
			},
			expected: "new string",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := normalizeValues(testCase.section, testCase.slot)

			if got != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, got)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	testCases := map[string]struct {
		section     string
		slot        Slot
		expectedErr string
	}{
		"Ok": {
			section: "123",
			slot: Slot{
				Label:           "test",
				ValidCharacters: `^[\d]+$`,
			},
			expectedErr: "",
		},
		"Error": {
			section: "abc",
			slot: Slot{
				Label:           "test",
				ValidCharacters: `^[\d]+$`,
			},
			expectedErr: "`abc` does not match expected pattern `^[\\d]+$` in `test`",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate(testCase.section, testCase.slot)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

			} else {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
					return
				}
			}
		})
	}
}

func TestNormalizeCharacters(t *testing.T) {
	testCases := map[string]struct {
		section  string
		slot     Slot
		expected string
	}{
		"No special char": {
			section: "123",
			slot: Slot{
				Min: 1,
				Max: 9,
			},
			expected: "123",
		},
		"Asterisk": {
			section: "0 * 1",
			slot: Slot{
				Min: 1,
				Max: 9,
			},
			expected: "0 1-9 1",
		},
		"Question mark": {
			section: "0 ? 1",
			slot: Slot{
				Min: 1,
				Max: 9,
			},
			expected: "0 1-9 1",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := normalizeCharacters(testCase.section, testCase.slot)

			if got != testCase.expected {
				t.Errorf("expected: %v\nbut got: %v", testCase.expected, got)
			}
		})
	}
}

func TestParseSection(t *testing.T) {
	testCases := map[string]struct {
		sections        []string
		slots           []Slot
		expectedResults []Result
		expectedErr     string
	}{
		"Normal": {
			sections: []string{"1", "2-3", "1,3"},
			slots: []Slot{
				{
					Label: "first",
					Min:   0,
					Max:   3,
				},
				{
					Label: "second",
					Min:   0,
					Max:   3,
				},
				{
					Label: "third",
					Min:   0,
					Max:   3,
				},
			},
			expectedResults: []Result{
				{
					Label: "first",
					Items: []int{1},
				},
				{
					Label: "second",
					Items: []int{2, 3},
				},
				{
					Label: "third",
					Items: []int{1, 3},
				},
			},
			expectedErr: "",
		},
		// TODO: more cases
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			results, err := parseSections(testCase.sections, testCase.slots)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if !reflect.DeepEqual(results, testCase.expectedResults) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedResults, results)
			}
		})
	}
}

func TestParseStep(t *testing.T) {
	testCases := map[string]struct {
		item         string
		expectedStep int
		expectedUnit string
		expectedErr  string
	}{
		"Normal": {
			item:         "1/2",
			expectedStep: 2,
			expectedUnit: "1",
			expectedErr:  "",
		},
		// TODO: more cases
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			step, unit, err := parseStep(testCase.item)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if step != testCase.expectedStep {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedStep, step)
			}

			if unit != testCase.expectedUnit {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedUnit, unit)
			}
		})
	}
}

func TestParseRange(t *testing.T) {
	testCases := map[string]struct {
		item          string
		slot          Slot
		step          int
		expectedItems []int
		expectedErr   string
	}{
		"Normal": {
			item: "1-6",
			slot: Slot{
				Max: 9,
			},
			step:          2,
			expectedItems: []int{1, 3, 5},
			expectedErr:   "",
		},
		// TODO: more cases
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			items, err := parseRange(testCase.item, testCase.slot, testCase.step)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if !reflect.DeepEqual(items, testCase.expectedItems) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedItems, items)
			}
		})
	}
}

func TestParseSingle(t *testing.T) {
	testCases := map[string]struct {
		item          string
		slot          Slot
		step          int
		expectedItems []int
		expectedErr   string
	}{
		"Normal": {
			item: "7",
			slot: Slot{
				Min: 0,
				Max: 9,
			},
			step:          0,
			expectedItems: []int{7},
			expectedErr:   "",
		},
		"With step": {
			item: "1",
			slot: Slot{
				Min: 0,
				Max: 9,
			},
			step:          2,
			expectedItems: []int{1, 3, 5, 7, 9},
			expectedErr:   "",
		},
		"Out of range": {
			item: "9",
			slot: Slot{
				Label: "test",
				Min:   0,
				Max:   1,
			},
			step:          2,
			expectedItems: []int{},
			expectedErr:   "item `9` out of range in `test`",
		},
		// TODO: more cases
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			items, err := parseSingle(testCase.item, testCase.slot, testCase.step)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if !reflect.DeepEqual(items, testCase.expectedItems) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedItems, items)
			}
		})
	}
}

func TestParseJoins(t *testing.T) {
	testCases := map[string]struct {
		section       string
		slot          Slot
		expectedItems []int
		expectedErr   string
	}{
		"Normal": {
			section: "1,3",
			slot: Slot{
				Label: "test",
				Min:   1,
				Max:   9,
			},
			expectedItems: []int{1, 3},
			expectedErr:   "",
		},
		// TODO: more cases
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			items, err := parseJoins(testCase.section, testCase.slot)

			if testCase.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error: %v\nbut got nothing", testCase.expectedErr)
					return
				}

				if err.Error() != testCase.expectedErr {
					t.Errorf("expected error: %v\nbut got: %v", testCase.expectedErr, err)
					return
				}

				// Ignore other results
				return
			}

			if testCase.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if !reflect.DeepEqual(items, testCase.expectedItems) {
				t.Errorf("expected: %v\nbut got: %v", testCase.expectedItems, items)
			}
		})
	}
}
