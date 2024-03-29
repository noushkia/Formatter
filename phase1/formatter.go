package phase1

import (
	"regexp"
	"strings"
)

func splitIntoSentences(sentences string) []string {
	var result []string
	last := 0
	for i, r := range sentences {
		if r == '.' || r == '!' || r == '?' {
			result = append(result, sentences[last:i+1])
			last = i + 1
		}
	}
	return result
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
	isNumericReg := regexp.MustCompile(`[0-9]+`)
	for i := range sentences {
		words := strings.Fields(sentences[i])
		for j := range words {
			if isNumericReg.MatchString(words[j]) {
				words[j] = string(isNumericReg.ReplaceAllFunc([]byte(words[j]), Ordinalize))
			}
		}
		sentences[i] = strings.Join(words, " ")
	}
	return sentences
}

func Format(sentences string) string {
	seperatedSentences := splitIntoSentences(sentences)
	seperatedSentences = capitalizeSentences(seperatedSentences)
	seperatedSentences = toOrdinal(seperatedSentences)

	return strings.Join(seperatedSentences, " ")
}
