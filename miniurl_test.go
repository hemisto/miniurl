package miniurl_test

import (
	"testing"

	"github.com/hemisto/miniurl"
	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	const (
		input    = "https://github.com/hemisto/miniurl"
		expected = 32
	)

	output := miniurl.Hast(input)
	assert.Len(t, output, expected)
}

func TestHashSameResultWithSameInput(t *testing.T) {
	const (
		input = "https://github.com/hemisto/miniurl"
	)

	output1 := miniurl.Hast(input)
	output2 := miniurl.Hast(input)
	assert.Equal(t, output1, output2)
}
