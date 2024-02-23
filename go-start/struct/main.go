package main

import "fmt"

// 構造体
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func fn1() {
	var john Employee
	fmt.Println(john) // {0   }

	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee) // {1001 John Doe Doe's Street}
	fmt.Println(employee.FirstName)

	employee2 := Employee{LastName: "Doe", FirstName: "John"} // 明示的にフィールドを指定
	fmt.Println(employee2)
}

// 構造体へのポインターは&演算子 (参照渡しになる)
func fn2() {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}

func main() {
	fn2()
}
