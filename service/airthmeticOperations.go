package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AdditionOpeartion struct {
	fetchOperatorFromDB DBContext
}

func (add *AdditionOpeartion) Calculate(firstNumber int, secondNumber int) int {
	operatorProvider := OperatorProvider{add.fetchOperatorFromDB}
	result := operatorProvider.fetchOperatorFromDB.getOperator()
	if result == "+" {
		return firstNumber + secondNumber
	} else if result == "-" {
		return firstNumber - secondNumber
	}
	return 0
}

type DBContext interface {
	getOperator() string
}

type OperatorProvider struct {
	fetchOperatorFromDB DBContext
}

type DBOperatorProvider_ struct {
	BaseUrl string
}

func (provideOperator *DBOperatorProvider_) ProvideOperator(url_ string) (string, error) {
	//op := provideOperator.fetchOperatorFromDB.getOperator()
	//return op

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
	operatorHolder := &ValueHolder_{}
	err = json.Unmarshal(content, operatorHolder)
	if err != nil {
		return "", errors.New("invalid input")
	}
	op := operatorHolder.Operator
	return op, nil
}

type ValueHolder_ struct {
	Operator string
}
