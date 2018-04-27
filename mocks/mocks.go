package mocks

import "github.com/prantoran/go-mock-microservice/printer"

type PrintFunc func(string) (*printer.Resp, error)

type MockPrinter struct {
	MockPrintFunc PrintFunc
}

func (m *MockPrinter) SetPrintFunc(f PrintFunc) {
	m.MockPrintFunc = f
}

func (m *MockPrinter) Print(s string) (*printer.Resp, error) {
	return m.MockPrintFunc(s)
}
