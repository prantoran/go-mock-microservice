package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/prantoran/go-mock-microservice/mocks"
	"github.com/prantoran/go-mock-microservice/printer"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	pr := mocks.NewMockPrinter()
	pr.SetPrintFunc(
		func(s string) (*printer.Resp, error) {
			return &printer.Resp{Body: "hello", Status: http.StatusText(http.StatusOK)}, nil
		})
	printer.Cur = pr
	os.Exit(m.Run())
}

func TestPrint(t *testing.T) {
	var tests = []struct {
		body, status string
	}{
		{"hello", http.StatusText(http.StatusOK)},
		{"notok", http.StatusText(http.StatusNotFound)},
	}
	for _, tt := range tests {
		res, _ := printer.Cur.Print(tt.body)
		if res.Status != tt.status {
			t.Errorf("status got %v, want %v", res.Status, tt.status)

		}
		if res.Body != tt.body {
			t.Errorf("body got %v, want %v", res.Body, tt.body)
		}
	}
}
