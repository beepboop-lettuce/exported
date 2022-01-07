package staff

import "log"

var OverpaidLimit = 75000
var UnderPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverpaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < UnderPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

//this fun has a receiver e. It has access to the values stored in the receiver variable e.
func (e *Office) notVisible() {
	log.Println("Hello, world")
}

//this func has no receiver. Does not have access to any of the values stored in receiver variable e
func myFunction() {
	log.Println("I m a function")
}
