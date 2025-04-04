package main

import (
	"testing"
	"github.com/gabbu98/dsa-go-training/arrays" // Replace with your actual module path
)

func main() {
	var test testing.T
    arrays.TestReverseInLine(&test)
	arrays.TestAddSliceOfTwoNumbers()
}