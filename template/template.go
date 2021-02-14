package template

//go:generate zonk tmpl -func FuncMap -format text -encode

import (
	"bytes"
	"io"
	"strings"
	"text/template"
	"unicode"
)

// Execute renders the named template and writes to io.Writer wr.
func Execute(wr io.Writer, name string, data interface{}) error {
	buf := new(bytes.Buffer)
	err := T.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}

	src, err := format(buf)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, src)
	return err
}

// FuncMap provides extra functions for the templates.
var FuncMap = template.FuncMap{
	"substr":   substr,
	"camelize": camelize,
	"hexdump":  hexdump,
}

func substr(s string, i int) string {
	return s[:i]
}

func camelize(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if !isCamelCase(runeValue) {
			continue
		}
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

func isCamelCase(r rune) bool {
	return r == '-' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// hexdump is a template function that creates a hux dump
// similar to xxd -i.
func hexdump(v interface{}) string {
	var data []byte
	switch vv := v.(type) {
	case []byte:
		data = vv
	case string:
		data = []byte(vv)
	default:
		return ""
	}
	var buf bytes.Buffer
	for i, b := range data {
		dst := make([]byte, 4)
		src := []byte{b}
		encode(dst, src, ldigits)
		buf.Write(dst)

		buf.WriteString(",")
		if (i+1)%cols == 0 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

// default number of columns
const cols = 12

// hex lookup table for hex encoding
const (
	ldigits = "0123456789abcdef"
	udigits = "0123456789ABCDEF"
)

func encode(dst, src []byte, hextable string) {
	dst[0] = '0'
	dst[1] = 'x'
	for i, v := range src {
		dst[i+1*2] = hextable[v>>4]
		dst[i+1*2+1] = hextable[v&0x0f]
	}
}
