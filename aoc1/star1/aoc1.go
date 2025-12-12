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

		switch direction {
		case 'L' :
			position -= amount
		case 'R' :
			position += amount
		}

		position = position % 100

		if position > 99 {
			position -= 100
		} else if position < 0 {
			position += 100
		}

		if position == 0 {
			password++
		}
	}

	fmt.Println(position, password)

}