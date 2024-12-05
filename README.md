# escape
Utility package to escape/unescape strings using a custom set of unwanted characters

## Usage

You can use it to escape unwanted characters from a string. For instance,
if you don't want your string to have spaces, you can create an escape.Escape like so:

```go
esc := escape.New(" ")
fmt.Println(esc.Escape("no spaces allowed"))
// prints no%20spaces%20allowed
```

You can have multiple unwanted characters, e.g.

```go
esc := escape.New(",.:")
fmt.Println(esc.Escape("foo,bar.baz:"))
// Output: foo%2cbar%2ebaz%3a
```