package phase1

import (
	"regexp"
	"strings"
)

func splitIntoSentences(s string) []string {
	var sentences []string
	last := 0
	for i, r := range s {
		if r == '.' || r == '!' || r == '?' {
			sentences = append(sentences, s[last:i+1])
			last = i + 1
		}
	}
	return sentences
}

func capitalizeSentences(sentences []string) []string {
	for i := range sentences {
		j := 1
		for ; sentences[i][j-1] == ' '; j++ {
		}
		sentences[i] = sentences[i][:j-1] + strings.ToUpper(sentences[i][j-1:j]) + sentences[i][j:]
	}
	return sentences
}

func toOrdinal(sentences []string) []string {
	// todo: The regex could be improved to check if the ordinal value is correct
	isNumericReg := regexp.MustCompile("[0-9]+")
	for i := range sentences {
		sentences[i] = string(isNumericReg.ReplaceAllFunc([]byte(sentences[i]), Ordinalize))
	}
	return sentences
}

func Format(sentences string) string {
	seperatedSentences := splitIntoSentences(sentences)
	seperatedSentences = capitalizeSentences(seperatedSentences)
	seperatedSentences = toOrdinal(seperatedSentences)

	return strings.Join(seperatedSentences, "")
}