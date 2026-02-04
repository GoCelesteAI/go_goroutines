package main

import (
  "fmt"
  "time"
)

// sayHello prints a greeting message
func sayHello() {
  fmt.Println("Hello from goroutine!")
}

// countUp counts from 1 to 5
func countUp() {
  for i := 1; i <= 5; i++ {
    fmt.Println("Count:", i)
    time.Sleep(500 * time.Millisecond)
  }
}

func main() {
  fmt.Println("=== Basic Goroutines ===")
  fmt.Println()

  // Start a goroutine with the go keyword
  fmt.Println("Starting goroutine...")
  go sayHello()

  // Main continues immediately
  fmt.Println("Main function continues...")

  // Wait a bit to see the goroutine output
  time.Sleep(100 * time.Millisecond)
  fmt.Println()

  // Start counting goroutine
  fmt.Println("Starting count goroutine...")
  go countUp()

  // Main does other work
  fmt.Println("Main is doing other work...")
  time.Sleep(3 * time.Second)

  fmt.Println()
  fmt.Println("Main function finished!")
}

