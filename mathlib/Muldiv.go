// Muldiv.go

package mathlib

import "fmt"

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
    result := a * b
    fmt.Printf("mathlib: Multiplying %d and %d gives us %d\n", a, b, result)
    return result
}

// Divide returns the result of dividing a by b.
func Divide(a, b int) float64 {
    if b == 0 {
        fmt.Println("mathlib: Division by zero is not allowed")
        return 0.0
    }
    result := float64(a) / float64(b)
    fmt.Printf("mathlib: Dividing %d by %d gives us %f\n", a, b, result)
    return result
}
