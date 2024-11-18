package main

import "fmt"

// mplmeneted implicitly
type employee interface {
	getName() string
	getSalary() int
	getGrade(role string) (grade string)
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getGrade(r string) (grade string) {
	// grade calculation based on salary
	if ft.salary < 50000 {
		return "A"
	} else if ft.salary < 70000 {
		return "B"
	} else {
		return "C"
	}

}

func main() {
	emp := fullTime{name: "namesa", salary: 60000}
	fmt.Print(emp)
}
