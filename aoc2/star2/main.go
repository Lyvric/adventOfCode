package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type RangePair struct {
	Start int
	End   int
}

var (
	invalidIDs   []string
	invalidIDSum int

	ranges []RangePair

	threads = 30
	mutex   = &sync.Mutex{}
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

		ranges = append(ranges, RangePair{Start: rangeOne, End: rangeTwo})
	}

	DeployWorkers()

	fmt.Println("")
	fmt.Println("Repeating IDs:", invalidIDs)

	for _, value := range invalidIDs {
		valueAsI, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		invalidIDSum += valueAsI
	}

	fmt.Println("")
	fmt.Println("Sum of invalid IDs:", invalidIDSum)
}

func DeployWorkers() {
	rangesChan := make(chan RangePair, len(ranges))
	for _, r := range ranges {
		rangesChan <- r
	}
	close(rangesChan)

	var wg sync.WaitGroup
	wg.Add(threads)

	for i := range threads {
		go func(workerID int) {
			defer wg.Done()

			var runtimeInvalidIDs []string
			for rangePair := range rangesChan {
				returnedList := Runtime(rangePair, runtimeInvalidIDs)
				runtimeInvalidIDs = returnedList
			}
			mutex.Lock()
			invalidIDs = append(invalidIDs, runtimeInvalidIDs...)
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
}

func Runtime(rangePair RangePair, runtimeInvalidIDs []string) []string {
	for i := rangePair.Start; i <= rangePair.End; i++ {
		idLength := len(strconv.Itoa(i))
		idIsInvalid := TryCatchCombinations(strconv.Itoa(i), idLength)

		if idIsInvalid {
			runtimeInvalidIDs = append(runtimeInvalidIDs, strconv.Itoa(i))
		}
	}

	return runtimeInvalidIDs
}

func TryCatchCombinations(s string, idLength int) bool {
	for patternLength := 1; patternLength <= idLength/2; patternLength++ {
		if idLength%patternLength != 0 {
			continue
		}

		pattern := s[0:patternLength]
		testString := strings.Repeat(pattern, idLength/patternLength)
		if testString == s {
			return true
		}
	}
	return false
}
