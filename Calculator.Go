package main

import (
  "fmt"
)

func main() {
  var tanya string
  var operator string
  var op_1 float64
  var op_2 float64
  var operand float64
  var result float64
  
  tanya = "y"
  result = 0
  operator = ""

  fmt.Println("Selamat Datang di Aplikasi Kalkulator Sederhana")
  fmt.Println("-----------------------------------------------")

  for tanya != "n" {
    fmt.Print("Apakah Anda Ingin Melakukan Perhitungan ( y / n): ")
    fmt.Scanln(&tanya)

    if tanya == "y" {
      fmt.Println("----------------------------------------------")
      fmt.Print("Input op_1 : ")
      fmt.Scanln(&op_1)
      fmt.Print("Input operator : ")
      fmt.Scanln(&operator)
      fmt.Print("input op_2 : ")
      fmt.Scanln(&op_2)

      if operator == "+" {
        operand = op_1 + op_2
      } else if operator == "-" {
        operand = op_1 - op_2
      } else if operator == "*" {
        operand = op_1 * op_2
      } else if operator == "/" {
        operand = op_1 / op_2
      } else if operator == "^" {
        operand = op_1
        
        for result > op_2 {
          operand = op_1 * op_2
          result++
        }
      } else {
        fmt.Print("Operator Tidak Terdefinisi, Harap Mulai Dari Awal.")
      }

      fmt.Print("Hasil Perhitungan ", )
      
      fmt.Print(op_1)
      fmt.Print(operator)
      fmt.Print(op_2)
      fmt.Print(" = ")
      fmt.Println(operand)
      fmt.Println("---------------------------------------------------")

    } else if tanya != "n" {
      fmt.Println("Inputan Hanya -Y- Untuk YES Dan -N- Untuk NOT")

    }
  }
  fmt.Println("Terima Kasih")
}
