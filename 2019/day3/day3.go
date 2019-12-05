package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

type vect struct {
	direction rune
	distance int
}

type coordinates struct {
	x, y int
}

func readVectorPaths(file string) [][]vect {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.TrimSpace(string(text))
	var vectorPaths [][]vec

	for _, line := range strings.Split(lines, "\n")

}