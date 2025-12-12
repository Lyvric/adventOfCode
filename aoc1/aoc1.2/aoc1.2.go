package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	position = 50
	password = 0

	oldPos = 999
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(data), "\n")

	for line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		oldPos = position

		modAmount := amount / 100
		amount = amount % 100

		if modAmount < 0 {
			modAmount *= -1
		}
		password += modAmount

		switch direction {
		case 'L':
			position -= amount
		case 'R':
			position += amount
		}

		if position == 0 {
			password++
			continue
		}

		if position > 99 {
			position -= 100
			password++
		} else if position < 0 {
			position += 100

			if oldPos != 0 {
				password++
			}
		}
	}

	fmt.Println(position, password)
}
