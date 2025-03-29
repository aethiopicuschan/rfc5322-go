package rfc5322_test

import (
	"rfc5322-go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageID(t *testing.T) {
	testCases := []struct {
		name        string
		left        string
		right       string
		expected    string
		expectedErr error
	}{
		{
			name:        "valid message ID",
			left:        "example",
			right:       "example.com",
			expected:    "<example@example.com>",
			expectedErr: nil,
		},
		{
			name:        "empty left part",
			left:        "",
			right:       "example.com",
			expected:    "",
			expectedErr: rfc5322.ErrorInvalidMessageID,
		},
		{
			name:        "empty right part",
			left:        "example",
			right:       "",
			expected:    "",
			expectedErr: rfc5322.ErrorInvalidMessageID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			messageID, err := rfc5322.NewMessageID(tc.left, tc.right)
			if tc.expectedErr == nil {
				assert.Equal(t, tc.expected, messageID.String())
			} else {
				assert.Nil(t, messageID)
				assert.EqualError(t, err, tc.expectedErr.Error())
			}
		})
	}
}
