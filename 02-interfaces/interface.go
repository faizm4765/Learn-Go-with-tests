package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectange struct {
	length  float64
	breadth float64
}

func areaR(rect Rectange) float64 {
	return rect.length * rect.breadth
}

func areaC(circle Circle) float64 {
	radius := circle.radius
	return math.Pi * radius * radius
}

func main() {
	// calculate area of circle
	circle := Circle{7}
	fmt.Println("Radius of circle: ", circle.radius)
	fmt.Println("Area of circle: ", areaC(circle)) // function call
	fmt.Println("Area of circle: ", circle.area()) // method call as it has a receiver type

	// calculate area of rectangle
	rect := Rectange{3, 5}
	fmt.Println("DImensions of rectangle: ", rect.length, rect.breadth)
	fmt.Println("Area of rectangle: ", areaR(rect))
	fmt.Println("Area of rectangle: ", rect.area())

	// here comes interface. As of now in second approach we are calling same methods but with just different receiver types. We want to call it with same receiver type.
	checkArea(circle)
	checkArea(rect)

	// calculate perimeter
	checkPerimeter(circle)
	checkPerimeter(rect)
}

// Now instead of having two different method names we will have just one.

func (rect Rectange) area() float64 {
	fmt.Println("Area of rect")
	return rect.length * rect.breadth
}

func (circle Circle) area() float64 {
	fmt.Println("Area of circle")
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) perimeter() float64 {
	fmt.Println("Perimeter of circle")
	return 2 * math.Pi * circle.radius
}

func (rect Rectange) perimeter() float64 {
	fmt.Println("Perimeter of rectangle")
	return 2 * (rect.length + rect.breadth)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func checkArea(s Shape) {
	ar := s.area()
	fmt.Println("Area of desired shape: ", ar)
}

func checkPerimeter(s Shape) {
	perimeter := s.perimeter()
	fmt.Println("Perimeter of desired shape: ", perimeter)
}
