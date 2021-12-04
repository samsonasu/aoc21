package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	size    int
	squares [][]Square
}

type Square struct {
	value  int64
	marked bool
}

const boardSize = 5

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var boards []Board
	// First row is the picks
	scanner.Scan()
	picks := strings.Split(scanner.Text(), ",")
	fmt.Println("Picks: ", picks)

	var board Board
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			//New Board
			board = newBoard(boardSize)
			boards = append(boards, board)
			fmt.Printf("%v", boards)
		} else {
			appendRow(&board, row)
		}

	}
}

func newBoard(size int) Board {
	b := Board{size: size}
	b.squares = make([][]Square, size)
	for i := range b.squares {
		b.squares[i] = make([]Square, size)
	}

	return b
}

func appendRow(b *Board, row string) {
	squares := make([]Square, b.size)
	re := regexp.MustCompile("\\s+")
	values := re.Split(row, b.size)
	fmt.Printf("Appending %v with size %d\n", values, b.size)
	for i, v := range values {
		parsed, _ := strconv.ParseInt(v, 10, 64)
		squares[i] = Square{marked: false, value: parsed}
	}
	b.squares = append(b.squares, squares)
}

func (b Board) String() string {
	return fmt.Sprintf("'%d x %d Board'", b.size, b.size)
}
