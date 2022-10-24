package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	result := Addition(1, 2)
	assert.Equal(t, 3, result, "Result must be 3")
	result = Addition(2, 3)
	assert.Equal(t, 5, result, "Result must be 5")
}

func TestSubtraction(t *testing.T) {
	if Subtraction(4, 2) != 2 {
		t.Error("Expected 4 (-) 2 to equal 2")
	}
	if Subtraction(6, 3) != 3 {
		t.Error("Expected 6 (-) 3 to equal 3")
	}
}

func TestDivision(t *testing.T) {
	if Division(4, 2) != 2 {
		t.Error("Expected 4 (/) 2 to equal 2")
	}
	if Division(9, 3) != 3 {
		t.Error("Expected 9 (/) 3 to equal 3")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(6, 2) != 12 {
		t.Error("Expected 6 (*) 2 to equal 12")
	}
	if Multiplication(2, 5) != 10 {
		t.Error("Expected 2 (*) 5 to equal 10")
	}
}
