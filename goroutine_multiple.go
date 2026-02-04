package main

import (
  "fmt"
  "time"
)

// printNumbers prints numbers with a label
func printNumbers(label string, count int) {
  for i := 1; i <= count; i++ {
    fmt.Printf("%s: %d\n", label, i)
    time.Sleep(300 * time.Millisecond)
  }
}

// printLetters prints letters
func printLetters() {
  letters := []string{"A", "B", "C", "D", "E"}
  for _, letter := range letters {
    fmt.Println("Letter:", letter)
    time.Sleep(400 * time.Millisecond)
  }
}

func main() {
  fmt.Println("=== Multiple Goroutines ===")
  fmt.Println()

  // Start multiple goroutines
  fmt.Println("Starting 3 concurrent goroutines...")
  fmt.Println()

  go printNumbers("First", 5)
  go printNumbers("Second", 5)
  go printLetters()

  // Wait for all to complete
  time.Sleep(3 * time.Second)

  fmt.Println()
  fmt.Println("All goroutines finished!")
}

