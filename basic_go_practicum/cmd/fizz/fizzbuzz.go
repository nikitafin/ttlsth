package main

import "fmt"

func main()  {
  for i:= 1; i < 101; i++{
    switch {
      case i % 3 == 0 && i % 5 == 0:
        fmt.Printf("%d: FizzBuzz", i)
      case i % 3 == 0:
        fmt.Printf("%d: Fizz", i)
      case i % 5 == 0:
        fmt.Printf("%d: Buzz", i)
      default:
        fmt.Printf("%d: %d", i, i)
    }
    fmt.Print("\n")
  }
}
