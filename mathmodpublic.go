package mathmodpublic

import "fmt"

func Greet(name string) {
	fmt.Printf("Welcome %s", name)
}

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func Sub(x, y int) int {
	return x - y
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}
