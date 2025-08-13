package main

import "fmt"

func main() {
	// This is the entry point of the application.
	// You can add your code here to start the application.
	fmt.Println("Hello, World!")

	// Example of a variable declaration and initialization
	var x int = 10
	fmt.Println("Value of x:", x)

	// Example of another variable declaration and initialization
	var y float64 = 20.5
	fmt.Println("Value of y:", y)

	// Example of a simple addition
	var sum int = x + int(y)
	fmt.Println("Sum of x and y:", sum)

	// Example of a conditional statement
	if sum > 20 {
		fmt.Println("The sum is greater than 20.")
	} else {
		fmt.Println("The sum is 20 or less.")
	}

	// Example of a loop
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// show the type of sum
	fmt.Printf("Type of sum: %T\n", sum)

	const pi float64 = 3.14
	fmt.Println("Value of pi:", pi)

	printMe()

	var a int = 10
	var b int = 0
	result, errorMsg := intDivision(a, b)
	fmt.Println("Result of division:", result)
	if errorMsg != nil {
		fmt.Println("Error from division:", errorMsg)
	}

	// define array of size 10 and print it
	var arr [10]int
	fmt.Println("Array of size 10:", arr)
}

func printMe() {
	fmt.Println("This is a function that prints a message.")
}

func intDivision(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return numerator / denominator, nil
}
