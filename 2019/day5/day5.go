package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type OpCodes struct {
	code int
	modeA int
	modeB int
	modeC int
}

func main() {
	line, err := ioutil.ReadFile("input5")
	if err != nil {
		panic(err)
	}

	splitLine := strings.Split(string(line), ",")

	// slice ints
	var nums []int
	for _, s := range splitLine {
		i, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			fmt.Println("Failed to split the array of numbers.")
			os.Exit(1)
		}
		nums = append(nums, i)
	}

	calculate(nums)

}

func defineOpCode(val int) OpCodes {
	s := fmt.Sprintf("%05d", val) // ensure input is 5 digits long with leading zeros if necessary

	// slice strings up
	opcode, _ := strconv.Atoi(s[3:])
	modeA, _ := strconv.Atoi(s[2:3])
	modeB, _ := strconv.Atoi(s[1:2])
	modeC, _ := strconv.Atoi(s[0:1])
	
	// set opcode values
	return OpCodes {
		code: opcode,
		modeA: modeA,
		modeB: modeB,
		modeC: modeC,
	}
}

func getValue(arr []int, position, mode int) int {
	if mode == 1 {
		return position
	}

	return arr[position]
}

func calculate(list []int) int {
	pointer := 0
	for pointer < len(list) {
		opcode := defineOpCode(list[pointer])

		switch opcode.code {
		case 1:
			// opcode 1: add
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			posC := list[pointer + 3]
			list[posC] = getValue(list, posA, opcode.modeA) + getValue(list, posB, opcode.modeB)
			pointer += 4 // increment by 4
		case 2:
			// opcode 2: multiply
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			posC := list[pointer + 3]
			list[posC] = getValue(list, posA, opcode.modeA) * getValue(list, posB, opcode.modeB)
			pointer += 4
		case 3:
			// read user input and save to given parameter
			posA := list[pointer + 1]
			var input int
			fmt.Scanln(&input)
			list[posA] = input
			pointer += 2 // increment by 2

		case 4:
			// returns the value at position
			posA := list[pointer + 1]
			value := getValue(list, posA, opcode.modeA)
			fmt.Println(value)
			pointer += 2
		case 5:
			// jump-if-true
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			valA := getValue(list, posA, opcode.modeA)
			valB := getValue(list, posB, opcode.modeB)

			if valA != 0 {
				pointer = valB
			} else {
				pointer += 3
			}
		case 6:
			// jump-if-false
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			valA := getValue(list, posA, opcode.modeA)
			valB := getValue(list, posB, opcode.modeB)

			if valA == 0 {
				pointer = valB
			} else {
				pointer += 3
			}
		case 7:
			// less than
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			posC := list[pointer + 3]
			valA := getValue(list, posA, opcode.modeA)
			valB := getValue(list, posB, opcode.modeB)

			if valA < valB {
				list[posC] = 1
			} else {
				list[posC] = 0
			}

			pointer += 4
		case 8:
			// equals
			posA := list[pointer + 1]
			posB := list[pointer + 2]
			posC := list[pointer + 3]
			valA := getValue(list, posA, opcode.modeA)
			valB := getValue(list, posB, opcode.modeB)

			if valA == valB {
				list[posC] = 1
			} else {
				list[posC] = 0
			}

			pointer += 4
		case 99:
			return pointer
		}
	}
	return pointer
}

