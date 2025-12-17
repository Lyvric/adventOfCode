package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	voltageEntries []int
	voltageSum int
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
		// Loop through each number in one line
		// 1. Save first number in var as `t` and then check if number is higher, if so overwrite value
		// 2. Starting from entry after number for variable `ten` search for highest number usable for variable `one`
		// Combine both vars to one and add to a global list

		splitted := strings.Split(line, "")
		sLength := len(splitted)

		var ten int
		var one int

		var tenPosition int

		tenLocked := false

		for i, number := range splitted {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				ten = value
				tenPosition = i
				continue
			}

			if !tenLocked && ten < value {
				if sLength-1 != i {
					ten = value
					tenPosition = i
				}
			}

			if sLength-1 == i {
				tenLocked = true
			}
		}

		for i, number := range splitted[tenPosition:] {
			if i == 0 {
				continue
			}

			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			if i == 1 {
				one = value
				continue
			}

			if one < value {
				one = value
			}
		}

		fmt.Println("Voltage:", ten, one)

		combined, err := strconv.Atoi(strconv.Itoa(ten) + strconv.Itoa(one))
		if err != nil {
			panic(err)
		}

		voltageEntries = append(voltageEntries, combined)
	}

	for _, entry := range voltageEntries {
		voltageSum += entry
	}

	fmt.Println("")
	fmt.Println("Voltage Entries:", voltageEntries)

	fmt.Println("Voltage Sum:", voltageSum)
}
