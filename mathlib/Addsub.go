// mathlib/Addsub.go

package mathlib

import (
	"fmt"
	"calc/interfaces"

)

type Addsub struct {
	add interfaces.IAdd
}

// NewAddsub creates a new instance of Addsub.
func NewAddsub(add interfaces.IAdd) Addsub {
    return Addsub{
        add: add,
    }
}

// Add returns the sum of two integers.
func (c Addsub) Add(a, b int) int {
    result := a + b
    fmt.Printf("mathlib: Adding %d and %d gives us %d\n", a, b, result)
    return result
}

// Subtract returns the result of subtracting b from a.
func (c Addsub) Subtract(a, b int) int {
    result := a - b
    fmt.Printf("mathlib: Subtracting %d from %d gives us %d\n", b, a, result)
    return result
}
