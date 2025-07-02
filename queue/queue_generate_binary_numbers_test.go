package queue

import (
	"math"
	"strings"
	"testing"
)

/*
Create a function with signature func GenerateBinaryNumbers(n int) []string
Given a positive integer like n count from 0 to n in binary
*/

func generateBinaryNumbers(n int) string {
	if n == 0 {
		return "00000000"
	}
	// 128 | 64 | 32 | 16 | 8 | 4 | 2 | 1
	var base int = 2
	var queue []float64 = []float64{}

	// populate the queue
	for i:=0; i<8; i++ {
		var powerResult float64 = math.Pow(float64(base),float64(i))
		queue = append([]float64{powerResult},queue...)
	}

	binary := make([]string, len(queue)); for i:= range binary {
		binary[i] = "0"
	}

	i:=0
	for n>0 {
		subtractor := int(queue[0])
		if subtractor <= n {
			n = n - subtractor
			binary[i] = "1"
		}
		queue = queue[1:]
		i++
	}

	return strings.Join(binary,"")
}

func TestGenerateBinaryNumbers(t *testing.T) {
	tests := []struct {
		n 			int
		expected 	string
	} {
		{0,"00000000"},
		{1, "00000001"},
		{2, "00000010"},
		{3, "00000011"},
		{4, "00000100"},
		{5, "00000101"},
		{6, "00000110"},
		{7, "00000111"},
		{8, "00001000"},
		{9, "00001001"},
		{10, "00001010"},
	}

	for i:=0; i<len(tests); i++ {
		if got := generateBinaryNumbers(tests[i].n); tests[i].expected != got {
			t.Fatalf("Expected %s but got %s for test case: %d", tests[i].expected, got, i)
		}
	}
}