package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayeredNils(t *testing.T) {
	var ld *Layered
	assert.Equal(t, Undraw, ld.GetLayer())
	var ldp *LayeredPoint
	assert.Equal(t, Undraw, ldp.GetLayer())
}
