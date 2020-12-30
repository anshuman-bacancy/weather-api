package main

import "fmt"

func print(mat [3][3]int) {
  for i := 0; i < len(mat); i++ {
    for j := 0; j < len(mat); j++ {
      fmt.Print(mat[i][j], " ")
    }
    fmt.Println()
  }
}
func Multiply(m1, m2 [][]int) {
  //result := [len(m1)][len(m2[0])]int{}
  result := [3][3]int{}
sum := 0

  if len(m1[0]) == len(m2) {
    for i := 0; i < 3; i++ {
      for j := 0; j < 3; j++ {
        for k := 0; k < 2; k++ {
          sum += m1[i][k]*m2[k][j]
        }
        result[i][j] = sum
        sum = 0;
      }
    }
    print(result)

  } else {
    fmt.Println("Mismatch rows and columns")
  }
}

func main() {
  m1 := [][]int{
    {3,5},
    {9,2},
    {5,0},
  }

  m2 := [][]int{
    {3,5,5},
    {9,2,7},
  }
  Multiply(m1, m2)
}
