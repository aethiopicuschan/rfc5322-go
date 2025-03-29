package rfc5322

import (
	"fmt"
	"strings"
)

// MessageID represents a Message-ID in the RFC 5322 format.
type MessageID struct {
	left  string
	right string
}

// NewMessageID creates a new MessageID instance with the given value.
func NewMessageID(left, right string) (m *MessageID, err error) {
	if left == "" || right == "" {
		err = ErrorInvalidMessageID
		return
	}
	m = &MessageID{
		left:  left,
		right: right,
	}
	return
}

// String returns the string representation of the MessageID.
func (m *MessageID) String() string {
	return fmt.Sprintf("<%s@%s>", m.left, m.right)
}

// MessageIDs represents a slice of MessageID.
type MessageIDs []MessageID

// NewMessageIDs creates a new MessageIDs instance with the given MessageID values.
func NewMessageIDs(messageIDs ...MessageID) MessageIDs {
	return messageIDs
}

// String returns the string representation of the MessageIDs.
func (m MessageIDs) String() string {
	list := make([]string, 0)
	for _, mi := range m {
		list = append(list, mi.String())
	}
	return strings.Join(list, " ")
}
