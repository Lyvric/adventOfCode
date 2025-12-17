package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	invalidIDs   []string
	invalidIDSum int
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

			idLength := len(strings.Split(strconv.Itoa(i), ""))
			fmt.Println("idLength:", idLength)

			idSplit := SplitByPosition(strconv.Itoa(i), idLength/2)
			fmt.Println("idSplit:", idSplit)

			if idSplit[0] == idSplit[1] {
				fmt.Println("ID repeating:", idSplit)
				invalidIDs = append(invalidIDs, strconv.Itoa(i))
			}
		}
	}

	fmt.Println("")
	fmt.Println("Repeating IDs:", invalidIDs)

	for id := range invalidIDs {
		value := invalidIDs[id]
		valueAsI, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		invalidIDSum += valueAsI
	}

	fmt.Println("")
	fmt.Println("Sum of invalid IDs:", invalidIDSum)

}

func SplitByPosition(s string, pos int) []string {
	var splitsOne []byte
	var splitsTwo []byte

	var value []string

	for i := range s {
		if i > pos-1 {
			splitsOne = append(splitsOne, s[i])
		} else {
			splitsTwo = append(splitsTwo, s[i])
		}
	}

	value = append(value, string(splitsOne))
	value = append(value, string(splitsTwo))

	return value
}
