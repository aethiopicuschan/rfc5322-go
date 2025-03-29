package rfc5322_test

import (
	"testing"

	"github.com/aethiopicuschan/rfc5322-go"
	"github.com/stretchr/testify/assert"
)

func TestContentType(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expected    string
	}{
		{
			"none header",
			"",
			"text/plain",
		},
		{
			"valid text type",
			"text/html; charset=utf-8",
			"text/html",
		},
		{
			"multipart header",
			"multipart/mixed; boundary=foo",
			"multipart/mixed",
		},
		{
			"invalid header",
			"invalid",
			"invalid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b := rfc5322.NewBody()
			if tc.headerValue != "" {
				b.SetHeader("Content-Type", tc.headerValue)
			}
			assert.Equal(t, tc.expected, b.ContentType())
		})
	}
}

func TestIsMultipart(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expected    bool
	}{
		{
			"ヘッダー未設定",
			"",
			false,
		},
		{
			"テキストタイプ",
			"text/plain",
			false,
		},
		{
			"multipart タイプ",
			"multipart/alternative; boundary=bar",
			true,
		},
		{
			"不正なヘッダー",
			"invalid",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b := rfc5322.NewBody()
			if tc.headerValue != "" {
				b.SetHeader("Content-Type", tc.headerValue)
			}
			assert.Equal(t, tc.expected, b.IsMultipart())
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *rfc5322.Body
		expected string
	}{
		{
			name: "Single part",
			setup: func() *rfc5322.Body {
				b := rfc5322.NewBody()
				b.SetContent([]byte("Hello, World!"))
				return b
			},
			expected: "\r\nHello, World!",
		},
		{
			name: "multipart",
			setup: func() *rfc5322.Body {
				b := rfc5322.NewBody()
				b.SetHeader("Content-Type", "multipart/mixed")
				b.SetContent([]byte("MainContent"))

				part := rfc5322.NewBody()
				part.SetHeader("Content-Type", "text/plain")
				part.SetContent([]byte("PartContent"))
				b.AddPart(part)
				return b
			},
			expected: "Content-Type: multipart/mixed; boundary=\"BOUNDARY-DEFAULT\"\r\n" +
				"\r\n" +
				"MainContent\r\n" +
				"--BOUNDARY-DEFAULT\r\n" +
				"Content-Type: text/plain\r\n" +
				"\r\n" +
				"PartContent\r\n" +
				"--BOUNDARY-DEFAULT--\r\n",
		},
		{
			name: "multipart without part",
			setup: func() *rfc5322.Body {
				b := rfc5322.NewBody()
				b.SetHeader("Content-Type", "multipart/alternative")
				b.SetContent([]byte("AlternativeContent"))
				return b
			},
			expected: "Content-Type: multipart/alternative; boundary=\"BOUNDARY-DEFAULT\"\r\n" +
				"\r\n" +
				"AlternativeContent\r\n" +
				"--BOUNDARY-DEFAULT--\r\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b := tc.setup()
			assert.Equal(t, tc.expected, b.String())
		})
	}
}
