package main

import "fmt"

func hasUniqueSubSquares(board [][]int) bool {
  s := make([][]int, 0)
  for row := 0; row < len(board); row += 3 {
    for col := 0; col < len(board); col += 3 {
      //fmt.Print(board[row:row+3][col:col+3], " ")
      s = board[row:row+3]
      //k := s[col:col]
      row += 3
    }
    fmt.Println(s, " ")
    fmt.Println()
  }
  return false
}

func checkDuplicate(arr []int) bool {
  for idx := 0; idx < len(arr)-1; idx++ {
    for i := idx+1; i < len(arr); i++ {
      if arr[idx] == arr[i] {
        return true
      }
    }
  }
  return false
}

func getColumn(board [][]int, columnIndex int) (column []int) {
  column = make([]int, 0)
  for _, row := range board {
    column = append(column, row[columnIndex])
  }
  return
}

func sudokuValid(board [][]int) (string, bool) {
  var validBoard bool

  // check if empty 
  if len(board) == 0 { return "Empty board", false }

  allRows := make([]int, 0)
  allCols := make([]int, 0)
  for idx, row := range board {
    allRows = append(allRows, row...)
    IsRowDuplicate := checkDuplicate(row)
    col := getColumn(board, idx)
    IsColumnDuplicate := checkDuplicate(col)
    allCols = append(allCols, col...)
    //fmt.Println(IsColumnDuplicate, IsRowDuplicate)
    if IsRowDuplicate == true || IsColumnDuplicate == true {
      validBoard = false
    } else {
      validBoard = true
    }
  }

  _ = hasUniqueSubSquares(board)

  return "Invalid board!", validBoard
}

func print(board [][]int) {
  for i := 0; i < len(board); i++ {
    //if i == 3 || i == 6 {
      //fmt.Println(" ")
    //}
    for j := 0; j < len(board); j++ {
      //if (j == 3) || (j == 6) {
        //fmt.Print(" ")
      //}
      fmt.Print(board[i][j], " ")
    }
    fmt.Println()
  }
}

func main() {
  board := [][]int{
    {5,3,4,6,7,8,9,1,2},
    {6,7,2,1,9,5,3,4,8},
    {1,9,8,3,4,2,5,6,7},
    {8,5,9,7,6,1,4,2,3},
    {4,2,6,8,5,3,7,9,1},
    {7,1,3,9,2,4,8,5,6},
    {9,6,1,5,3,7,2,8,4},
    {2,8,7,4,1,9,6,3,5},
    {3,4,5,2,8,6,1,7,9},
  }
  fmt.Println("Actual board")
  print(board)
  fmt.Println()
  isValid, message := sudokuValid(board)
  fmt.Println(isValid, message)
}

