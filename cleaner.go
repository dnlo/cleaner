// Package cleaner is for composing functions to clean strings
package cleaner

import (
	"regexp"
	"strings"
)

// StringFunction takes a string and returns a string
type StringFunction func(string) string

// Clean takes a set of StringFunctions that a string will be passed through
func Clean(fns ...StringFunction) StringFunction {
	return func(s string) string {
		for _, v := range fns {
			if v != nil {
				s = v(s)
			}
		}
		return s
	}
}

// Delete takes a set of strings to be deleted
func Delete(s ...string) StringFunction {
	return func(r string) string {
		for _, v := range s {
			r = strings.Replace(r, v, "", -1)
		}
		return r
	}
}

// Replace takes a part of the string to be replaced and what to replace it with
func Replace(old, new string) StringFunction {
	return func(s string) string {
		return strings.Replace(s, old, new, -1)
	}
}

// Extract takes a regexp pattern and returns the first match
func Extract(pat string) StringFunction {
	r := regexp.MustCompile(pat)
	return func(s string) string {
		return r.FindString(s)
	}
}

// CutBefore returns part of the string before the first occurence of d
func CutBefore(d string) StringFunction {
	return func(s string) string {
		return (strings.Split(s, d))[0]
	}
}

// CutAt returns part of the string after the last occurence of d
func CutAt(d string) StringFunction {
	return func(s string) string {
		sl := strings.Split(s, d)
		if len(sl) < 1 {
			return ""
		}
		return sl[len(sl)-1]
	}
}
