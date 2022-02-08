package main

import "fmt"

func main() {
  fmt.Println("Selamat Datang Di Program Array Sederhana")
  fmt.Println("=========================================")
  var array = [9][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}}
  
  for i := 0; i < 9; i++{
    for j := 0; j < 2; i++{
      
      fmt.Printf("[%d][%d][%d] = %d\n", i, i, j, array [i][j])
    }
  }
}
