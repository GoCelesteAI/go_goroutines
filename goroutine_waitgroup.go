package main

import (
  "fmt"
  "sync"
  "time"
)

func worker(id int, wg *sync.WaitGroup) {
  defer wg.Done() // Signal completion when function exits

  fmt.Printf("Worker %d starting\n", id)
  time.Sleep(time.Duration(id*200) * time.Millisecond)
  fmt.Printf("Worker %d done\n", id)
}

func main() {
  fmt.Println("=== WaitGroup for Coordination ===")
  fmt.Println()

  // Create a WaitGroup
  var wg sync.WaitGroup

  // Start 5 workers
  fmt.Println("Starting 5 workers...")
  fmt.Println()

  for i := 1; i <= 5; i++ {
    wg.Add(1) // Increment counter before starting goroutine
    go worker(i, &wg)
  }

  // Wait for all workers to complete
  fmt.Println("Waiting for workers to finish...")
  wg.Wait()

  fmt.Println()
  fmt.Println("All workers completed!")
  fmt.Println()

  // WaitGroup with anonymous function
  fmt.Println("Using WaitGroup with anonymous goroutines...")
  fmt.Println()

  var wg2 sync.WaitGroup

  tasks := []string{"Download", "Process", "Upload"}
  for _, task := range tasks {
    wg2.Add(1)
    go func(t string) {
      defer wg2.Done()
      fmt.Printf("Task '%s' started\n", t)
      time.Sleep(500 * time.Millisecond)
      fmt.Printf("Task '%s' completed\n", t)
    }(task)
  }

  wg2.Wait()
  fmt.Println()
  fmt.Println("All tasks finished!")
}

