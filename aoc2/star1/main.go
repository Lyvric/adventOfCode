package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	invalidIDs []string
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(data), ",")

	for line := range lines {
		rangeOne, err := strconv.Atoi(strings.Split(line, "-")[0])
		if err != nil {
			panic(err)
		}

		rangeTwo, err := strconv.Atoi(strings.Split(line, "-")[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(rangeOne, rangeTwo)

		for i := rangeOne; i <= rangeTwo; i++ {
			fmt.Println("ID:", i)

			splitedID := strings.Split(strconv.Itoa(i), "")
			cSplitedID := splitedID

			for id := range splitedID {
				for id2 := range cSplitedID {
					if id != id2 {
						break
					}
				}
			}
		}
	}
}
