// mycustomadd/customadd.go

package mycustomadd

import (
    "fmt"
    "calc/interfaces"
)

// MyCustomAdd is a custom implementation of the IAdd interface.
type MyCustomAdd struct{
    interfaces.IAdd
}

// Add returns the sum of two integers using the custom logic.
func (m MyCustomAdd) Add(a, b int) int {
    result := a + b 
    fmt.Printf("Custom Add: Adding %d and %d gives us %d\n", a, b, result)
    return result
}

// Add returns the sum of two integers using the custom logic.
func (m MyCustomAdd) Subtract(a, b int) int {
    // result := a + b
    fmt.Printf("Stubbed subtract\n")
    return 0
}
