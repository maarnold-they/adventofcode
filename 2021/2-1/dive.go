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

  posX := 0
  posY := 0

  for scanner.Scan() {
    urr := scanner.Text()
    curr := strings.Split(urr, " ")
    direction := curr[0]
    dist, err := strconv.Atoi(curr[1])
    if err != nil {
      log.Fatal(err)
    }
    switch direction {
    case "forward":
      posX = posX + dist
    case "down":
      posY = posY + dist
    case "up":
      posY = posY - dist
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(posX * posY)
}
