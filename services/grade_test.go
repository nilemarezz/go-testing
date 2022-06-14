package services_test

import (
	"fmt"
	"testing"
	"testing/services"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {
	type TestCase struct {
		name     string
		score    int
		expected string
	}

	cases := []TestCase{
		{name: "Score more than 80", score: 85, expected: "A"},
		{name: "Score more than 70", score: 72, expected: "B"},
		{name: "Score more than 60", score: 69, expected: "C"},
		{name: "Score more than 50", score: 52, expected: "D"},
		{name: "Score less than 50", score: 30, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := services.Checkgrade(c.score)
			assert.Equal(t, c.expected, result)
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.Checkgrade(10)
	}
}

func ExampleCheckGrade() {
	grade := services.Checkgrade(80)
	fmt.Println(grade)
	// Output: A
}
