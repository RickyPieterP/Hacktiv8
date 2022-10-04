package webserver

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Arieel", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 22, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 21, Division: "GA"},
}

func getEmployees() []Employee {
	return employees
}

func createEmployees(in Employee) []Employee {
	in.ID = len(employees) + 1
	employees = append(employees, in)
	return employees
}
