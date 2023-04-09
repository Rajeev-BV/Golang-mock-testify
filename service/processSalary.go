package service

// Repo is fake database interface.
type empSal struct {
	basic int
	HRA   int
}

type empDetails struct {
	empID         int
	empName       string
	bankAccNumber string
	ifscCode      string
}

type Repo interface {
	FetchEmployeeSalaryDetails(empID string) (empSal, error)
	FetchEmployeeDetails(empID string) (empDetails, error)
}

type banlAPI interface {
	CreditSalaryToEmployee(Salary int, AccNumber string) error
}

type FetchSal struct {
	database Repo
	bankAPI  banlAPI
}

func (g FetchSal) FetchSalComp() empSal {
	msg, _ := g.database.FetchEmployeeSalaryDetails("aa")
	return msg
}

func (g FetchSal) ProcessSalary() error {
	empSal, _ := g.database.FetchEmployeeSalaryDetails("aa")
	empDetails, _ := g.database.FetchEmployeeDetails("aa")
	netSal := empSal.basic + empSal.HRA
	err := g.bankAPI.CreditSalaryToEmployee(netSal, empDetails.bankAccNumber)
	return err
}
