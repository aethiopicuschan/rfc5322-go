# rfc5322-go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/rfc5322-go.svg)](https://pkg.go.dev/github.com/aethiopicuschan/rfc5322-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/rfc5322-go)](https://goreportcard.com/report/github.com/aethiopicuschan/rfc5322-go)
[![CI](https://github.com/aethiopicuschan/rfc5322-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/rfc5322-go/actions/workflows/ci.yaml)

Builder for [RFC 5322](https://tools.ietf.org/html/rfc5322) compliant email text in Go.

## Installation

```bash
go get -u github.com/aethiopicuschan/rfc5322-go
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/aethiopicuschan/rfc5322-go"
)

func main() {
	// Address
	alice, _ := rfc5322.NewAddressWithName("Alice", "alice@example.com")
	bob, _ := rfc5322.NewAddressWithName("Bob", "bob@example.com")

	// Header
	now := time.Now()
	header := rfc5322.NewHeader(*rfc5322.NewDate(now), rfc5322.NewAddresses(*alice))
	header.AddTo(*bob) // AddTo() is used to add a recipient
	mi, _ := rfc5322.NewMessageID(now.Format("20060102150405"), "example.com")
	header.SetMessageID(*mi)
	header.SetSubject("Hello World!")

	// Body
	body := rfc5322.NewBody()
	body.SetHeader("Content-Type", "multipart/alternative")
	plainBody := rfc5322.NewBody()
	plainBody.SetContent([]byte("Hi Bob!"))
	plainBody.SetHeader("Content-Type", "text/plain; charset=UTF-8")
	htmlBody := rfc5322.NewBody()
	htmlBody.SetContent([]byte("<html><body>Hi Bob!</body></html>"))
	htmlBody.SetHeader("Content-Type", "text/html; charset=UTF-8")
	body.AddPart(plainBody)
	body.AddPart(htmlBody)

	// Email
	email := rfc5322.NewEMail(header, body)
	s, _ := email.String()
	fmt.Println(s)
	/*Output:
	MIME-Version: 1.0
	Date: Sat, 29 Mar 2025 14:33:59 +0900
	From: Alice <alice@example.com>
	To: Bob <bob@example.com>
	Message-ID: <20250329143359@example.com>
	Subject: Hello World!
	Content-Type: multipart/alternative; boundary="BOUNDARY-DEFAULT"

	--BOUNDARY-DEFAULT
	Content-Type: text/plain; charset=UTF-8

	Hi Bob!
	--BOUNDARY-DEFAULT
	Content-Type: text/html; charset=UTF-8

	<html><body>Hi Bob!</body></html>
	--BOUNDARY-DEFAULT--
	*/
}
```

## Features

If you want to disable encoding, you can use `rfc5322.DisableEncode()`.

## Testing

```bash
go test ./...
```
