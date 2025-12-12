package main

import (
	"os"
	"strings"
)


func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
	}
}