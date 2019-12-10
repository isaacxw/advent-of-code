package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Planet struct {
	name     string
	children []*Planet
	parent   *Planet
}

func main() {
	data := getData()
	fmt.Println("Part 1: " + strconv.Itoa(part1(data)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(data)))
}

// Part 1
func part1(data []string) int {
	planets := createMap(data)

	numberOfOrbits := 0
	for i := range planets {
		planet := planets[i]
		numberOfOrbits += countParents(planet)
	}
	return numberOfOrbits
}

// Part 2
func part2(data []string) int {
	planets := createMap(data)
	return evaluateTransfers(planets["YOU"], planets["SAN"])
}

// Read data from input file
func getData() []string {
	file, err := os.Open("input06")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return l
}

func createMap(data []string) map[string]*Planet {
	planets := make(map[string]*Planet)

	for i := range data {
		orbit := strings.Split(data[i], ")")
		nameOfOrbitee := orbit[0]
		nameOfOrbiter := orbit[1]

		from, status := planets[nameOfOrbitee]
		if !status {
			from = &Planet{name: nameOfOrbitee, children: nil, parent: nil}
			planets[nameOfOrbitee] = from
		}

		to, status := planets[nameOfOrbiter]
		if !status {
			to = &Planet{name: nameOfOrbiter, children: nil, parent: nil}
			planets[nameOfOrbiter] = to
		}

		from.children = append(from.children, to)
		to.parent = from
	}
	return planets
}

// Use a recursive call to get parent count
func countParents(planet *Planet) int {
	if planet.parent != nil {
		return countParents(planet.parent) + 1
	}
	return 0
}

// part 2
func evaluateTransfers(planetFrom *Planet, planetTo *Planet) int {
	visited := routeToRoot(planetFrom)

	currentPlanet := planetTo
	found := false
	samePlanet := Planet{}
	steps := -1

	for !found {
		if contains(visited, currentPlanet) {
			found = true
			samePlanet = *currentPlanet
			break
		}

		if currentPlanet.parent != nil {
			currentPlanet = currentPlanet.parent
		}
		steps += 1
	}

	foundFrom := false
	currentFromPlanet := planetFrom
	for !foundFrom {
		if currentFromPlanet.name == samePlanet.name {
			foundFrom = true
			break
		}
		if currentFromPlanet.parent != nil {
			currentFromPlanet = currentFromPlanet.parent
		}
		steps += 1
	}
	return steps - 1
}

// Get the path to the root node
func routeToRoot(planet *Planet) []string {
	var visited []string
	rootFound := false
	currentPlanet := planet
	for !rootFound {
		if currentPlanet.parent == nil {
			visited = append(visited, currentPlanet.name)
			rootFound = true
			break
		} else {
			visited = append(visited, currentPlanet.name)
			currentPlanet = currentPlanet.parent
		}
	}
	return visited
}

func contains(s []string, e *Planet) bool {
	for _, search := range s {
		if search == e.name {
			return true
		}
	}
	return false
}
