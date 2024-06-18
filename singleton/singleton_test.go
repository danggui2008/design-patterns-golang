package singleton_test

import (
	"design-patterns/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.True(t, singleton.GetInstance() == singleton.GetInstance())
	assert.True(t, singleton.GetLazyInstance() == singleton.GetLazyInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test error")
			}
		}
	})
}
