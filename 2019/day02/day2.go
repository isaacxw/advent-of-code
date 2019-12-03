package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func getData() (output []int) {
	data, err := ioutil.ReadFile("input02")
	if err != nil {
		panic(err)
	}

	splitStr := string(data)
	for _, s := range strings.Split(strings.ReplaceAll(splitStr, "\n", ""), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		output = append(output, i)
	}
	return
}

func calculate(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		firstIndex, secondIndex, store := input[i + 1], input[i + 2], input[i + 3]
		switch opcode {
		case 1:
			input[store] = input[firstIndex] + input[secondIndex]
		case 2:
			input[store] = input[firstIndex] * input[secondIndex]
		case 99:
			return input
		}
	}
	return input
}

func part1() {
	data := getData()
	data[1] = 12
	data[2] = 2
	data = calculate(data)
	fmt.Printf("%v\n", data[0])
}

func part2() {
	var answer int
	target := 19690720

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			data := getData()
			data[1] = i
			data[2] = j
			data = calculate(data)
			if data[0] == target {
				answer = (100 * i) + j
			}
		}
	}	
	fmt.Printf("%v\n", answer)
}