package escape

import (
	"fmt"
	"strings"
)

// Escape escapes and unespaces strings containing a custom list of
// characters.
type Escape struct {
	// list of characters that we don't want in the output string.
	unwanted []byte
	// character used to signal an escape sequence. % by default.
	marker byte
}

var (
	// Spaces is a pre-configured Escape that escapes spaces.
	Spaces = New(" ")
)

// New creates a new Escape with the given list of characters to be escaped
// using % as the escape sequence marker.
func New(unwanted string) Escape {
	return NewWithMarker(unwanted, '%')
}

// NewWithMarker creates a new Escape with the given list of characters to be
// escaped and a custom escape sequence marker.
func NewWithMarker(unwanted string, marker byte) Escape {
	bs := make([]byte, len(unwanted)+1)
	bs[0] = marker
	for _, c := range unwanted {
		bs = append(bs, byte(c))
	}
	return Escape{
		unwanted: bs,
		marker:   marker,
	}
}

// Escape escapes the given string by replacing all unwanted characters with a
// hex representation.
func (e Escape) Escape(s string) string {
	return e.escapeString(s)
}

func (e Escape) Unescape(s string) (string, error) {
	return e.unescapeString(s)
}

func (e Escape) escapeString(s string) string {
	// count the escapable chars
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if e.isUnwanted(c) {
			n++
		}
	}

	if n == 0 {
		return s
	}

	buf := make([]byte, len(s)+2*n)
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if e.isUnwanted(b) {
			buf[j] = e.marker
			buf[j+1] = hexChars[b>>4]
			buf[j+2] = hexChars[b&0x0f]
			j += 3
			continue
		}
		buf[j] = b
		j++
	}
	return string(buf)
}

func (e Escape) unescapeString(s string) (string, error) {
	var n int
	for i := 0; i < len(s); i++ {
		if s[i] == e.marker {
			if i+2 >= len(s) {
				return "", fmt.Errorf("unfinished escape sequence at position %d", i)
			}
			if !isHex(s[i+1]) || !isHex(s[i+2]) {
				return "", fmt.Errorf("invalid escape sequence at position %d", i)
			}
			n++
			i += 3
		}
	}

	if n == 0 {
		return s, nil
	}

	var buf strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == e.marker {
			buf.WriteByte(hexValue(s[i+1])<<4 | hexValue(s[i+2]))
			i += 2
			continue
		}
		buf.WriteByte(s[i])
	}

	return buf.String(), nil
}

const hexChars = "0123456789abcdef"

func isHex(b byte) bool {
	return '0' <= b && b <= '9' || 'a' <= b && b <= 'f' || 'A' <= b && b <= 'F'
}

func hexValue(b byte) byte {
	switch {
	case '0' <= b && b <= '9':
		return b - '0'
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10
	}
	return 0
}

func (e Escape) isUnwanted(b byte) bool {
	for _, u := range e.unwanted {
		if b == u {
			return true
		}
	}
	return false
}
