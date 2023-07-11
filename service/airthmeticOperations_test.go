package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDB struct {
	mock.Mock
}

func (d *mockDB) getOperator() string {
	args := d.Called()
	return args.String(0)
}

func TestAddOperation_Given_2_Numbers_Then_Add(t *testing.T) {
	//_mockDB := mockDB{}
	//_mockDB.On("getOperator").Return("+")
	//opertions := AdditionOpeartion{&_mockDB}
	//resultAdd := opertions.Calculate(5, 3)
	//assert.Equal(t, resultAdd, 8)

	//Arrange
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, requestReader *http.Request) {
		if requestReader.URL.Path != "/getOperator" {
			t.Errorf("Expected to request '/getOperator', got: %s", requestReader.URL.Path)
		}
		if requestReader.Method != "GET" {
			t.Errorf("Expected to request 'GET', got: %s", requestReader.Method)
		}
		responseWriter.WriteHeader(200)
		responseWriter.Write([]byte(`{"operator":"+"}`))
	}))
	defer server.Close()
	opertorManager := DBOperatorProvider_{
		BaseUrl: server.URL,
	}
	result, _ := opertorManager.ProvideOperator(opertorManager.BaseUrl)
	_mockDB := mockDB{}
	_mockDB.On("getOperator").Return(result)
	opertions := AdditionOpeartion{&_mockDB}
	//Act
	resultAdd := opertions.Calculate(5, 3)
	//Assert
	assert.Equal(t, resultAdd, 8)
}

func TestAddOpCallFail(t *testing.T) {
	//Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()
	opertorManager := DBOperatorProvider_{
		BaseUrl: server.URL,
	}
	//Act
	_, err := opertorManager.ProvideOperator(opertorManager.BaseUrl)
	if err == nil {
		t.Error("Call failed.")
		return
	}
	//Assert
	assert.Contains(t, "invalid input", err.Error())
}

func TestAddOp1(t *testing.T) {
	tests := map[string]struct {
		firstNumber  int
		secondNumber int
		expected     int
	}{
		"10+5": {
			firstNumber:  10,
			secondNumber: 5,
			expected:     15,
		},
		"5+2": {
			firstNumber:  5,
			secondNumber: 2,
			expected:     7,
		},
		"5-2": {
			firstNumber:  5,
			secondNumber: 2,
			expected:     3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			opertions := AdditionOpeartion{}
			result := opertions.Calculate(test.firstNumber, test.secondNumber)
			assert.Equal(t, result, test.expected)
		})
	}
}
