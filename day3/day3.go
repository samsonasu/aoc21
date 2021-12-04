package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Zero = 48
const One = 49

func main() {

	bit_counter := []int{0, 0, 0, 0, 0}
	gamma := []byte("00000")
	epsilon := []byte("00000")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		input := scanner.Text()
		fmt.Println("Reading:", input)

		for i, ch := range input {
			fmt.Print(ch, i)
			switch ch {
			case Zero:
				fmt.Print("Got a Zero")
				bit_counter[i] -= 1
			case One:
				fmt.Print("Got a One")
				bit_counter[i] += 1
			default:
				os.Exit(1)
			}
			fmt.Println("bit_counter is", bit_counter)
		}

		for i, _ := range gamma {
			if bit_counter[i] > 0 {
				// More ones than zeros
				gamma[i] = One
				epsilon[i] = Zero
			} else {
				gamma[i] = Zero
				epsilon[i] = One
			}
		}
		gamma_dec, _ := strconv.ParseInt(string(gamma), 2, 64)
		fmt.Println("Gamma: ", string(gamma), gamma_dec)

		epsilon_dec, _ := strconv.ParseInt(string(epsilon), 2, 64)
		fmt.Println("Epsilon: ", string(epsilon), epsilon_dec)

		fmt.Println("Multiply for Power: ", gamma_dec*epsilon_dec)
	}
}
