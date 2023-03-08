package flag

import (
	"math/rand"
	"strings"
)

type replacementTable struct {
	value string
	corresponding string
}
type Flag struct {
	prefix string
	initialValue string
	value string
	random float32
	caseFormat string
}

var (
	invalidChars = []string{
		"!",
		"\"",
		"#",
		"$",
		"%",
		"&",
		"'",
		"(",
		")",
		"*",
		"+",
		",",
		"-",
		".",
		"/",
		":",
		";",
		"<",
		"=",
		">",
		"?",
		"@",
		"[",
		"\\",
		"]",
		"^",
		"_",
		"`",
		"{",
		"|",
		"}",
		"~",
	}
	replaceTable = []replacementTable{
		{
			value: "A",
			corresponding: "4",
		},
		{
			value: "E",
			corresponding: "3",
		},
		{
			value: "I",
			corresponding: "1",
		},
		{
			value: "O",
			corresponding: "0",
		},
		{
			value: "S",
			corresponding: "5",
		},
	}
)

func NewFlagBuilder() *Flag {
	return &Flag{
		random: 0.5,
	}
}

func (f *Flag) SetPrefix(prefix string) {
	f.prefix = prefix
}

func (f *Flag) SetCase(caseOption string) {
	f.caseFormat = caseOption
}

func (f *Flag) SetInitialValue(initial string) {
	f.initialValue = initial
}

func (f *Flag) SetRandom(random float32) {
	f.random = random
}

func (f *Flag) formatCase() {
	switch f.caseFormat {
	case "preserve":
		return
	case "upper":
		f.value = strings.ToUpper(f.value)
	case "lower":
		f.value = strings.ToLower(f.value)
	}
}

func (f *Flag) Build() string {
	f.cleanString()
	f.formatCase()
	f.randomReplaceChars()
	return f.prefix + "{" + f.value + "}"
}

func (f *Flag) cleanString() {
	purgedFlag := f.initialValue
	for _, char := range invalidChars {
		purgedFlag = strings.ReplaceAll(purgedFlag, char, "")
	}
	purgedFlag = strings.TrimSpace(purgedFlag)
	purgedFlag = strings.ReplaceAll(purgedFlag, " ", "_")
	f.value = purgedFlag
}


func (f *Flag) randomReplaceChars() {
	var final string
  for _, char := range f.value {
		// if the char can be replaced
		canBeReplaced := false
		for _, table := range replaceTable {
			if strings.ToUpper(string(char)) == table.value {
				canBeReplaced = true
			}
		}
		// if the char can be replaced and the random number is less than the random value
		if canBeReplaced && rand.Float32() < f.random {
			for _, table := range replaceTable {
				if strings.ToUpper(string(char)) == table.value {
					final += table.corresponding
				}
			}
		} else {
			final += string(char)
		}
	}
	f.value = final
}
