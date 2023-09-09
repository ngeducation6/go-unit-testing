package interfaces

// Define the interfaces.ICalculator interface
type ICalculator interface {
    Add(a, b int) int
    Subtract(a, b int) int
    AddMultiply(a, b int) int
}