package rfc5322

import (
	"fmt"
	"mime"
	"strings"

	"github.com/moznion/go-optional"
)

// Header represents the header of an email message as per RFC 5322.
type Header struct {
	// minimum required fields
	date Date
	from Addresses

	// optional fields
	sender     optional.Option[Address]
	to         optional.Option[Addresses]
	cc         optional.Option[Addresses]
	bcc        optional.Option[Addresses]
	messageID  optional.Option[MessageID]
	replyTo    optional.Option[Address]
	inReplyTo  optional.Option[string]
	references optional.Option[MessageIDs]
	subject    optional.Option[string]
	comments   optional.Option[string]
	keywords   optional.Option[[]string]

	// resent fields
	resentDate      optional.Option[Date]
	resentFrom      optional.Option[Addresses]
	resentSender    optional.Option[Address]
	resentTo        optional.Option[Addresses]
	resentCc        optional.Option[Addresses]
	resentBcc       optional.Option[Addresses]
	resentMessageID optional.Option[MessageID]
	resentReplyTo   optional.Option[Address]

	// extra fields
	extra optional.Option[map[string]string]
}

// NewHeader creates a new Header instance with the given date and from addresses.
func NewHeader(date Date, from Addresses) *Header {
	return &Header{
		date: date,
		from: from,
	}
}

func (h *Header) SetSender(sender Address) *Header {
	h.sender = optional.Some(sender)
	return h
}

func (h *Header) AddTo(to Address) *Header {
	if h.to.IsSome() {
		old := h.to.Unwrap()
		h.to = optional.Some(append(old, to))
	} else {
		h.to = optional.Some(Addresses{to})
	}
	return h
}

func (h *Header) AddCc(cc Address) *Header {
	if h.cc.IsSome() {
		old := h.cc.Unwrap()
		h.cc = optional.Some(append(old, cc))
	} else {
		h.cc = optional.Some(Addresses{cc})
	}
	return h
}

func (h *Header) AddBcc(bcc Address) *Header {
	if h.bcc.IsSome() {
		old := h.bcc.Unwrap()
		h.bcc = optional.Some(append(old, bcc))
	} else {
		h.bcc = optional.Some(Addresses{bcc})
	}
	return h
}

func (h *Header) SetReplyTo(replyTo Address) *Header {
	h.replyTo = optional.Some(replyTo)
	return h
}

func (h *Header) SetMessageID(messageID MessageID) *Header {
	h.messageID = optional.Some(messageID)
	return h
}

func (h *Header) SetInReplyTo(inReplyTo string) *Header {
	h.inReplyTo = optional.Some(inReplyTo)
	return h
}

func (h *Header) AddReference(reference MessageID) *Header {
	if h.references.IsSome() {
		old := h.references.Unwrap()
		h.references = optional.Some(append(old, reference))
	} else {
		h.references = optional.Some(MessageIDs{reference})
	}
	return h
}

func (h *Header) SetSubject(subject string) *Header {
	h.subject = optional.Some(subject)
	return h
}

func (h *Header) SetComments(comment string) *Header {
	h.comments = optional.Some(comment)
	return h
}

func (h *Header) AddKeyword(keyword string) *Header {
	if h.keywords.IsSome() {
		old := h.keywords.Unwrap()
		h.keywords = optional.Some(append(old, keyword))
	} else {
		h.keywords = optional.Some([]string{keyword})
	}
	return h
}

func (h *Header) AddKeywords(keywords []string) *Header {
	if h.keywords.IsSome() {
		old := h.keywords.Unwrap()
		h.keywords = optional.Some(append(old, keywords...))
	} else {
		h.keywords = optional.Some(keywords)
	}
	return h
}

func (h *Header) SetResentDate(date Date) *Header {
	h.resentDate = optional.Some(date)
	return h
}

func (h *Header) AddResentFrom(from Address) *Header {
	if h.resentFrom.IsSome() {
		old := h.resentFrom.Unwrap()
		h.resentFrom = optional.Some(append(old, from))
	} else {
		h.resentFrom = optional.Some(Addresses{from})
	}
	return h
}

func (h *Header) SetResentSender(sender Address) *Header {
	h.resentSender = optional.Some(sender)
	return h
}

func (h *Header) AddResentTo(to Address) *Header {
	if h.resentTo.IsSome() {
		old := h.resentTo.Unwrap()
		h.resentTo = optional.Some(append(old, to))
	} else {
		h.resentTo = optional.Some(Addresses{to})
	}
	return h
}

func (h *Header) AddResentCc(cc Address) *Header {
	if h.resentCc.IsSome() {
		old := h.resentCc.Unwrap()
		h.resentCc = optional.Some(append(old, cc))
	} else {
		h.resentCc = optional.Some(Addresses{cc})
	}
	return h
}

func (h *Header) AddResentBcc(bcc Address) *Header {
	if h.resentBcc.IsSome() {
		old := h.resentBcc.Unwrap()
		h.resentBcc = optional.Some(append(old, bcc))
	} else {
		h.resentBcc = optional.Some(Addresses{bcc})
	}
	return h
}

func (h *Header) SetResentMessageID(messageID MessageID) *Header {
	h.resentMessageID = optional.Some(messageID)
	return h
}

func (h *Header) SetResentReplyTo(replyTo Address) *Header {
	h.resentReplyTo = optional.Some(replyTo)
	return h
}

