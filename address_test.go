package rfc5322_test

import (
	"testing"

	"github.com/aethiopicuschan/rfc5322-go"
	"github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
)

func TestAddress(t *testing.T) {
	testCases := []struct {
		name        string
		inputAddr   string
		inputName   optional.Option[string]
		addr        string
		expected    string
		expectedErr error
	}{
		{
			name:        "valid email",
			inputAddr:   "example@example.com",
			expected:    "example@example.com",
			expectedErr: nil,
		},
		{
			name:        "invalid email",
			inputAddr:   "invalid-email",
			expected:    "",
			expectedErr: rfc5322.ErrorInvalidAddress,
		},
		{
			name:        "empty email",
			inputAddr:   "",
			expected:    "",
			expectedErr: rfc5322.ErrorInvalidAddress,
		},
		{
			name:        "valid email with name",
			inputAddr:   "example@example.com",
			inputName:   optional.Some("Example User"),
			expected:    "Example User <example@example.com>",
			expectedErr: nil,
		},
		{
			name:        "valid email with empty name",
			inputAddr:   "example@example.com",
			inputName:   optional.Some(""),
			expected:    "",
			expectedErr: rfc5322.ErrorInvalidName,
		},
		{
			name:        "valid email with None name",
			inputAddr:   "example@example.com",
			inputName:   optional.None[string](),
			expected:    "example@example.com",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var address *rfc5322.Address
			var err error
			if tc.inputName.IsSome() {
				address, err = rfc5322.NewAddressWithName(tc.inputName.Unwrap(), tc.inputAddr)
			} else {
				address, err = rfc5322.NewAddress(tc.inputAddr)
			}
			if tc.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, address.Value())
			}
		})
	}
}

func TestAddresses(t *testing.T) {
	addr1, _ := rfc5322.NewAddressWithName("Example User", "example@example.com")
	addr2, _ := rfc5322.NewAddress("example2@example.com")
	testCases := []struct {
		name      string
		inputAddr rfc5322.Addresses
		expected  string
	}{
		{
			name:      "multiple addresses",
			inputAddr: rfc5322.NewAddresses(*addr1, *addr2),
			expected:  "Example User <example@example.com>, example2@example.com",
		},
		{
			name:      "single address",
			inputAddr: rfc5322.NewAddresses(*addr1),
			expected:  "Example User <example@example.com>",
		},
		{
			name:      "no addresses",
			inputAddr: rfc5322.NewAddresses(),
			expected:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.inputAddr.Value())
		})
	}
}
