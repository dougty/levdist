package levdist

import (
	"strings"
)

func multiMin(v ...int) int {
	num := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < num {
			num = v[i]
		}
	}
	return num
}

// Measure Levenshtein distance between two strings
func Measure(s string, t string) int {
	s = strings.TrimSpace(s)
	t = strings.TrimSpace(t)

	m := len(s)
	n := len(t)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	s = strings.ToLower(s)
	t = strings.ToLower(t)

	if s == t {
		return 0
	}

	v0 := make([]int, n+1)
	v1 := make([]int, n+1)

	for i := 0; i <= n; i++ {
		v0[i] = i
	}

	for i := 0; i <= m-1; i++ {
		v1[0] = i + 1

		for j := 0; j <= n-1; j++ {
			delCost := v0[j+1] + 1
			insCost := v1[j] + 1
			subCost := v0[j]
			if s[i] != t[j] {
				subCost = v0[j] + 1
			}

			v1[j+1] = multiMin(delCost, insCost, subCost)
		}

		v1, v0 = v0, v1
	}

	return v0[n]
}
