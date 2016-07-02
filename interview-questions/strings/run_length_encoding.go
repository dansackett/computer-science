package interview

import (
	"bytes"
	"strconv"
)

// Encode takes a string and compresses sequences into numeric
// representations such as converting `aabcccccc` to `a2bc6`.
func Encode(s string) string {
	if len(s) == 0 {
		return s
	}

	var buffer bytes.Buffer
	last, count := s[0], 1

	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
		} else {
			buffer.WriteString(string(last))
			if count > 1 {
				buffer.WriteString(strconv.Itoa(count))
			}
			last = s[i]
			count = 1
		}
	}

	buffer.WriteString(string(last))
	if count > 1 {
		buffer.WriteString(strconv.Itoa(count))
	}
	return buffer.String()
}

// Decode takes an encoded string and expands the numeric representations into
// repeated characters again.
func Decode(s string) string {
	if len(s) == 0 {
		return s
	}

	var buffer bytes.Buffer
	last := s[0]

	for i := 1; i < len(s); i++ {
		if val, err := strconv.Atoi(string(s[i])); err == nil {
			for i := 1; i < val; i++ {
				buffer.WriteString(string(last))
			}
		} else {
			buffer.WriteString(string(last))
			last = s[i]
		}
	}

	buffer.WriteString(string(last))

	return buffer.String()
}
