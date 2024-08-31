package nits

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_nit(t *testing.T) {
	f := []byte(`
  import kfdjas
  import fdaslkjf

  def main():
    blah
    blah
    blah
    blah
    blah
  `)

	nit := NewNit("f.py", "hash", "hashy", "hashyy")
	assert.NotEmpty(t, nit)

	byt, err := AddNit(f, nit, 0.2)
	assert.NoError(t, err)
	spew.Dump(byt)

}
