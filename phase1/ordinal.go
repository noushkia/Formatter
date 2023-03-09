package phase1

// Courtesy of https://github.com/sbani/go-humanizer/blob/master/numbers/ordinal.go

import (
	"math"
	"strconv"
)

// Ordinal returns the ordinal string for a specific integer.
func Ordinal(number int) string {
	absNumber := int(math.Abs(float64(number)))

	i := absNumber % 100
	if i == 11 || i == 12 || i == 13 {
		return "th"
	}

	switch absNumber % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// Ordinalize the number by adding the Ordinal to the number.
func Ordinalize(number []byte) []byte {
	numberInt, _ := strconv.Atoi(string(number[:]))
	return append(number, []byte(Ordinal(numberInt))...)
}
