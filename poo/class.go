package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Println(e)

	e.id = 1
	e.name = "Matías"
	fmt.Println(e)
	e.SetId(5)
	fmt.Println(e)
	e.SetName("Ignacio")
	fmt.Println(e)

	fmt.Println("El id es: ", e.GetId())
	fmt.Println("El nombre es: ", e.GetName())

	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}

	fmt.Println(e2)

	e3 := new(Employee)

	e3.id = 4
	e3.name = "Matías 3"

	fmt.Println(*e3)

	e4 := NewEmployee(10, "Mati", true)
	fmt.Println(*e4)

}
