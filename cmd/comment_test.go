package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_comment(t *testing.T) {
  c := generateComment("test.go", "nit") 
  assert.NotEmpty(t, c)
  assert.Equal(t, "// nit", c)
  c = generateComment("test.py", "nit")
  assert.NotEmpty(t, c)
  assert.Equal(t, "# nit", c)
  c = generateComment("invalid.txt", "nit")
  assert.Empty(t, c)
  c = generateComment("invalid", "nit")
  assert.Empty(t, c)
  c = generateComment("", "")
  assert.Empty(t, c)
}
