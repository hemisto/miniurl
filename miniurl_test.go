package miniurl_test

import (
	"fmt"
	"testing"

	"github.com/hemisto/miniurl"
	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	const (
		input    = "https://github.com/hemisto/miniurl"
		expected = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expected)
}

func TestHashSameResultWithSameInput(t *testing.T) {
	const (
		input = "https://github.com/hemisto/miniurl"
	)

	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}

func TestHashDifferentResultsWithDifferentInputs(t *testing.T) {
	const (
		input1 = "https://github.com/hemisto/miniurl"
		input2 = "asdfasdf"
	)

	output1 := miniurl.Hash(input1)
	output2 := miniurl.Hash(input2)
	assert.NotEqual(t, output1, output2)
}

func ExampleHash() {
	const (
		input    = "https://github.com/hemisto/miniurl"
		expected = 32
	)

	output := miniurl.Hash(input)
	fmt.Println(output)
	// output:
	// 5f8ae3d142f2dca236f941f71b263949
}

func BenchmarkHash(b *testing.B) {
	const input = "https://github.com/hemisto/miniurl"
	for n := 0; n < b.N; n++ {
		miniurl.Hash(input)
	}
}

func FuzzHash(f *testing.F) {
	f.Add("some string")
	f.Fuzz(func(t *testing.T, input string) {
		output1 := miniurl.Hash(input)
		output2 := miniurl.Hash(input)

		assert.Equal(t, output1, output2)
		assert.Len(t, output1, 32)
		assert.Len(t, output2, 32)
	})
}
