package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	voltageEntries []int
	voltageSum     int

	batterySlotAmount = 12
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

		var batteries []int
		lastBatteryPos := 0
		remainingBatterySlots := batterySlotAmount
		currentBatterySlot := 0

		absolutePosition := 0

		fmt.Println("Checking these batteries:", splitted)

		for range batterySlotAmount {
			fmt.Println("Checking for next battery")
			for j, number := range splitted[absolutePosition:] {

				fmt.Println("splitted:", splitted[absolutePosition:])

				if j == 0 && lastBatteryPos != 0 {
					continue
				}

				value, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}

				if j == 1 && lastBatteryPos != 0 || j == 0 {
					batteries = append(batteries, value)
					lastBatteryPos = j

					fmt.Println("Added battery to list:", value, j)

					continue
				}

				if batteries[currentBatterySlot] < value {
					if sLength-remainingBatterySlots-1 != j {
						batteries[currentBatterySlot] = value
						lastBatteryPos = j

						fmt.Println("Added battery to list:", value, j)
					}
				}


				// Problem is, that sLength is the length of overall battery list
				// We need to somehow calc how much we need to reduct from sLength for the real length, or cal it new
				fmt.Println("END", sLength-1 == j, sLength-1, j)

				if sLength-1 == j {
					remainingBatterySlots--
					currentBatterySlot++

					absolutePosition += lastBatteryPos

					fmt.Println("Absolute Pos:", absolutePosition)

					break
				}
			}
		}

		fmt.Println("Voltage:", batteries)

		// fmt.Println("Voltage:", ten, one)

		// combined, err := strconv.Atoi(strconv.Itoa(ten) + strconv.Itoa(one))
		// if err != nil {
		// 	panic(err)
		// }

		// voltageEntries = append(voltageEntries, combined)
	}

	for _, entry := range voltageEntries {
		voltageSum += entry
	}

	fmt.Println("")
	fmt.Println("Voltage Entries:", voltageEntries)

	fmt.Println("Voltage Sum:", voltageSum)
}
