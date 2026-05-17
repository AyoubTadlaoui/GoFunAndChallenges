package main

import "fmt"

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	fmt.Println("=== Methods & Interfaces ===")
	fmt.Println("Rectangle.Area =", r.Area())
	fmt.Println("Circle.Area    =", c.Area())
	fmt.Println("TotalArea      =", TotalArea(r, c))

	var counter Counter
	counter.Inc()
	counter.Inc()
	counter.Inc()
	fmt.Println("Counter        =", counter.Count())
}
