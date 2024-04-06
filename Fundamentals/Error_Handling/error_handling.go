package error_handling

import (
	"errors"
	"fmt"
)

func ErrorHandling() {
	fmt.Println("------------------ Error Handling ------------------------")
	fmt.Println(" - Write a function that divides two numbers and returns an error if the divisor is zero.")

	var x, y float64

	fmt.Println("Enter the Devident")
	fmt.Scanf("%f", &x)

	fmt.Println("Enter the Devisor")
	fmt.Scanf("%f", &y)

	devide := func(a, b float64) (float64, error) {
		if b == 0 {
			return -1, errors.New("You CANNOT DEVIDE BY ZERO!!! 😡")
		}

		return a / b, nil
	}

	result, err := devide(x, y)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f devided by %.2f is %.2f\n", x, y, result)
}
