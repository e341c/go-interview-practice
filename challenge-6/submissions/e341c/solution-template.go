package challenge6
import (
	"strings"
	"regexp"
)
func CountWordFrequency(text string) map[string]int {
	res := make(map[string]int)
	re := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
	split := re.FindAllString(text, -1)
	for _, w := range split {
		lw := strings.ToLower(w)
		cw := strings.ReplaceAll(lw, "'", "")
		res[cw]++
	}
	return res
} 