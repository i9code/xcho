package base

import (
	"regexp"
	"strings"
)

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func Camel(s string) string {
	return toCamelInitCase(s, true)
}

func LowerCamel(s string) string {
	return toCamelInitCase(s, false)
}

func Snake(s string) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		nextIsCapital := false
		if i+1 < len(s) {
			w := s[i+1]
			nextIsCapital = w >= 'A' && w <= 'Z'
		}
		if i > 0 && v >= 'A' && v <= 'Z' && n[len(n)-1] != '_' && !nextIsCapital {
			n += "_" + string(v)
		} else if v == ' ' {
			n += "_"
		} else {
			n = n + string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}

func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}
