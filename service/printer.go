package service

type printer interface {
	PrintText(inputText string) error
}

type Printer struct {
	_printer printer
}

func (printerstruct *Printer) Print(inputText string) error {
	err := printerstruct._printer.PrintText(inputText)
	return err
}
