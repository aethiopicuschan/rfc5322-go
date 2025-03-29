package rfc5322_test

import (
	"rfc5322-go"
	"testing"
	"time"

	"github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	addr1, _ := rfc5322.NewAddressWithName("Example User", "example@example.com")
	addr2, _ := rfc5322.NewAddress("example2@example.com")
	date := rfc5322.NewDate(time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC))
	testCases := []struct {
		name        string
		inputAddr   rfc5322.Addresses
		inputSender optional.Option[rfc5322.Address]
		inputDate   *rfc5322.Date
		expectedErr error
	}{
		{
			name:        "multiple addresses without sender",
			inputAddr:   rfc5322.NewAddresses(*addr1, *addr2),
			inputSender: optional.None[rfc5322.Address](),
			inputDate:   date,
			expectedErr: rfc5322.ErrorNeedSender,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			header := rfc5322.NewHeader(*tc.inputDate, tc.inputAddr)
			_, err := header.String()
			assert.EqualError(t, err, tc.expectedErr.Error())
		})
	}
}
