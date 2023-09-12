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

type DataStrategy interface{
  Calculate([]int) int
}

type SumStrategy struct{}

func (s SumStrategy) Calculate(group []int) int {
	var sum int
	for _, val := range group {
		sum += val
	}

	return sum
}

type MaxThreeStrategy struct{}

func (s MaxThreeStrategy) Calculate(group []int) int {
 	slices.SortFunc(group, func(a, b int) int {
		return cmp.Compare(b, a)
	})
  maxThreeValues := 0

  for i := 0; i < 3 && i < len(group); i++ {
    maxThreeValues += group[i]
  }

	return maxThreeValues
}

type GroupData struct {
	groups   [][]int
	data     []int
  calculatedValue int
  strategy DataStrategy
}

func (g *GroupData) loadGroups(fileName string) {
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
}

func (g *GroupData) SetStrategy(s DataStrategy) {
	g.strategy = s
}

func (g *GroupData) CalculateGroup() {
	g.data = []int{}
	for _, group := range g.groups {
		g.data = append(g.data, g.strategy.Calculate(group))
	}
}

func (g *GroupData) CalculateData() {
  g.calculatedValue = g.strategy.Calculate(g.data)
}

func (g *GroupData) SortData() {
	slices.SortFunc(g.data, func(a, b int) int {
		return cmp.Compare(b, a)
	})
}

func main() {
	groupData := GroupData{}
	groupData.loadGroups("calories.txt")

  groupData.SetStrategy(SumStrategy{})
	groupData.CalculateGroup()
  groupData.SortData()
  fmt.Printf("Using Sum Strategy: %v\n", groupData.data)
  fmt.Printf("Using Sum Strategy. Elf with most food has these much: %v\n", groupData.data[0])

	groupData.SetStrategy(MaxThreeStrategy{})
	groupData.CalculateData()
	fmt.Printf("Using Max Strategy: %v\n", groupData.calculatedValue)
}
