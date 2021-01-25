package parser

import "strings"

// Ordered list of cron parts
var Slots = []Slot{
	{
		Label:           "minute",
		Min:             0,
		Max:             59,
		ValidCharacters: `^[\d|\*|\-|,|/]+$`,
	},
	{
		Label:           "hour",
		Min:             0,
		Max:             23,
		ValidCharacters: `^[\d|\*|\-|,|/]+$`,
	},
	{
		Label:           "day of month",
		Min:             1,
		Max:             31,
		ValidCharacters: `^[\d|\*|\-|,|/|?]+$`,
	},
	{
		Label:           "month",
		Min:             1,
		Max:             12,
		ValidCharacters: `^[\d|\*|\-|,|/]+$`,
		Replacer: strings.NewReplacer(
			"jan", "1",
			"feb", "2",
			"mar", "3",
			"apr", "4",
			"may", "5",
			"jun", "6",
			"jul", "7",
			"aug", "8",
			"sep", "9",
			"oct", "10",
			"nov", "11",
			"dec", "12",
		),
	},
	{
		Label:           "day of week",
		Min:             0,
		Max:             6,
		ValidCharacters: `^[\d|\*|\-|,|/|?]+$`,
		Replacer: strings.NewReplacer(
			"sun", "0",
			"mon", "1",
			"tue", "2",
			"wed", "3",
			"thr", "4",
			"fri", "5",
			"sat", "6",
		),
	},
}

