package printer

import (
	"io/ioutil"
	"net/http"
)

type Printer interface {
	Print(s string) (*Resp, error)
}

type Resp struct {
	Body, Status string
}

type HomePrinter struct {
	URL string
}

func (p *HomePrinter) Print(s string) (*Resp, error) {
	addr := p.URL + "?v=" + s
	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &Resp{Body: string(bodyBytes), Status: res.Status}, nil
}

func New(URL string) *HomePrinter {
	return &HomePrinter{URL}
}
