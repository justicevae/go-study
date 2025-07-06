package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Println("员工信息:")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("工号: %s\n", e.EmployeeID)
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "哈哈哈",
			Age:  20,
		},
		EmployeeID: "11111",
	}
	emp.PrintInfo()
}
