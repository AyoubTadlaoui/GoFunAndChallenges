package error_handling

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ErrorHandlingOne() {
	fmt.Println("------------------ Error Handling ------------------------")
	fmt.Println(" - Write a function that divides two numbers and returns an error if the divisor is zero.")

	var x, y float64
	fmt.Println("Enter the devidend:")
	fmt.Scanf("%f", &x)
	fmt.Println("Enter the divisor:")
	fmt.Scanf("%f", &y)

	division := func(a, b float64) (float64, error) {
		//Error condition:
		if b == 0 {
			return 0, errors.New("ERROR: Division By Zero")
		}

		return a / b, nil
	}

	result, err := division(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f devided by: %.2f = %0.2f\n", x, y, result)

}

func ErrorHandlingTwo() {
	fmt.Println("- Implement a file reading function that returns an error if the file does not exist or cannot be read. ")

	filepath := "README.md"

	// OpenFile opens the file at the specified path and returns a file handle.
	openFile := func(fileP string) (*os.File, error) {

		file, err := os.Open(fileP)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	file, err := openFile(filepath)

	if err != nil {

		return
	}
	defer file.Close()
	fmt.Println("File Openned Successfully ✅")

	//Read the Opened File

	readFile := func(file *os.File) (string, error) {

		content, err := io.ReadAll(file)
		if err != nil {
			return "", err
		}

		return string(content), nil

	}
	fileContent, err := readFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%s\n", fileContent)
}
