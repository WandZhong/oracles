package dbu

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"

	bat "github.com/robert-zaremba/go-bat"
)

// EmptyArray represents a value of empty Postgresql Array
var EmptyArray = []byte("{}")

var (
	openingArray   = byte('{')
	closingArray   = byte('}')
	arraySeparator = byte(',')

	arraySeparatorSlice = []byte{arraySeparator}
)

// parseArray parses array returned by postgres for []Text column. This implementation
// is based on observed behaviour of postrgresql regarding syntax of output for []Text columns.
//
// Array consists of tokens (entries), separated by commas. Token can be
// either quoted (`"mary had a little lamb"`) or unquoted (`mary`).
//
// Unquoted token does not contain commas, spaces nor quotes. This implementation
// is lax when it comes to parsing unquoted tokens. The only restriction is that
// unquoted token cannot contain comma.
//
// Quoted token does not contain quote (it is always escaped). Other escape symbols may be used as well, but
// as such they should not appear because we never send them. So we will never send "\\n", only "\n".  The
// only exception is backslash itself, which is also always escaped.
//
// Sample valid inputs: `{mary}`, `{"mary"}`, `{"mary had a little lamb"}`, `{"\""}`, `{"\n"}`.
//
// Parsing is done in loop, token by token.
//
// This function is tested in `strings_test.go` (focuses on general correctness) and `pgu_strings_test.go`
// (focuses on integration with postgresql).
func parseArray(source string) ([]string, error) {
	// return empty array and not nil, because web client cannot handle nil
	tokens := make([]string, 0, 0)
	source = strings.Trim(source, "{}")
	for len(source) > 0 {
		token, remaining, err := parseToken(source)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
		source = remaining
	}
	return tokens, nil
}

func parseToken(source string) (string, string, error) {
	rune, _ := utf8.DecodeRuneInString(source)
	if rune == '"' {
		return parseQuotedToken(source[1:])
	}
	return parseUnquotedToken(source)
}

func parseUnquotedToken(source string) (string, string, error) {
	commaPos := strings.IndexRune(source, ',')
	if commaPos == -1 {
		return source, "", nil
	}
	tail := source[(commaPos + 1):]
	if tail == "" {
		return "", "", strconv.ErrSyntax
	}
	return source[:commaPos], source[(commaPos + 1):], nil
}

func parseQuotedToken(source string) (string, string, error) {
	var token []byte
	var runeTmp [utf8.UTFMax]byte

	if len(source) == 0 {
		return "", source, strconv.ErrSyntax
	}
	for source[0] != '"' {
		c, multibyte, tail, err := strconv.UnquoteChar(source, '"')
		if err != nil {
			return "", source, err
		}

		// append rune to buffer
		if c < utf8.RuneSelf || !multibyte {
			token = append(token, byte(c))
		} else {
			n := utf8.EncodeRune(runeTmp[:], c)
			token = append(token, runeTmp[:n]...)
		}

		// move source to the new tail
		source = tail
		if len(source) == 0 {
			return "", source, strconv.ErrSyntax
		}
	}

	//trim tail so that it starts with next token
	if len(source) == 1 {
		source = ""
	} else {
		if source[1] != ',' {
			return "", source, strconv.ErrSyntax
		}
		source = source[2:]
	}

	return bat.UnsafeByteArrayToStr(token), source, nil
}

// SplitArray splits Postgresql encoded Array into list of bytes of elements.
// It trims {} characters and split by ','
func SplitArray(src []byte) [][]byte {
	l := len(src)
	if l < 2 {
		return [][]byte{}
	}
	src = src[1 : l-1]
	return bytes.Split(src, arraySeparatorSlice)
}

// SplitNestedArray splits Postgresql encoded Array of simple types which doesn't require any
// escape charaters into list of bytes of elements.
func SplitNestedArray(src []byte) [][]byte {
	var resp = [][]byte{}
	l := len(src)
	if l < 2 {
		return resp
	}
	src = src[1 : l-1]
	var level, start int
	for i, c := range src {
		if c == openingArray {
			level++
		} else if c == closingArray {
			level--
		} else if level == 0 && c == arraySeparator {
			resp = append(resp, src[start:i])
			start = i + 1
		}
	}
	if l != 2 { // add last element if array is not empty
		resp = append(resp, src[start:])
	}
	return resp
}

// ParseFloatArray parses float array column
func ParseFloatArray(src []byte) ([]float64, error) {
	if bytes.Equal(src, EmptyArray) {
		return []float64{}, nil
	}
	vals := SplitArray(src)
	var results = make([]float64, len(vals))
	var err error
	for i := range vals {
		if results[i], err = bat.Atof64(bat.UnsafeByteArrayToStr(vals[i])); err != nil {
			return nil, err
		}
	}
	return results, nil
}

// ParseInt64Array parses float array column
func ParseInt64Array(src []byte) ([]int64, error) {
	if bytes.Equal(src, EmptyArray) {
		return []int64{}, nil
	}
	vals := SplitArray(src)
	var results = make([]int64, len(vals))
	var err error
	for i := range vals {
		if results[i], err = bat.Atoi64(bat.UnsafeByteArrayToStr(vals[i])); err != nil {
			return nil, err
		}
	}
	return results, nil
}
