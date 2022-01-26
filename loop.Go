package main

import "fmt"

func main() {
  var ulang int
  
  fmt.Println("Selamat Datang Di Perulangan Looping")
   
  fmt.Println("Berapa Kali Anda Ingin Ulang?: ")
  fmt.Scanln(&ulang)
  
  for i := 0; i < 3; i++ {
    fmt.Println("Belajar Go-lang Ke-", i) 
    ulang = ulang - 1
  }
  fmt.Println("Terima Kasih")
  
  /*
    var i int
    var ulang int
    fmt.Println("Berapa Kali Anda Ingin Ulang : ")
    fmt.Scanln(&ulang)
    i = 0
    for i := 1; i<i {
      i = i + 1
      fmt.Println(i, "Belajar Go-lang")
    } */
}
