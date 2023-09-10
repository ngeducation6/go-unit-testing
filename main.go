// main.go

package main

import (
    "fmt"
    "calc/ops"
	"calc/mycustomadd"
)

func main() {

    // Create a new instance of Calculator without providing an Addsub instance.
    ca := ops.NewCalculator(nil) // No Addsub provided, will use default

    // Perform some calculations.
    sum := ca.Add(5, 3)
    difference := ca.Subtract(10, 4)
    addmul := ca.AddMultiply(3, 5)

    // Display the results.
    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", difference)
    fmt.Printf("AddMultiply: %d\n", addmul)

 
	ca = ops.NewCalculator(mycustomadd.MyCustomAdd{})

    // Perform some calculations.
    sum = ca.Add(5, 3)
    // difference := ca.Subtract(10, 4)
    // addmul := ca.AddMultiply(3, 5)

    // Display the results.
    fmt.Printf("Sum: %d\n", sum)
}
