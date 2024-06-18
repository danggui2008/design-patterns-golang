package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShape(t *testing.T) {
	shape := Shape{"shape1"}
	shapeClone := shape.Clone()
	assert.True(t, shape.sType == shapeClone.sType)

}
