package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type printer interface {
	PrintText(inputText string) error
}

type operand interface {
	GetOperand() (string, error)
}

type operatorType interface {
	Calculate() int
}

type priceProvider interface {
	GetPrice(itemNumber string) (int, error)
}

type dayOfTheWeekProvider interface {
	GetDayOfTheWeek() time.Weekday
}

type AddOperation struct {
	firstNumber  int
	secondNumber int
}

func (add *AddOperation) Calculate() int {
	return add.firstNumber + add.secondNumber
}

type SubOperation struct {
	firstNumber  int
	secondNumber int
}

func (add *SubOperation) Calculate() int {
	return add.firstNumber - add.secondNumber
}

func NewOperator(kind string, firstNumber int, secondNumber int) operatorType {
	if kind == "+" {
		return &AddOperation{firstNumber: firstNumber, secondNumber: secondNumber}
	}
	return &SubOperation{firstNumber: firstNumber, secondNumber: secondNumber}
}

type Printer struct {
	_printer printer
}

type Operand struct {
	_opernd operand
}

type PriceProvider struct {
	_priceProvider        priceProvider
	_dayOfTheWeekProvider dayOfTheWeekProvider
}

func (printerstruct *Printer) Print(inputText string) error {
	var outputText = strings.ToUpper(inputText)
	err := printerstruct._printer.PrintText(outputText)
	return err
}

func (operandstruct *Operand) operate(firstNumber int, secondNumber int) int {
	var result int
	msg, _ := operandstruct._opernd.GetOperand()
	operator := NewOperator(msg, firstNumber, secondNumber)
	result = operator.Calculate()

	return result
}

func (priceProvider *PriceProvider) GetNetPrice(itemName string) float32 {
	var netPrice = 0
	result, _ := priceProvider._priceProvider.GetPrice(itemName)
	dayOfTheWeek := priceProvider._dayOfTheWeekProvider.GetDayOfTheWeek()
	if dayOfTheWeek.String() == time.Saturday.String() || dayOfTheWeek.String() == time.Sunday.String() {
		netPrice = (result * 90) / 100
	} else {
		netPrice = (result)
	}

	return float32(netPrice)
}

type DBOperatorProvider struct {
	BaseUrl string
	//Client  http.Client
}

func (dbOperateProvider *DBOperatorProvider) GetOperator(url_ string) (string, error) {
	//dbOperator := DBOperatorProvider{}
	url := fmt.Sprintf("%s/getOperator", url_)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	client := &http.Client{}
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	v := &ValueHolder{}
	err = json.Unmarshal(content, v)
	if err != nil {
		return "", err
	}
	return v.Operator, nil
}

type ValueHolder struct {
	Operator string
}
