package main

import (
  "bufio"
  "fmt"
  "log"
  "math"
  "os"
)

func main() {
  file, err := os.Open("binary_diagnostic.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var oxyVals []string
  var co2Vals []string

  for scanner.Scan() {
    inputTxt := scanner.Text()
    oxyVals = append(oxyVals, inputTxt)
    co2Vals = append(co2Vals, inputTxt)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  countOnes, countInput := 0, 0
  var newOxyVals []string
  var newCo2Vals []string

  for i := 0; len(oxyVals) > 1; i++ {
    for _, val := range oxyVals {
      if val[i] == '1' {
        countOnes++
      }
      countInput++
    }
    if countOnes >= countInput / 2 {
      for _, val := range oxyVals {
        if val[i] == '1' {
          newOxyVals = append(newOxyVals, val)
        }
      }
    } else {
      for _, val := range oxyVals {
        if val[i] == '0' {
          newOxyVals = append(newOxyVals, val)
        }
      }
    }
    oxyVals = make([]string, len(newOxyVals))
    copy(oxyVals, newOxyVals)
    newOxyVals = nil
    countOnes, countInput = 0, 0
  }

  for i:= 0; len(co2Vals) > 1; i++ {
    for _, val := range co2Vals {
      if val[i] == '1' {
        countOnes++
      }
      countInput++
    }
    if countOnes < countInput / 2 {
      for _, val := range co2Vals {
        if val[i] == '1' {
          newCo2Vals = append(newCo2Vals, val)
        }
      }
    } else {
      for _, val := range co2Vals {
        if val[i] == '0' {
          newCo2Vals = append(newCo2Vals, val)
        }
      }
    }
    co2Vals = make([]string, len(newCo2Vals))
    copy(co2Vals, newCo2Vals)
    newCo2Vals = nil
    countOnes, countInput = 0, 0
  }

  oxyVal, co2Val := 0, 0

  for i := 0; i < 12; i++ {
    if oxyVals[0][i] == '1' {
      oxyVal = oxyVal + int(math.Pow(float64(2), float64(11-i)))
    }
    if co2Vals[0][i] == '1' {
      co2Val = co2Val + int(math.Pow(float64(2), float64(11-i)))
    }
  }

  fmt.Println(oxyVal * co2Val)
}
