package basics

import "time"

var name = "Lemuel"
var points int = 25
var dateOfBirth time.Time = time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)

func Variables() (string, int, float64, bool, time.Time) {
	area := 34.65
	isValid := true

	return name, points, area, isValid, dateOfBirth
}
