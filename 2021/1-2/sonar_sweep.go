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
  sum := -2
  sumA := 0
  sumB := 0
  sumC := 0
  sumD := 0
  flag := -1

  for scanner.Scan() {
    curr, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal(err)
    }
    switch flag = (flag + 1) % 4; flag {
    case 0:
      if sumB > sumA {
        sum++
      }
      sumA = curr
      sumC = sumC + curr
      sumD = sumD + curr
    case 1:
      if sumC > sumB {
        sum++
      }
      sumB = curr
      sumD = sumD + curr
      sumA = sumA + curr
    case 2:
      if sumD > sumC {
        sum++
      }
      sumC = curr
      sumA = sumA + curr
      sumB = sumB + curr
    case 3:
      if sumA > sumD {
        sum++
      }
      sumD = curr
      sumB = sumB + curr
      sumC = sumC + curr
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(sum)
}
