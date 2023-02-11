package main

import (
	"math"
)

// how many digit's groups to process
const groupsNumber int = 4

var _smallNumbers = []string{
	"kosong", "satu", "dua", "tiga", "empat",
	"lima", "enam", "tujuh", "lapan", "sembilan",
	"sepuluh", "sebelas",
}
var _tens = []string{
	"", "sepuluh",
}
var _scaleNumbers = []string{
	"", "ribu", "juta", "bilion",
}

type digitGroup int

// Convert converts number into the words representation.
func Convert(number int) string {
	return convert(number, false)
}

// ConvertAnd converts number into the words representation
// with " and " added between number groups.
func ConvertAnd(number int) string {
	return convert(number, true)
}

func convert(number int, useAnd bool) string {
	// Zero rule
	if number == 0 {
		return _smallNumbers[0]
	}

	// Divide into three-digits group
	var groups [groupsNumber]digitGroup
	positive := math.Abs(float64(number))

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = digitGroup(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i], useAnd)
	}
	combined := textGroup[0]
	and := useAnd && (groups[0] > 0 && groups[0] < 100)

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + _scaleNumbers[i]

			if len(combined) != 0 {
				prefix += separator(and)
			}

			and = false

			combined = prefix + combined
		}
	}

	if number < 0 {
		combined = "negatif " + combined
	}

	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group digitGroup, useAnd bool) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		ret += _smallNumbers[hundreds] + " ratus"

		if tensUnits != 0 {
			ret += separator(useAnd)
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		if tens == 1 {
			ret += _tens[tens]
		} else {
			ret += _smallNumbers[tens] + " puluh"
		}

		if units != 0 {
			ret += " " + _smallNumbers[units]
		}
	} else if tensUnits != 0 {
		if tensUnits < 12 {
			ret += _smallNumbers[tensUnits]
		} else {
			ret += _smallNumbers[units] + " belas"
		}
	}

	return
}

// separator returns proper separator string between
// number groups.
func separator(useAnd bool) string {
	if useAnd {
		return " dan "
	}
	return " "
}
