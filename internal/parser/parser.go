package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string, slots []Slot) (results []Result, command string, err error) {
	input = cleanInput(input)
	n := len(slots) + 1 // Number of slots + command
	sections := strings.SplitN(input, " ", n)

	if len(sections) != n {
		return nil, "", errors.New("invalid number of sections")
	}

	sections, command = separateCommand(sections, n)
	results, err = parseSections(sections, slots)
	return results, command, err
}

func parseSections(sections []string, slots []Slot) (results []Result, err error) {
	for i := range sections {
		section := sections[i]
		slot := slots[i]
		result := Result{Label: slot.Label}

		section = normalizeValues(section, slot)

		err := validate(section, slot)
		if nil != err {
			return nil, err
		}

		section = normalizeCharacters(section, slot)

		items, err := parseJoins(section, slot)
		if nil != err {
			return nil, err
		}

		result.AddItems(items)
		results = append(results, result)
	}
	return results, nil
}

// Normalize special values such as: Sun => 0, jan => 1 ...
func normalizeValues(section string, slot Slot) string {
	if slot.Replacer != nil {
		section = strings.ToLower(section)
		section = slot.Replacer.Replace(section)
	}
	return section
}

func validate(section string, slot Slot) error {
	pattern := slot.ValidCharacters
	match, _ := regexp.MatchString(pattern, section)
	if !match {
		return errors.New(fmt.Sprintf("`%s` does not match expected pattern `%s` in `%s`", section, pattern, slot.Label))
	}
	return nil
}

func normalizeCharacters(section string, slot Slot) string {
	sectionRange := fmt.Sprintf("%d-%d", slot.Min, slot.Max)
	section = strings.ReplaceAll(section, "*", sectionRange)
	section = strings.ReplaceAll(section, "?", sectionRange)
	return section
}

// Only one step separator `/` per item is expected. If more present, error will be returned.
func parseStep(item string) (int, string, error) {
	defaultStep := 0
	stepParts := strings.SplitN(item, "/", 2)
	if len(stepParts) != 2 {
		return defaultStep, item, nil
	}

	u := stepParts[0]
	step, err := strconv.Atoi(stepParts[1])
	if nil != err {
		return 0, "", errors.New("invalid step")
	}
	if step == 0 {
		return 0, "", errors.New("invalid step")
	}
	return step, u, nil
}

// Only one range `-` is expected. If more present, error will be returned.
func parseRange(item string, slot Slot, step int) (items []int, err error) {
	rangeParts := strings.SplitN(item, "-", 2)

	start, err := strconv.Atoi(rangeParts[0])
	if nil != err {
		return nil, errors.New(fmt.Sprintf("invalid start in `%s`", slot.Label))
	}
	if start < slot.Min {
		return nil, errors.New(fmt.Sprintf("invalid range start in `%s`", slot.Label))
	}

	end, err := strconv.Atoi(rangeParts[1])
	if nil != err {
		return nil, errors.New(fmt.Sprintf("invalid end in `%s`", slot.Label))
	}
	if end > slot.Max {
		return nil, errors.New(fmt.Sprintf("invalid range end in `%s`", slot.Label))
	}

	if start > end {
		return nil, errors.New(fmt.Sprintf("invalid range, start `%d` > end `%d` in `%s`", start, end, slot.Label))
	}

	rangeStep := 1
	if step > 0 {
		rangeStep = step
	}

	for k := start; k <= end; k += rangeStep {
		items = append(items, k)
	}
	return items, nil
}

func parseSingle(item string, slot Slot, step int) (items []int, err error) {
	start, err := strconv.Atoi(item)
	if nil != err {
		return nil, errors.New(fmt.Sprintf("invalid item `%s` in `%s`", item, slot.Label))
	}
	if start < slot.Min || start > slot.Max {
		return nil, errors.New(fmt.Sprintf("item `%s` out of range in `%s`", item, slot.Label))
	}

	if step > 0 {
		for k := start; k <= slot.Max; k += step {
			items = append(items, k)
		}
	} else {
		items = append(items, start)
	}

	return items, nil
}

func parseJoins(section string, slot Slot) (items []int, err error) {
	units := strings.Split(section, ",")
	for j := range units {
		unit := units[j]

		step, unit, err := parseStep(unit)
		if nil != err {
			return nil, errors.New(fmt.Sprintf("`%s` in `%s`", err, slot.Label))
		}

		if isRange(unit) {
			rangeItems, err := parseRange(unit, slot, step)
			if nil != err {
				return nil, err
			}
			items = append(items, rangeItems...)
			continue
		}

		singleItems, err := parseSingle(unit, slot, step)
		if nil != err {
			return nil, err
		}
		items = append(items, singleItems...)
	}

	return items, nil
}
