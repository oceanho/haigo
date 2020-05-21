package haigorr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectorGet(t *testing.T) {
	selector := New(1000)
	for i := 0; i < 1000; i++ {
		assert.Equal(t,int64(i),selector.Next())
	}
	assert.Equal(t,int64(0),selector.Next())
}


func BenchmarkGet(b *testing.B){
	selector := New(int64(b.N))
	for i := 0; i < b.N; i++ {
		assert.Equal(b,int64(i),selector.Next())
	}
	assert.Equal(b,int64(0),selector.Next())
}
