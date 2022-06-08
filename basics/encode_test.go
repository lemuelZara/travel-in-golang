package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encoding(t *testing.T) {
	t.Run("should represent in binary, decimal and hexadecimal", func(t *testing.T) {
		letters := "ab"

		assert.Equal(t, letters, string([]byte{0b01100001, 0b01100010}))
		assert.Equal(t, letters, string([]byte{97, 98}))
		assert.Equal(t, letters, string([]byte{0x61, 0x62}))
	})

	t.Run("should convert string correctly", func(t *testing.T) {
		str := "á😋"
		hex, bytes := ConvertString(str)

		assert.Equal(t, "c3a1f09f988b", hex)
		assert.Equal(t, []byte{0xc3, 0xa1, 0xf0, 0x9f, 0x98, 0x8b}, bytes)
		assert.Equal(t, "á", string(str[:2]))
	})
}
