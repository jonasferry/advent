package main

import (
  "fmt"
  "os"
  "strconv"
  )

func main() {
  in := os.Args[1]
  lvl := 0
  
  for i := 0; i < len(in); i++ {
    switch string(in[i]) {
      case "(": lvl++
      case ")": lvl--
      default: panic("Error")
    }
  }
  
  fmt.Println("Santa is on level: " + strconv.Itoa(lvl))
}