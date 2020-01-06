package ch4

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(m map[string]int, list []string) {
	m[k(list)]++
}

func Count(m map[string]int, list []string) int {
	return m[k(list)]
}
