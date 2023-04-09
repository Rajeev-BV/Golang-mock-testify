package service

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
	empSal := getEmployeeDetails(g)
	empDetails := getEmployeeSalaryDetails(g)
	netSal := calculateSalary(empSal)
	err := creditSalaryToBank(g, netSal, empDetails)
	return err
}

func creditSalaryToBank(g FetchSal, netSal int, empDetails empDetails) error {
	err := g.bankAPI.CreditSalaryToEmployee(netSal, empDetails.bankAccNumber)
	return err
}

func calculateSalary(empSal empSal) int {
	netSal := empSal.basic + empSal.HRA
	return netSal
}

func getEmployeeSalaryDetails(g FetchSal) empDetails {
	empDetails, _ := g.database.FetchEmployeeDetails("aa")
	return empDetails
}

func getEmployeeDetails(g FetchSal) empSal {
	empSal, _ := g.database.FetchEmployeeSalaryDetails("aa")
	return empSal
}
