package main

import (
	"fmt"
	"strconv"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
	Son       *Employee
}

func main() {
	employee1 := EmployeeByID(123)
	employee2 := EmployeeByID(123)

	// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的
	fmt.Println(employee2 == employee1)
}

func EmployeeByID(id int) Employee {
	return Employee{
		ID:        id,
		Name:      "N" + strconv.Itoa(id),
		Address:   "A" + strconv.Itoa(id),
		DoB:       time.Time{},
		Position:  "P" + strconv.Itoa(id),
		Salary:    id * 1000,
		ManagerID: id + 10000,
	}
}
