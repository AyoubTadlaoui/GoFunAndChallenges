package main

import "fmt"

func main() {
	fmt.Println("=== Variables & Constants ===")
	fmt.Println("Pi =", Pi)
	fmt.Println("Rectangle 3 x 4 area =", RectangleArea(3, 4))
	fmt.Println("Circle r=5 area =", CircleArea(5))
	fmt.Println("Circle r=5 circumference =", CircleCircumference(5))
	fmt.Println("Type of 42 =", DescribeType(42))
	fmt.Println("Type of \"hi\" =", DescribeType("hi"))
	fmt.Println("Type of 3.14 =", DescribeType(3.14))
	fmt.Println("Type of true =", DescribeType(true))
}
