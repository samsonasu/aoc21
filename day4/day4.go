package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Board struct {
	size    int
	squares [][]Square
}

type Square struct {
	value  int64
	marked bool
}

func (s *Square) Mark() {
	s.marked = true
}

const boardSize = 5

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var allBoards []Board
	// First row is the picks
	scanner.Scan()
	picks := strings.Split(scanner.Text(), ",")
	fmt.Println("Picks: ", picks)
	scanner.Scan() // read a blank line

	tmpBoard := newBoard(boardSize)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			// Done reading the board, so add the last one to the list of all boards
			allBoards = append(allBoards, tmpBoard)
			tmpBoard = newBoard(boardSize)
		} else {
			appendRow(&tmpBoard, row)
		}
	}
	allBoards = append(allBoards, tmpBoard)
	fmt.Printf("Finished scanning, boards = %v", allBoards)

	fmt.Println("Picks: ", picks)
	for _, pick := range picks { 
		p, _ := strconv.ParseInt(pick, 10, 64)
		fmt.Println("Picked ", p)
		for b := range allBoards {
			for row := 0; row < boardSize; row++ {
				for col := 0; col < boardSize; col++ {
					square := allBoards[b].squares[row][col]
					if square.value == p {
						allBoards[b].squares[row][col].Mark()
					}
				}
			}
		}
		fmt.Printf("%v", allBoards)
		time.Sleep(1 * time.Second)
	}
}

func newBoard(size int) Board {
	b := Board{size: size}

	return b
}

func appendRow(b *Board, row string) {
	squares := make([]Square, b.size)
	re := regexp.MustCompile("(\\d+)")
	values := re.FindAllString(row, b.size)
	fmt.Printf("Appending %v with size %d\n", values, b.size)
	for i, v := range values {
		parsed, _ := strconv.ParseInt(v, 10, 64)
		squares[i] = Square{marked: false, value: parsed}
	}
	b.squares = append(b.squares, squares)
	fmt.Println(b)
}

func (b Board) String() string {
	str := ""
	str += fmt.Sprintf("\n'%d x %d Board:'\n", b.size, b.size)
	for _, row := range b.squares {
		for i, square := range row {
			var space string
			if i == 0 {
				space = ""
			} else {
				space = " "
			}
			if square.marked {
				str += fmt.Sprintf("\033[31;1;4m%s%2d\033[0m", space, square.value)
			} else {
				str += fmt.Sprintf("%s%2d", space, square.value)
			}
		}
		str += "\n"
	}
	return str
}
