package main

import "fmt"

type employee struct {
	employeeId   string
	employeeName string
	phone        string
	age          int
	salary       float64
}

func main() {
	employee1 := employee{
		employeeId:   "1",
		employeeName: "John",
		phone:        "1234567890",
		age:          25,
		salary:       50000.00,
	}

	fmt.Println("Employee1 = ", employee1)

	employeeList := []employee{
		{
			employeeId:   "2",
			employeeName: "Jane",
			phone:        "0987654321",
			age:          28,
			salary:       60000.00,
		},
		{
			employeeId:   "3",
			employeeName: "Bob",
			phone:        "1112223333",
			age:          35,
			salary:       75000.00,
		},
	}

	employeeList = append(employeeList, employee{
		employeeId:   "4",
		employeeName: "Alice",
		phone:        "4445556666",
		age:          30,
		salary:       80000.00,
	}, employee{
		employeeId:   "5",
		employeeName: "Tom",
		phone:        "7778889999",
		age:          32,
		salary:       70000.00,
	})

	fmt.Println("EmployeeList = ", employeeList)
}
