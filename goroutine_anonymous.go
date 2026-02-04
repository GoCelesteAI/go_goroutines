package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("=== Anonymous Goroutines ===")
  fmt.Println()

  // Anonymous goroutine - function defined inline
  go func() {
    fmt.Println("Hello from anonymous goroutine!")
  }()

  time.Sleep(100 * time.Millisecond)
  fmt.Println()

  // Anonymous goroutine with parameters
  message := "Go is awesome!"
  go func(msg string) {
    fmt.Println("Message:", msg)
  }(message)

  time.Sleep(100 * time.Millisecond)
  fmt.Println()

  // Multiple anonymous goroutines
  fmt.Println("Starting multiple anonymous goroutines...")
  for i := 1; i <= 5; i++ {
    go func(num int) {
      fmt.Printf("Goroutine %d running\n", num)
      time.Sleep(time.Duration(num*100) * time.Millisecond)
      fmt.Printf("Goroutine %d done\n", num)
    }(i)
  }

  // Wait for all to complete
  time.Sleep(1 * time.Second)

  fmt.Println()
  fmt.Println("All anonymous goroutines finished!")
}

