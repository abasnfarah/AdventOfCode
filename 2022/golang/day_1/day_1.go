package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type groupData struct {
	groups   [][]int
	data     []int
	maxValue int
	topThree int
}

func (g *groupData) generateGroups(fileName string) {
	d, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(d), "\n")

	var group []int
	for _, line := range lines {
		if line == "" {
			if len(group) > 0 {
				g.groups = append(g.groups, group)
				group = []int{}
			}
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		group = append(group, val)
	}

	if len(group) > 0 {
		g.groups = append(g.groups, group)
	}

	g.setData()
	g.setMaxValue()
	g.setTopThree()
}

func (g *groupData) setData() {
	for _, group := range g.groups {
		var sum int
		for _, val := range group {
			sum += val
		}
		g.data = append(g.data, sum)
		slices.SortFunc(g.data, func(a, b int) int {
			return cmp.Compare(b, a)
		})
	}
}

func (g *groupData) setMaxValue() {
	g.maxValue = slices.Max(g.data)
}

func (g *groupData) setTopThree() {
	for i := 0; i < 3; i++ {
		g.topThree += g.data[i]
	}
}

func main() {
	groupData := groupData{}
	groupData.generateGroups("calories.txt")

	fmt.Printf("Groups: %v\n\r", groupData.groups)
	fmt.Printf("Data: %d\n\r", groupData.data)
	fmt.Printf("Max: %d\n\r", groupData.maxValue)
	fmt.Printf("Data: %d\n\r", groupData.topThree)
}
