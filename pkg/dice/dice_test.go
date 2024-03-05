package dice

import (
	"testing"
)

func TestD4(t *testing.T) {
	result := D4()
	if result < 1 || result > 4 {
		t.Errorf("D4() = %v, want a number between 1 and 4", result)
	}
}

func TestD6(t *testing.T) {
	result := D6()
	if result < 1 || result > 6 {
		t.Errorf("D6() = %v, want a number between 1 and 6", result)
	}
}

func TestD8(t *testing.T) {
	result := D8()
	if result < 1 || result > 8 {
		t.Errorf("D8() = %v, want a number between 1 and 8", result)
	}
}

func TestD10(t *testing.T) {
	result := D10()
	if result < 1 || result > 10 {
		t.Errorf("D10() = %v, want a number between 1 and 10", result)
	}
}

func TestD12(t *testing.T) {
	result := D12()
	if result < 1 || result > 12 {
		t.Errorf("D12() = %v, want a number between 1 and 12", result)
	}
}

func TestD20(t *testing.T) {
	result := D20()
	if result < 1 || result > 20 {
		t.Errorf("D20() = %v, want a number between 1 and 20", result)
	}
}

func TestD100(t *testing.T) {
	result := D100()
	if result < 1 || result > 100 {
		t.Errorf("D100() = %v, want a number between 1 and 100", result)
	}
}

func TestAdd(t *testing.T) {
	// Initialize a variable to hold the result
	result := int64(100)

	// Define the value to be added
	valueToAdd := int64(5)

	// Calculate manually
	newVal := result + valueToAdd

	// Call the Add function with the result and valueToAdd
	Add(&result, valueToAdd)

	// Check if the result has been correctly modified
	if result != newVal {
		t.Errorf("Add() = %v, want %v", result, newVal)
	}
}

func TestSubtract(t *testing.T) {
	// Initialize a variable to hold the result
	result := int64(10)

	// Define the value to be subtracted
	valueToSubtract := int64(5)

	// Calculate manually
	newVal := result - valueToSubtract

	// Call the Add function with the result and valueToAdd
	Subtract(&result, valueToSubtract)

	// Check if the result has been correctly modified
	if result != newVal {
		t.Errorf("Subtract() = %v, want %v", result, newVal)
	}
}

func TestParseDiceCmd(t *testing.T) {
	// Test case: "1d4" should return a number between 1 and 4
	result := ParseDiceCmd("1d4")
	if result < 1 || result > 4 {
		t.Errorf("ParseDiceCmd(\"1d4\") = %v, want a number between 1 and 4", result)
	}

	// Test case: "2d6" should return a number between 2 and 12
	result = ParseDiceCmd("2d6")
	if result < 2 || result > 12 {
		t.Errorf("ParseDiceCmd(\"2d6\") = %v, want a number between 2 and 12", result)
	}

	// Test case: "1d20+5" should return a number between 6 and 25
	result = ParseDiceCmd("1d20+5")
	if result < 6 || result > 25 {
		t.Errorf("ParseDiceCmd(\"1d20+5\") = %v, want a number between 6 and 25", result)
	}

	// Test case: "1d20-5" should return a number between -4 and 15
	result = ParseDiceCmd("1d20-5")
	if result < -4 || result > 15 {
		t.Errorf("ParseDiceCmd(\"1d20-5\") = %v, want a number between -4 and 15", result)
	}

	// Test case: Invalid command should return -1
	result = ParseDiceCmd("invalid")
	if result != -1 {
		t.Errorf("ParseDiceCmd(\"invalid\") = %v, want -1", result)
	}
}
