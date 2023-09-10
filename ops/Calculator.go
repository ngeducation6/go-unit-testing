// ops/Calculator.go

package ops

import (
    "fmt"
    "calc/interfaces"
    as "calc/mathlib"
)

type Calculator struct {
    interfaces.ICalculator
    addsub interfaces.IAdd
}


// NewCalculator creates a new instance of Calculator with the given IAdd dependency.
// If addsub is nil, a default instance will be created.
func NewCalculator(addsub interfaces.IAdd) Calculator {
    if addsub == nil {
        addsub = as.Addsub{} // Replace with the actual function to create a default instance
    }
    
    return Calculator{
        addsub: addsub,
    }
}



// Add uses the Addsub struct from the mathlib package to add two integers.
func (c Calculator) Add(a, b int) int {
    result := c.addsub.Add(a, b)
    fmt.Printf("Adding %d and %d: %d\n", a, b, result)
    return result
}

// Subtract uses the math library to subtract b from a.
func (c Calculator) Subtract(a, b int) int {
    result := c.addsub.Subtract(a, b)
    fmt.Printf("Subtracting %d from %d: %d\n", b, a, result)
    return result
}

// AddMultiply takes two numbers, calculates the sum using a method from the mathlib package,
// calculates the product of the same two numbers using a method from the mathlib package,
// and finally adds both results using a method from the mathlib package.
func (c Calculator) AddMultiply(a, b int) int {
    // Calculate the sum of the two numbers using a method from the mathlib package
    sum := c.addsub.Add(a, b)

    // Calculate the product of the same two numbers using a method from the mathlib package
    product := as.Multiply(a, b) // Assuming Multiply is defined in mathlib

    // Add the sum and the product using a method from the mathlib package
    result := c.addsub.Add(sum, product)

    msg := "Calculating the sum of %d and %d, the product of %d and %d, and adding them again: %d\n"
    fmt.Printf(msg, a, b, a, b, result)

    return result
}
