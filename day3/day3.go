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

	var bit_counter []int

	var input_rows []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		input_rows = append(input_rows, input) // Save these for part 2
		// First time initialize the arrays to the correct length
		if bit_counter == nil {
			l := len(input)
			bit_counter = make([]int, l)

		}
		bit_counter = count_bits(input, bit_counter)
	}
	calc_part1(bit_counter)

	// oxygen generator
	var oxygen string
	tmp_rows := input_rows

	for i := range bit_counter {
		var search byte
		if bit_counter[i] >= 0 {
			search = One
		} else {
			search = Zero
		}
		tmp_rows = filter_pos(tmp_rows, i, search)
		fmt.Println(tmp_rows)

		if len(tmp_rows) == 1 {
			oxygen = tmp_rows[0]
			break
		}

		//now reset bit counter and recalculate on the filtered list
		bit_counter = make([]int, len(tmp_rows[0]))
		for _, in := range tmp_rows {
			bit_counter = count_bits(in, bit_counter)
		}
		fmt.Println("counter: ", bit_counter)
	}

	oxygen_dec, _ := strconv.ParseInt(string(oxygen), 2, 64)
	fmt.Println("Oxygen value is: ", oxygen, oxygen_dec)
	fmt.Println("======================================")
	// CO2 Score

	// oxygen generator
	var co2 string
	tmp_rows = input_rows

	for i := range bit_counter {
		var search byte
		if bit_counter[i] >= 0 {
			search = Zero
		} else {
			search = One
		}
		tmp_rows = filter_pos(tmp_rows, i, search)
		fmt.Println(tmp_rows)

		if len(tmp_rows) == 1 {
			co2 = tmp_rows[0]
			break
		}

		//now reset bit counter and recalculate on the filtered list
		bit_counter = make([]int, len(input_rows[0]))
		for _, in := range tmp_rows {
			bit_counter = count_bits(in, bit_counter)
		}
		fmt.Println("counter: ", bit_counter)
	}

	co2_dec, _ := strconv.ParseInt(string(co2), 2, 64)
	fmt.Println("CO2 value is: ", co2, co2_dec)

	fmt.Println("Multiply O2 by CO2 and get: ", co2_dec*oxygen_dec)

}

func filter_pos(rows []string, index int, value byte) (ret []string) {
	fmt.Printf("Filtering for %v at index %d in %d rows", value, index, len(rows))
	for _, s := range rows {
		if s[index] == value {
			ret = append(ret, s)
		}
	}
	return ret
}

func calc_part1(bit_counter []int) {
	l := len(bit_counter)
	var gamma []byte = make([]byte, l)
	var epsilon []byte = make([]byte, l)

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

func count_bits(input string, bit_counter []int) []int {
	for i, ch := range input {
		switch ch {
		case Zero:
			bit_counter[i] -= 1
		case One:
			bit_counter[i] += 1
		default:
			os.Exit(1)
		}
	}
	return bit_counter
}
