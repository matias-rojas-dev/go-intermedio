package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (TEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 20
	ftEmployee.id = 5

	tEmployee := TemporaryEmployee{}

	getMessage(tEmployee)
	getMessage(ftEmployee)

}
