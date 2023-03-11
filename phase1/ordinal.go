package phase1

// Courtesy of https://github.com/sbani/go-humanizer/blob/master/numbers/ordinal.go

import (
	"math"
	"strconv"
	"unicode"
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
	var numeric []byte
	for i := 0; i < len(number); i++ {
		if unicode.IsDigit(rune(number[i])) {
			numeric = append(numeric, number[i])
		} else {
			break
		}
	}
	numberInt, _ := strconv.Atoi(string(numeric))
	ordinal := []byte(Ordinal(numberInt))
	return append(numeric, append([]byte(""), append(ordinal, number[len(numeric):]...)...)...)
}
