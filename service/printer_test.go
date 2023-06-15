package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockPrinter struct {
	mock.Mock
}

func (d *mockPrinter) PrintText(inputText string) error {
	args := d.Called(inputText)
	return args.Error(0)
}

func TestPrint(t *testing.T) {
	_mockPrinter := mockPrinter{}
	_mockPrinter.On("PrintText", "abc").Return(nil)

	print := Printer{&_mockPrinter}
	print.Print("abc")

	_mockPrinter.AssertCalled(t, "PrintText", "abc")
	_mockPrinter.AssertNumberOfCalls(t, "PrintText", 1)
	_mockPrinter.AssertExpectations(t)

}
