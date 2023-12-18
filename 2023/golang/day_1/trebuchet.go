package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func genSubstrings(s string) []string {
  var subs []string
  for i := 0; i < len(s); i++ {
    for j := i + 1; j <= len(s); j++ {
      if len(s[i:j]) >= 1 {
        subs = append(subs, s[i:j])
      }
    }
  }
  return subs
}

func main() {

  stringToNumsTable := map[string]int{
    "one":  1, "two": 2, "three":3, 
    "four": 4, "five": 5, "six":  6, 
    "seven":7, "eight":8, "nine": 9,}

  d, err := os.ReadFile("trebuchet.txt")
  if err != nil {
    log.Fatal(err)
  }

  lines := strings.Split(string(d), "\n")

  sum := 0

  for _, line := range lines {
    fmt.Println(line)    
    firstNum := -1
    secondNum := -1
    if line == "" {
      continue
    }

    subs := genSubstrings(line)
    
    for _, sub := range subs {
      val, ok := stringToNumsTable[sub]
      if ok {
        fmt.Println("Found sub: ", sub)
        if firstNum == -1 {
          firstNum = val
        } else {
          secondNum = val
        }
      } else if len(sub) == 1 {
        newInt, err := strconv.Atoi(sub)
        if err != nil {
          continue
        }
        if firstNum == -1 {
          firstNum = newInt
        } else {
          secondNum = newInt
        }
      }
      fmt.Printf("firstNum: %d, SecondkNum: %d \n", firstNum, secondNum)
    }

    if secondNum == -1 {
      secondNum = firstNum
    }

    fmt.Printf("firstNum: %d, SecondkNum: %d \n", firstNum, secondNum)

    sum += firstNum * 10 + secondNum
    fmt.Println("sum: ", sum)

  }

  fmt.Println(sum)
}

