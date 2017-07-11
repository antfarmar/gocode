package main

import "fmt"

func main() {
	s, t := "ALBONCGEDST", "LOABNGCESDT"
	m, n := len(s), len(t)
	lcs := lcs(s, t)
	printlcs(s, lcs, m, n)
	fmt.Println()
}

// Longest Common Subsequence
func lcs(s1, s2 string) [][]int {
	m, n := len(s1), len(s2)
	var memo [][]int = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case s1[i-1] == s2[j-1]:
				memo[i][j] = memo[i-1][j-1] + 1 // + cost
			case memo[i][j-1] >= memo[i-1][j]:
				memo[i][j] = memo[i][j-1] + 0 // + cost
			default:
				memo[i][j] = memo[i-1][j] + 0 // + cost
			}
		}
	}
	return memo
}

func printlcs(s string, r [][]int, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	switch {
	case r[i][j] > r[i-1][j] && r[i][j] > r[i][j-1]:
		printlcs(s, r, i-1, j-1)
		fmt.Printf("%c", s[i-1])
	case r[i][j-1] >= r[i-1][j]:
		printlcs(s, r, i, j-1)
	default:
		printlcs(s, r, i-1, j)
	}
}
