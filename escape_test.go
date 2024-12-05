package escape_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rselbach/escape"
)

func Test_Escape(t *testing.T) {
	tests := map[string]struct {
		g    escape.Escape
		s    string
		want string
	}{
		"empty string": {
			g:    escape.New(""),
			s:    "",
			want: "",
		},
		"no escape": {
			g:    escape.New(""),
			s:    "hello",
			want: "hello",
		},
		"escape %": {
			g:    escape.New(""),
			s:    "hello%world",
			want: "hello%25world",
		},
		"escape %%": {
			g:    escape.New(""),
			s:    "hello%%world",
			want: "hello%25%25world",
		},
		"several": {
			g:    escape.New(":,$"),
			s:    "foo: $12,34",
			want: "foo%3a %2412%2c34",
		},
		"spaces": {
			g:    escape.New(" "),
			s:    "foo bar",
			want: "foo%20bar",
		},
		"beginning": {
			g:    escape.New("."),
			s:    ".foo",
			want: "%2efoo",
		},
		"end": {
			g:    escape.New("."),
			s:    "foo.",
			want: "foo%2e",
		},
		"beginning and end": {
			g:    escape.New("."),
			s:    ".foo.",
			want: "%2efoo%2e",
		},
		"with custom marker": {
			g:    escape.NewWithMarker(":", '$'),
			s:    "foo: $12,34",
			want: "foo$3a $2412,34",
		},
		"with custom marker and %": {
			g:    escape.NewWithMarker(":", '$'),
			s:    "foo:%:",
			want: "foo$3a%$3a",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.g.Escape(tt.s); got != tt.want {
				t.Errorf("Escape.Escape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Unescape(t *testing.T) {
	tests := map[string]struct {
		ge      escape.Escape
		s       string
		want    string
		wantErr string
	}{
		"empty string": {
			ge:   escape.New(""),
			s:    "",
			want: "",
		},
		"no escape": {
			ge:   escape.New(""),
			s:    "hello",
			want: "hello",
		},
		"escape %": {
			ge:   escape.New(""),
			s:    "hello%25world",
			want: "hello%world",
		},
		"escape %%": {
			ge:   escape.New(""),
			s:    "hello%25%25world",
			want: "hello%%world",
		},
		"several": {
			ge:   escape.New(":,$"),
			s:    "foo%3a %2412%2c34",
			want: "foo: $12,34",
		},
		"spaces": {
			ge:   escape.New(" "),
			s:    "foo%20bar",
			want: "foo bar",
		},
		"beginning": {
			ge:   escape.New("."),
			s:    "%2efoo",
			want: ".foo",
		},
		"end": {
			ge:   escape.New("."),
			s:    "foo%2e",
			want: "foo.",
		},
		"beginning and end": {
			ge:   escape.New("."),
			s:    "%2efoo%2e",
			want: ".foo.",
		},
		"invalid sequence": {
			ge:      escape.New("."),
			s:       "foo%xx",
			wantErr: "invalid escape sequence at position 3",
		},
		"unfinished sequence 1": {
			ge:      escape.New("."),
			s:       "foo%",
			wantErr: "unfinished escape sequence at position 3",
		},
		"unfinished sequence 2": {
			ge:      escape.New("."),
			s:       "foo%a",
			wantErr: "unfinished escape sequence at position 3",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tc.ge.Unescape(tc.s)
			if tc.wantErr != "" {
				if err == nil {
					t.Errorf("Escape.Unescape() expected error, got nil")
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("Escape.Unescape() error = %v, want %v", err, tc.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Escape.Unescape() unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Escape.Unescape() = %v, want %v", got, tc.want)
			}
		})
	}
}

func ExampleEscape_Escape() {
	e := escape.New("$,")
	fmt.Println(e.Escape("foo: $12,34"))
	// Output: foo: %2412%2c34
}

func ExampleEscape_Escape_customMarker() {
	ge := escape.New(",.:")
	fmt.Println(ge.Escape("foo,bar.baz:"))
	// Output: foo%2cbar%2ebaz%3a
}
