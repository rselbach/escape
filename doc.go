// escape is a utility package to escape/unescape strings using a custom set of
// unwanted characters.
//
// # Introduction
//
// Sometimes you need strings not to contain certain characters, similar to how
// URLs and query strings are escaped.
//
// This package provides an equivalent way to encode and decode strings using a
// custom list of "unwanted" characters.
//
//	g := escape.New(":,$")
//	escaped := g.Escape("foo: $12,34")
//	fmt.Println(escaped) // foo%3a %2412%2c34
//
// Similarly, you can unescape the string:
//
//	unescaped, err := g.Unescape(escaped)
//	fmt.Println(unescaped) // foo: $12,34
//
// # Customization
//
// By default, it uses the % character as the escape sequence marker, but you
// can also specify a custom marker.
//
//	g := escape.NewWithMarker(":,$", '#')
//	escaped := g.Escape("foo: $12,34")
//	fmt.Println(escaped) // foo#3a $#12#2c34
//
// The marker is always escaped, even if it doesn't appear in the list of
// unwanted characters.
package escape
