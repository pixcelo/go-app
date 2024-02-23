package main

import "fmt"

// type Person struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Address   string
// }

// このように宣言すると、Employeeのフィールドの変更で壊れる可能性がある
// type Employee struct {
// 	Information Person
// 	ManagerID   int
// }

// func fn1() {
// 	var employee Employee
// 	employee.Information.FirstName = "John"
// }

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func fn1() {
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)
}

func main() {
	fn1()
}
