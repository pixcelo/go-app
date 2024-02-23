package main

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Information Person
	ManagerID   int
}

func fn1() {
	var employee Employee
	employee.Information.FirstName = "John"
}

func main() {
	fn1()
}
