// ops/calculator_test

package ops

import (
	"testing"
    "calc/ops"
	"calc/mycustomadd"
)


func TestCalculator_Add(t *testing.T) {
    // Create a Calculator with the mock IAdd implementation
    calculator := ops.NewCalculator(mycustomadd.MyCustomAdd{})


    // Test the Add operation
    // result := calculator.AddMultiply(2, 3)
	result := calculator.Add(2, 3)


    // Check if the result matches the expected value
    expected := 2 + 3
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}