func (h *Header) SetExtra(key, value string) *Header {
	if h.extra.IsSome() {
		old := h.extra.Unwrap()
		old[key] = value
		h.extra = optional.Some(old)
	} else {
		h.extra = optional.Some(map[string]string{key: value})
	}
	return h
}

// String returns the string representation of the Header.
func (h Header) String() (s string, err error) {
	var sb strings.Builder

	// minimum required fields
	sb.WriteString("MIME-Version: 1.0\r\n")
	sb.WriteString(fmt.Sprintf("Date: %s\r\n", h.date.String()))
	addr, err := h.from.String()
	if err != nil {
		return
	}
	sb.WriteString(fmt.Sprintf("From: %s\r\n", addr))

	// If there are multiple addresses in the "From" field, include the "Sender" field
	if len(h.from) > 1 {
		if h.sender.IsSome() {
			sender := h.sender.Unwrap()
			addr, err = sender.String()
			if err != nil {
				return
			}
			sb.WriteString(fmt.Sprintf("Sender: %s\r\n", addr))
		} else {
			err = ErrorNeedSender
			return
		}
	}

	// Check if "To", "Cc", or "Bcc" fields are set
	if h.to.IsSome() || h.cc.IsSome() || h.bcc.IsSome() {
		set := false
		if h.to.IsSome() {
			set = true
			addr, err = h.to.Unwrap().String()
			if err != nil {
				return
			}
			sb.WriteString(fmt.Sprintf("To: %s\r\n", addr))
		}
		if h.cc.IsSome() {
			set = true
			addr, err = h.cc.Unwrap().String()
			if err != nil {
				return
			}
			sb.WriteString(fmt.Sprintf("Cc: %s\r\n", addr))
		}
		if h.bcc.IsSome() {
			set = true
			addr, err = h.bcc.Unwrap().String()
			if err != nil {
				return
			}
			sb.WriteString(fmt.Sprintf("Bcc: %s\r\n", addr))
		}
		if !set {
			err = ErrorNeedToCcBcc
			return
		}
	}

	// Optional fields
	if h.messageID.IsSome() {
		mi := h.messageID.Unwrap()
		sb.WriteString(fmt.Sprintf("Message-ID: %s\r\n", mi.String()))
	}
	if h.replyTo.IsSome() {
		replyTo := h.replyTo.Unwrap()
		addr, err = replyTo.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Reply-To: %s\r\n", addr))
	}
	if h.inReplyTo.IsSome() {
		inReplyTo := h.inReplyTo.Unwrap()
		sb.WriteString(fmt.Sprintf("In-Reply-To: %s\r\n", inReplyTo))
	}
	if h.references.IsSome() {
		references := h.references.Unwrap()
		sb.WriteString(fmt.Sprintf("References: %s\r\n", references.String()))
	}
	if h.subject.IsSome() {
		subject := h.subject.Unwrap()
		if encode {
			subject = mime.QEncoding.Encode("utf-8", subject)
		}
		sb.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	}
	if h.comments.IsSome() {
		comments := h.comments.Unwrap()
		if encode {
			comments = mime.QEncoding.Encode("utf-8", comments)
		}
		sb.WriteString(fmt.Sprintf("Comments: %s\r\n", comments))
	}
	if h.keywords.IsSome() {
		keywords := h.keywords.Unwrap()
		if encode {
			for i, keyword := range keywords {
				keywords[i] = mime.QEncoding.Encode("utf-8", keyword)
			}
		}
		sb.WriteString(fmt.Sprintf("Keywords: %s\r\n", strings.Join(keywords, ", ")))
	}

	if h.resentDate.IsSome() {
		resentDate := h.resentDate.Unwrap()
		sb.WriteString(fmt.Sprintf("Resent-Date: %s\r\n", resentDate.String()))
	}

	if h.resentFrom.IsSome() {
		resentFrom := h.resentFrom.Unwrap()
		addr, err = resentFrom.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-From: %s\r\n", addr))
	}

	if h.resentSender.IsSome() {
		resentSender := h.resentSender.Unwrap()
		addr, err = resentSender.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-Sender: %s\r\n", addr))
	}

	if h.resentTo.IsSome() {
		resentTo := h.resentTo.Unwrap()
		addr, err = resentTo.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-To: %s\r\n", addr))
	}

	if h.resentCc.IsSome() {
		resentCc := h.resentCc.Unwrap()
		addr, err = resentCc.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-Cc: %s\r\n", addr))
	}

	if h.resentBcc.IsSome() {
		resentBcc := h.resentBcc.Unwrap()
		addr, err = resentBcc.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-Bcc: %s\r\n", addr))
	}

	if h.resentMessageID.IsSome() {
		resentMessageID := h.resentMessageID.Unwrap()
		sb.WriteString(fmt.Sprintf("Resent-Message-ID: %s\r\n", resentMessageID.String()))
	}

	if h.resentReplyTo.IsSome() {
		resentReplyTo := h.resentReplyTo.Unwrap()
		addr, err = resentReplyTo.String()
		if err != nil {
			return
		}
		sb.WriteString(fmt.Sprintf("Resent-Reply-To: %s\r\n", addr))
	}

	if h.extra.IsSome() {
		extra := h.extra.Unwrap()
		for key, value := range extra {
			if encode {
				value = mime.QEncoding.Encode("utf-8", value)
			}
			sb.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}

	s = sb.String()
	return
}
