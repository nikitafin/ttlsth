package main

import (
  "fmt"
)


func main(){
  var year, month, day int
  fmt.Println("Enter birthday(year-month-day):")
  fmt.Scanf("%d-%d-%d", &year, &month, &day)
  switch {
    case year >= 1946 && year <= 1964:
      fmt.Println("Hello boomer")
    case year >= 1965 && year <= 1980:
      fmt.Println("Hello X")
    case year >= 1981 && year <= 1996:
      fmt.Println("Hello mil")
    case year >= 1997 && year <= 2012:
      fmt.Println("Hello zoomer")
    case year >= 2013:
      fmt.Println("Hello alpha")
    default:
      fmt.Println("Hello there")
  }
}
