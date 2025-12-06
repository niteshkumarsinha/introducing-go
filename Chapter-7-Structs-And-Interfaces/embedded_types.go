package main


import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Talk(){
	fmt.Println("Hi, my name is: ", p.Name)
}

type Employee struct {
	Person
	EmployeeID string
	Position   string
}

type Android struct {
	Person
	Model string
}

func embedded_types_example() {
	e := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
		Position:   "Software Engineer",
	}

	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Employee ID: %s\n", e.EmployeeID)
	fmt.Printf("Position: %s\n", e.Position)
}

func main() {
	embedded_types_example()

	p := Person{Name: "Nitesh", Age: 35}
	fmt.Println(p.Name)

	a := new(Android)
	a.Person = Person{Name: "Nitesh Kumar", Age: 35}
	a.Person.Talk()
	a.Talk()

}