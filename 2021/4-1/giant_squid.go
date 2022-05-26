package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "regexp"
  "strings"
  "strconv"
)

func checkBoard(board [5][5]int) bool {
  for row := 0; row < 5; row++ {
    if board[row][0] == -1 {
      bingo := true
      for col := 0; col < 5  && bingo; col++ {
        if board[row][col] != -1 {
          bingo = false
        }
      }
      if bingo {
        return true
      }
    }
  }
  for col := 0; col < 5; col++ {
    if board[0][col] == -1 {
      bingo := true
      for row:= 0; row < 5 && bingo; row++ {
        if board[row][col] != -1 {
          bingo = false
        }
      }
      if bingo {
        return true
      }
    }
  }
  return false
}

func main() {
  file, err := os.Open("giant_squid.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var bingoNums []int

  scanner.Scan()

  for _, num := range strings.Split(scanner.Text(), ",") {
    tmp, _ := strconv.Atoi(num)
    bingoNums = append(bingoNums, tmp)
  }

  var bingoBoards [][5][5]int

  for scanner.Scan() {
    var board[5][5]int
    for row := 0; row < 5; row++ {
      scanner.Scan()
      for col, num := range regexp.MustCompile(`\s+`).Split(scanner.Text(), 5) {
        board[row][col], _ = strconv.Atoi(num)
      }
    }
    bingoBoards = append(bingoBoards, board)
  }

  for _, num := range bingoNums {
    for i := 0; i < len(bingoBoards); i++ {
      for row := 0; row < 5; row++ {
        for col := 0; col < 5; col++ {
          if bingoBoards[i][row][col] == num {
            bingoBoards[i][row][col] = -1
          }
        }
      }
    if checkBoard(bingoBoards[i]) {
      sum := 0
      for row := 0; row < 5; row++ {
        for col := 0; col < 5; col++ {
          if bingoBoards[i][row][col] != -1 {
            sum = sum + bingoBoards[i][row][col]
          }
        }
      }
      fmt.Println(sum * num)
      return
    }
    }
  }
}
