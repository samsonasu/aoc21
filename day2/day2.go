package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	x := int64(0)
	aim := int64(0)
	z := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), " ")

		dir := parsed[0]
		val, _ := strconv.ParseInt(parsed[1], 10, 0)

		switch dir {
		case "forward":
			x += val
			z += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		}

		fmt.Printf("got [%s], new location is %d: (%d, %d)\n", scanner.Text(), aim, x, z)

	}

	fmt.Printf("multiplying for some reason yields %d", x*z)
}

func part1() {
	x := int64(0)
	z := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), " ")

		dir := parsed[0]
		val, _ := strconv.ParseInt(parsed[1], 10, 0)

		switch dir {
		case "forward":
			x += val
		case "down":
			z += val
		case "up":
			z -= val
		}

		fmt.Printf("got [%s], new location is (%d, %d)\n", scanner.Text(), x, z)

	}

	fmt.Printf("multiplying for some reason yields %d", x*z)
}
