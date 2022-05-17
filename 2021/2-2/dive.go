package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, err := os.Open("dive.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  hPos := 0
  depth := 0
  aim := 0

  for scanner.Scan() {
    inputTxt := scanner.Text()
    inputArr := strings.Split(inputTxt, " ")
    direction := inputArr[0]
    dist, err := strconv.Atoi(inputArr[1])
    if err != nil {
      log.Fatal(err)
    }
    switch direction {
    case "forward":
      hPos = hPos + dist
      depth = depth + (aim * dist)
    case "down":
      aim = aim + dist
    case "up":
      aim = aim - dist
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(hPos * depth)
}
