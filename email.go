package rfc5322

import "fmt"

// Package rfc5322 provides a simple implementation of RFC 5322 email format.
type EMail struct {
	header *Header
	body   *Body
}

func NewEMail(header *Header, body *Body) *EMail {
	return &EMail{
		header: header,
		body:   body,
	}
}

func (e *EMail) String() (s string, err error) {
	hs, err := e.header.String()
	if err != nil {
		return
	}
	bs := e.body.String()
	s = fmt.Sprintf("%s%s", hs, bs)
	return
}
