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

			// Build splits with loops
			// Loop through id until number repeats
			// Then continue on to next split part
			// Compare the parts like this: part1 ?= part2 -> if true: part2 ?= part3; ...
			// If all parts are the same add id to invalidIDs, if one part is different ignore id

			splittedID := strings.Split(strconv.Itoa(i), "")
			idParts := SplitToParts(splittedID)

			fmt.Println("Parts:", idParts)

			hits := 0

			for i := range idParts {
				value := idParts[i]

				if len(idParts)-1 == i {
					break
				}

				nextValue := idParts[i+1]

				if value != nextValue {
					break
				}

				hits++
			}

			fmt.Println("Hits:", hits)
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

func SplitToParts(splittedID []string) []string {
	var idParts []string
	var partBuilder []string

	indexed := 0

	for i := range splittedID {
		iValue := splittedID[i]

		for indexed > i {
			fmt.Println("SKIP", indexed, i)
			continue
		}

		partBuilder = append(partBuilder, iValue)

		for j := range splittedID {
			jValue := splittedID[j]

			if j == 0 {
				continue
			}

			fmt.Println("values:", iValue, jValue)

			if iValue == jValue {
				fmt.Println("partBuilder:", partBuilder)
				idParts = append(idParts, strings.Join(partBuilder, ""))
				fmt.Println("idParts:", idParts)
				partBuilder = nil
				break
			}

			partBuilder = append(partBuilder, jValue)
			indexed++
		}
	}

	return idParts
}
