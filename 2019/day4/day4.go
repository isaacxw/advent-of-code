package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", part1())
	fmt.Printf("%v\n", part2())
}

func part1() (sum int) {
	lower, upper := 307237, 769058 // inputs
	for lower <= upper {
		str := strconv.Itoa(lower)
		if increasing(str) && containsDouble(str) {
			sum++
		}
		lower++
	}
	return
}

func part2() (sum int) {
	lower, upper := 307237, 769058
	for lower <= upper {
		str := strconv.Itoa(lower)
		if increasing(str) && exactlyTwo(str) {
			sum++
		}
		lower++
	}
	return
}


func increasing(ss string) bool {
	for i := 0; i < len(ss)-1; i++ {
		a, _ := strconv.Atoi(string(ss[i]))
		b, _ := strconv.Atoi(string(ss[i+1]))
		if a > b {
			return false
		}
	}
	return true
}

func containsDouble(ss string) bool {
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			return true
		}
	}
	return false
}

// Repeated digits are exactly 2 digits long
func exactlyTwo(ss string) bool {
	ss = "a" + ss + "a"
	for i := 1; i < len(ss)-2; i++ {
		if ss[i-1] != ss[i] && ss[i] == ss[i+1] && ss[i+2] != ss[i] {
			return true
		}
	}
	return false
}
