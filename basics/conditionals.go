package basics

import (
	"fmt"
	"time"
)

func IsWeekend(t *time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

func GetTypeFromVariable(i interface{}) string {
	switch t := i.(type) {
	case bool:
		return "bool"
	case string:
		return "string"
	case int:
		return "int"
	case float64:
		return "float64"
	default:
		return fmt.Sprintf("Don't know type %T", t)
	}
}
