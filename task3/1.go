package main

import (
	"fmt"
)

//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

// 查询指定部门的所有员工
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = $1`
	err := db.Select(&employees, query, department)
	return employees, err
}

// 查询薪资最高的员工
func getHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`
	err := db.Get(&employee, query)
	return employee, err
}

func main() {
	techEmployees, _ := getEmployeesByDepartment(db, "技术部")
	highestPaid, _ := getHighestPaidEmployee(db)
	fmt.Println(techEmployees)
	fmt.Println(highestPaid)
}
