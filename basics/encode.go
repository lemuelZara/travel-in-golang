package basics

import (
	"fmt"
)

func ConvertString(str string) (string, []byte) {
	hex := fmt.Sprintf("%x", str)
	bytes := []byte(str)

	return hex, bytes
}

func SpamMask(str string) []byte {
	size := len(str)
	buf := []byte(str)
	in := false

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	for i := 0; i < size; i++ {
		if len(str[i:]) >= nlink && str[i:i+nlink] == link {
			in = true

			i += nlink
		}

		c := str[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			buf[i] = mask
		}
	}

	return buf
}
