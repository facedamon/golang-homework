package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct{}

type Circle struct{}

func (*Rectangle) Area() {
	fmt.Println("Rectangle Area called.")
}

func (*Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter called.")
}

func (*Circle) Area() {
	fmt.Println("Circle Area called.")
}

func (*Circle) Perimeter() {
	fmt.Println("Circle Perimeter called.")
}

type Person struct {
	Name string
	Age  int8
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("name=%s, age=%d, id=%d\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	var s Shape
	s = new(Rectangle)
	s.Perimeter()
	s.Area()

	s = new(Circle)
	s.Perimeter()
	s.Area()

	em := Employee{Person{"jack", 25}, 12}
	em.PrintInfo()
}
