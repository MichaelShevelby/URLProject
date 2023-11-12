package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUrlCorrectSpace(t *testing.T) {
	assert.False(t, IsUrlCorrect(" google.com"))
	assert.False(t, IsUrlCorrect("http:// google.com"))
	assert.False(t, IsUrlCorrect("http://google .com"))
	assert.False(t, IsUrlCorrect("http://google. com"))
	assert.False(t, IsUrlCorrect(" https://google.com"))
}

func TestIsUrlCorrectWrongPrefix(t *testing.T) {
	assert.False(t, IsUrlCorrect("google.com"))
	assert.False(t, IsUrlCorrect("http:/google.com"))
	assert.False(t, IsUrlCorrect("htp:/google.com"))
	assert.False(t, IsUrlCorrect("://google.com"))
	assert.True(t, IsUrlCorrect("http://google.com"))
	assert.True(t, IsUrlCorrect("https://google.com"))
}
