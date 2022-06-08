package basics

import (
	"fmt"
)

func ConvertString(str string) (string, []byte) {
	hex := fmt.Sprintf("%x", str)
	bytes := []byte(str)

	return hex, bytes
}
