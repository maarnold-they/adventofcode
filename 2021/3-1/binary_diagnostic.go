package main

import (
  "bufio"
  "fmt"
  "log"
  "math"
  "os"
  "strings"
)

func main() {
  file, err := os.Open("binary_diagnostic.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var countOnes [12]int
  countInput := 0

  for scanner.Scan() {
    inputTxt := scanner.Text()
    inputArr := strings.Split(inputTxt, "")
    for i, val := range inputArr {
      if(val == "1") {
        countOnes[i]++
      }
    }
    countInput++
  }

  gammaRate := 0
  epsilonRate := 0

  for i, val := range countOnes {
    if val >= countInput / 2 {
      gammaRate = gammaRate + int(math.Pow(float64(2), float64(11-i)))
    } else {
      epsilonRate = epsilonRate + int(math.Pow(float64(2), float64(11-i)))
    }
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(gammaRate * epsilonRate)
}
