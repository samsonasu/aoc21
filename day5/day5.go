package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const gridSize = 1000

func main() {
	grid := make([][]int64, gridSize)
	for i := range grid {
    	grid[i] = make([]int64, gridSize)
	}
	scanner := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile("\\d+")

	for scanner.Scan() {
		row := scanner.Text()
		numbers := re.FindAllString(row, 4)
		x1, _ := strconv.ParseInt(numbers[0], 10, 64)
		y1, _ := strconv.ParseInt(numbers[1], 10, 64)
		x2, _ := strconv.ParseInt(numbers[2], 10, 64)
		y2, _ := strconv.ParseInt(numbers[3], 10, 64)
		fmt.Println()
		fmt.Println(x1, y1, "->", x2, y2)


		if x1 == x2 { // vertical line
			fmt.Println("vert")
			if y2 > y1 {
				for y := y1; y <= y2; y++ {
					grid[x1][y] += 1
				}
			} else {
				for y := y2; y <= y1; y++ {
					grid[x1][y] += 1
				}
			}
		} else if y1 == y2 { // horizontal line
			fmt.Println("horz")
			if x2 > x1 {
				for x := x1; x <= x2; x++ {
					grid[x][y1] += 1
				}
			} else {
				for x := x2; x <= x1; x++ {
					grid[x][y1] += 1
				}
			}
		} else {
			fmt.Println("Diagonal")
		}
	}
	// printGrid(grid)
}

func printGrid(grid [][]int64) {
	overlap := 0
	debug := true
	for y, _ := range grid {
		for x, _ := range grid {
			v := grid[x][y]
			if v > 1 {
				overlap++
			}
			if (debug) {
				if (v == 0) {
					fmt.Print(".")
				} else {
					fmt.Print(v)
				}
			}
		}
		if debug { fmt.Println() }
	}
	fmt.Println("Overlap: ", overlap)
}