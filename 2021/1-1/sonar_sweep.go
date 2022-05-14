package main

import (
  "bufio"
	"fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  file, err := os.Open("sonar_sweep.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  sum := -1
  last := 0

  for scanner.Scan() {
    curr, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal(err)
    }
    if curr > last {
      sum++
    }
    last = curr
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(sum)
}
