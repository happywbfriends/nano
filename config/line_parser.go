package config

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errUnterminatedQuote = errors.New("unterminated quote")
	errEmptyKey          = errors.New("empty key")
	errInvalidQuotes     = errors.New("invalid quotes")
)

// Converts line of format `a=1 b=2 c=hello d="hello world"` into map
func ParseConfigLine(s string, makeKeysUpperCase bool) (map[string]string, error) {
	quoted := rune(0)
	tokens := strings.FieldsFunc(s, func(r rune) bool {
		if r == quoted {
			quoted = 0
		} else if quoted == 0 && (r == '"' || r == '\'') {
			quoted = r
		}
		return quoted == 0 && r == ' '
	})
	if quoted != 0 {
		return nil, errUnterminatedQuote
	}

	values := map[string]string{}
	for _, t := range tokens {
		k, v, found := strings.Cut(t, "=")
		if !found {
			return nil, fmt.Errorf("cant split '%s' into 'key=value'", t)
		}
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if k == "" {
			return nil, errEmptyKey
		}

		if strings.HasPrefix(v, `"`) || strings.HasPrefix(v, `'`) {
			if len(v) == 1 {
				return nil, errInvalidQuotes
			}
			if v[0] != v[len(v)-1] { //open and close quotes should be the same
				return nil, errInvalidQuotes
			}
			v = v[1 : len(v)-1]
		}
		if makeKeysUpperCase {
			values[strings.ToUpper(k)] = v
		} else {
			values[k] = v
		}
	}
	return values, nil
}
