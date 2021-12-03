package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	find_increases(3)
}

func find_increases(window_size int) {
	// part 1
	scanner := bufio.NewScanner(os.Stdin)

	increased_count := 0
	var history []int64
	for scanner.Scan() {
		parsed, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		fmt.Printf("Got %d, history is %v\n", parsed, history)
		if len(history) < window_size {
			history = append(history, parsed)
			continue
		}
		prev := sum(history)

		history = append(history, parsed)
		if len(history) >= window_size {
			history = history[1:]
		}
		curr := sum(history)
		if curr > prev {
			increased_count = increased_count + 1
		}
	}
	fmt.Printf("SONAR increased %d times", increased_count)
}

func sum(arr []int64) int64 {
	var acc int64 = 0
	for _, v := range arr {
		acc += v
	}
	return acc
}
