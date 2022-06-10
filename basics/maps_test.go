package basics

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Maps(t *testing.T) {
	t.Run("should translate correct word", func(t *testing.T) {
		w := Dictionary("car")
		assert.Equal(t, "carro", w)

		w = Dictionary("bread")
		assert.Equal(t, "pão", w)

		w = Dictionary("wood")
		assert.Equal(t, "madeira", w)

		w = Dictionary("abcd")
		assert.Equal(t, "abcd not found.", w)
	})

	t.Run("should return error if missing visitors", func(t *testing.T) {
		_, err := LogParser("../tmp/log_err_missing.txt")

		assert.Equal(t, errors.New("wrong input: [hi.com] (line #1)"), err)
	})

	t.Run("should return error if visitors is negative", func(t *testing.T) {
		_, err := LogParser("../tmp/log_err_negative.txt")

		assert.Equal(t, errors.New("wrong input: [google.com -9] (line #4)"), err)
	})

	t.Run("should return error if failed to convert number of visitors", func(t *testing.T) {
		_, err := LogParser("../tmp/log_err_str.txt")

		assert.Equal(t, errors.New("wrong input: [golang.org FOUR] (line #5)"), err)
	})

	t.Run("should return correct sum of site visitors", func(t *testing.T) {
		r, _ := LogParser("../tmp/log.txt")

		assert.Equal(t, 10, r["golang.org"])
		assert.Equal(t, 29, r["google.com"])
		assert.Equal(t, 20, r["hi.com"])
	})
}
