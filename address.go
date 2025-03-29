package rfc5322

import (
	"fmt"
	"mime"
	"net/mail"
	"strings"

	"github.com/moznion/go-optional"
	"golang.org/x/net/idna"
)

// Address represents an email address as per RFC 5322.
type Address struct {
	name  optional.Option[string]
	value string
}

// NewAddress creates a new Address instance with the given value.
func NewAddress(value string) (a *Address, err error) {
	_, err = mail.ParseAddress(value) // Validate the email address format
	if err != nil {
		err = ErrorInvalidAddress
		return
	}
	a = &Address{
		name:  optional.None[string](),
		value: value,
	}
	return
}

// NewAddressWithName creates a new Address instance with the given name and value.
func NewAddressWithName(name, value string) (a *Address, err error) {
	if name == "" {
		err = ErrorInvalidName
		return
	}
	_, err = mail.ParseAddress(value) // Validate the email address format
	if err != nil {
		err = ErrorInvalidAddress
		return
	}
	a = &Address{
		name:  optional.Some(name),
		value: value,
	}
	return
}

// Value returns the value of the Address.
func (a *Address) Value() string {
	if a.name.IsSome() {
		return fmt.Sprintf("%s <%s>", a.name.Unwrap(), a.value)
	} else {
		return a.value
	}
}

// String returns the string representation of the Address.
func (a *Address) String() (s string, err error) {
	if a.name.IsSome() {
		var name string
		if encode {
			name = mime.QEncoding.Encode("utf-8", a.name.Unwrap())
			s, err = idna.ToASCII(a.value)
			if err != nil {
				return
			}
		} else {
			s = a.value
			name = a.name.Unwrap()
		}
		s = fmt.Sprintf("%s <%s>", name, s)
		return
	} else {
		if encode {
			s, err = idna.ToASCII(a.value)
		}
	}
	return
}

// Addresses represents a slice of Address.
type Addresses []Address

// NewAddresses creates a new Addresses instance.
func NewAddresses(addr ...Address) Addresses {
	return Addresses(addr)
}

// Value returns the value of the Addresses.
func (a Addresses) Value() string {
	list := make([]string, 0)
	for _, addr := range a {
		list = append(list, addr.Value())
	}
	return strings.Join(list, ", ")
}

// String returns the string representation of the Addresses.
func (a Addresses) String() (s string, err error) {
	list := make([]string, 0)
	for _, addr := range a {
		addr, err := addr.String()
		if err != nil {
			return "", err
		}
		list = append(list, addr)
	}
	s = strings.Join(list, ", ")
	return
}
