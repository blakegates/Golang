package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30_000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50_000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60_000, FullTime: true},
	{FirstName: "Alan", LastName: "Anderson", Salary: 15_000, FullTime: false},
	{FirstName: "Margeret", LastName: "Carter", Salary: 100_000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}
	// myStaff.All()

	// log.Println(myStaff.All())
	// staff.OverPaidLimit = 10_000
	log.Println("Overpaid staff", myStaff.OverPaid())
	log.Println("Underpaid staff", myStaff.Underpaid())

	myStaff.NotVisible()
}
