package main

import (
  "fmt"
  "github.com/lukebond/martian-robots/pkg/instructions"
)

func main() {
  i, err := instructions.LoadInstructionsFromFile("sample-input.txt")
  if err != nil {
    fmt.Printf("Error loading instructions: %v\n", err)
  } else {
    output := i.ProcessInstructions()
    fmt.Printf(output)
  }
}
