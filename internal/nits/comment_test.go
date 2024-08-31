package nits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_comment(t *testing.T) {
	c := GenerateComment("test.go", "nit")
	assert.NotEmpty(t, c)
	assert.Equal(t, "// nit", c)
	c = GenerateComment("test.py", "nit")
	assert.NotEmpty(t, c)
	assert.Equal(t, "# nit", c)
	c = GenerateComment("invalid.txt", "nit")
	assert.Empty(t, c)
	c = GenerateComment("invalid", "nit")
	assert.Empty(t, c)
	c = GenerateComment("", "")
	assert.Empty(t, c)
}
