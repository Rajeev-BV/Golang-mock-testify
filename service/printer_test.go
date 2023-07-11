package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPrinter struct {
	mock.Mock
}
type mockOperand struct {
	mock.Mock
}

type mockPriceProvider struct {
	mock.Mock
}

type mockDayOfTheWeekProvider struct {
	mock.Mock
}

func (d *mockPrinter) PrintText(inputText string) error {
	args := d.Called(inputText)
	return args.Error(0)
}

func (d *mockOperand) GetOperand() (string, error) {
	args := d.Called()
	return args.String(0), args.Error(1)
}

func (d *mockPriceProvider) GetPrice(itemNumber string) (int, error) {
	args := d.Called((itemNumber))
	return args.Int(0), args.Error(1)
}

func (d *mockDayOfTheWeekProvider) GetDayOfTheWeek() time.Weekday {
	args := d.Called()
	return time.Weekday(args.Int(0))
}

func TestPrint(t *testing.T) {
	_mockPrinter := mockPrinter{}
	_mockPrinter.On("PrintText", "ABC").Return(nil)

	print := Printer{&_mockPrinter}
	print.Print("abc")

	_mockPrinter.AssertCalled(t, "PrintText", "ABC")
	_mockPrinter.AssertNumberOfCalls(t, "PrintText", 1)
	_mockPrinter.AssertExpectations(t)

}

func TestAdd(t *testing.T) {
	//Arrange
	_mockOperand := mockOperand{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"operator":"+"}`))
	}))
	defer server.Close()
	om := DBOperatorProvider{
		BaseUrl: server.URL,
	}
	result, _ := om.GetOperator(om.BaseUrl)
	_mockOperand.On("GetOperand").Return(result, nil)
	print := Operand{&_mockOperand}
	//Act
	resultAdd := print.operate(2, 3)
	//Assert
	assert.Equal(t, resultAdd, 5)
}

func TestSub(t *testing.T) {
	_mockOperand := mockOperand{}
	_mockOperand.On("GetOperand").Return("-", nil)
	print := Operand{&_mockOperand}
	result := print.operate(5, 3)
	assert.Equal(t, result, 2)
}

func TestOperations(t *testing.T) {
	tests := map[string]struct {
		firstNumber  int
		secondNumber int
		operator     string
		expected     int
	}{
		"10+5": {
			firstNumber:  10,
			secondNumber: 5,
			operator:     "+",
			expected:     15,
		},
		"5-2": {
			firstNumber:  5,
			secondNumber: 2,
			operator:     "-",
			expected:     3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_mockOperand := mockOperand{}
			_mockOperand.On("GetOperand").Return(test.operator, nil)
			print := Operand{&_mockOperand}
			result := print.operate(test.firstNumber, test.secondNumber)
			assert.Equal(t, result, test.expected)
		})
	}

}

func Test_DiscountForAnItem_OnWeekEnds(t *testing.T) {
	var expectedPrice float32 = 450.0
	_mockPriceProvider := mockPriceProvider{}
	_mockDayOfTheWeekProvider := mockDayOfTheWeekProvider{}

	_mockPriceProvider.On("GetPrice", "Laptop").Return(500, nil)
	_mockDayOfTheWeekProvider.On("GetDayOfTheWeek").Return(6, nil)

	_priceProvider := PriceProvider{&_mockPriceProvider, &_mockDayOfTheWeekProvider}
	result := _priceProvider.GetNetPrice("Laptop")
	assert.Equal(t, expectedPrice, result)

}

func Test_DiscountForAnItem_OnWeekDays(t *testing.T) {
	var expectedPrice float32 = 500.0
	_mockPriceProvider := mockPriceProvider{}
	_mockDayOfTheWeekProvider := mockDayOfTheWeekProvider{}

	_mockPriceProvider.On("GetPrice", "Laptop").Return(500, nil)
	_mockDayOfTheWeekProvider.On("GetDayOfTheWeek").Return((2), nil)

	_priceProvider := PriceProvider{&_mockPriceProvider, &_mockDayOfTheWeekProvider}

	result := _priceProvider.GetNetPrice("Laptop")
	assert.Equal(t, expectedPrice, result)

}

func TestGetOperatorFromAPI(t *testing.T) {
	// create a new reader with that JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if r.URL.Path != "/getOperator" {
			t.Errorf("Expected to request '/getOperator', got: %s", r.URL.Path)
		}
		w.Write([]byte(`{"operator":"+"}`))
	}))
	defer server.Close()

	om := DBOperatorProvider{
		BaseUrl: server.URL,
	}
	result, err := om.GetOperator(om.BaseUrl)
	if err != nil {
		t.Error("TestGitHubCallSuccess failed.")
		return
	}
	assert.True(t, len(result) > 0)
	assert.Equal(t, result, "+")
}

//Add 2 numbers
//Sub 2 numbers
//Take opertor from config file
//Introduce stubbing
//Injection
//Parameterize
//Save to DB (Mock)
