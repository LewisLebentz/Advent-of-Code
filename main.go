package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
  segment int
  total  int
}

type Elves struct {
  elf []Elf
}

var elves []Elf

func main() {

  content, err := ioutil.ReadFile("input.txt")

   if err != nil {
        log.Fatal(err)
   }

  input := string(content)
  splitInput := strings.Split(input, "\n\n")
  fmt.Println("Segment length: ", len(strings.Split(splitInput[0], "\n")))
  for i := 0; i < len(splitInput); i++ {
    segment := strings.Split(splitInput[i], "\n")
    var total int
    for j := 0; j < len(segment); j++ {
      fmt.Println(segment[j])
      intVar, err := strconv.Atoi(segment[j])
       if err != nil {
            log.Fatal(err)
       }
      
      
      total = total + intVar
    }
    fmt.Println("Total for segment: ", total)
    elf := Elf{
      segment: i,
      total: total,
    }
    fmt.Println(elf)
    elves = append(elves, elf)

  }
  // fmt.Println(elves.segment)
  sort.Slice(elves, func(i, j int) bool {
      return elves[i].total < elves[j].total
  })
  for _, v := range elves {
      fmt.Println(v)
  }  
}
