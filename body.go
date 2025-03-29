package rfc5322

import (
	"fmt"
	"mime"
	"strings"
)

// Body represents the body of an email message.
type Body struct {
	headers map[string]string
	content []byte
	parts   []*Body
}

func NewBody() *Body {
	return &Body{
		headers: make(map[string]string),
		content: make([]byte, 0),
		parts:   make([]*Body, 0),
	}
}

func (b *Body) SetHeader(key, value string) {
	b.headers[key] = value
}

func (b *Body) SetContent(content []byte) {
	b.content = content
}

func (b *Body) AddPart(part *Body) {
	b.parts = append(b.parts, part)
}

func (b *Body) ContentType() string {
	if ct, ok := b.headers["Content-Type"]; ok {
		mediaType, _, err := mime.ParseMediaType(ct)
		if err == nil {
			return mediaType
		}
		return ct
	}
	return "text/plain"
}

func (b *Body) IsMultipart() bool {
	return strings.HasPrefix(b.ContentType(), "multipart/")
}

func (b *Body) ensureBoundary() string {
	if !b.IsMultipart() {
		return ""
	}
	ct, ok := b.headers["Content-Type"]
	if !ok {
		return ""
	}
	mediaType, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return ""
	}
	boundary, exists := params["boundary"]
	if !exists || boundary == "" {
		boundary = "BOUNDARY-DEFAULT"
		b.headers["Content-Type"] = fmt.Sprintf("%s; boundary=%q", mediaType, boundary)
	}
	return boundary
}

func (b *Body) String() string {
	var builder strings.Builder

	boundary := b.ensureBoundary()
	for key, value := range b.headers {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	builder.WriteString("\r\n")

	if b.IsMultipart() {
		if len(b.content) > 0 {
			builder.WriteString(string(b.content))
			builder.WriteString("\r\n")
		}

		for _, part := range b.parts {
			builder.WriteString("--" + boundary + "\r\n")
			builder.WriteString(part.String())
			builder.WriteString("\r\n")
		}
		builder.WriteString("--" + boundary + "--\r\n")
	} else {
		builder.WriteString(string(b.content))
	}

	return builder.String()
}
