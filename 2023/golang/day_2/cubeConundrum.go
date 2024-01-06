package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	linesbuff, err := os.ReadFile("gameRecords.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(linesbuff), "\n")
	fmt.Println(lines)

	// for line := range lines {
	// fmt.Println(line)
	// }

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	// Sum of the id's of those games
	// Store in map of strings to integers
	// all at onnce red = 12, green = 13, blue = 14
	// hashmap := map[string]int{"red": 12, "green": 13, "blue": 14}
	// var hashmapCopy map[string]int
	// copy(hashmapCopy, hashmap)
}
