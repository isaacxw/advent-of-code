package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input01")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var modules []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		modules = append(modules, weight)

	}

	var totalFuel int

	for i := 0; i < len(modules); i++ {
		fuel := modules[i] / 3

		if (fuel - 2) > 0 {
			fuel = fuel - 2
			totalFuel += fuel
			modules = append(modules, fuel)
		}
	}

	fmt.Println("Total fuel: " + strconv.Itoa(totalFuel))
}
