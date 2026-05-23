package cowsay

import (
	"fmt"
	"slices"
	"strings"

)

func CowSay(lines []string) {
	// Find the longest length in our slice of strings
	longestLine := slices.MaxFunc(lines, func(a string, b string) int {
		return len(a) - len(b)
	})
	boxWidth := len(longestLine) + 4
	repeated := strings.Repeat("-", boxWidth - 2)

	fmt.Printf(" %s ", repeated)
	for _, line := range lines {
		padLen := boxWidth - 4 - len(line)
		if padLen < 0 {
			padLen = 0
		}

		padding := strings.Repeat(" ", padLen)
		str := "| " + line + padding + " |\n"
		fmt.Printf("%s", str)
	}
	fmt.Printf(" %s\n", repeated)
}
