package lerp

import (
	"github.com/Lundis/go-gmath/vec2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLerp(t *testing.T) {
	a := float32(0)
	b := float32(10)
	res := Lerp(a, b, 0.5)

	assert.Equal(t, float32(5), res)
}

func TestLerp2(t *testing.T) {
	a := vec2.F{0, 0}
	b := vec2.F{10, 10}
	res := Lerp2(a, b, 0.5)

	assert.Equal(t, float32(5), res.X)
	assert.Equal(t, float32(5), res.Y)
}
