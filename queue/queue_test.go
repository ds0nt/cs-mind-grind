package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const chainSize = 3

func TestQueueSimple(t *testing.T) {
	q := NewQueue(chainSize)
	val, ok := q.Dequeue()
	assert.False(t, ok)
	assert.Nil(t, val)

	assert.Equal(t, q.popI, 0)
	assert.Equal(t, q.pushI, 0)
	q.Enqueue(1)
	assert.Equal(t, q.popI, 0)
	assert.Equal(t, q.pushI, 1)
	val, ok = q.Dequeue()
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 1)
	assert.Equal(t, val, 1)
	assert.True(t, ok)

	q.Enqueue(2)
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 2)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 2)

	q.Enqueue(3)
	assert.Equal(t, q.popI, 2)
	assert.Equal(t, q.pushI, 0)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 3)

	q.Enqueue(4)
	assert.Equal(t, q.popI, 0)
	assert.Equal(t, q.pushI, 1)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 4)
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 1)

	q.Enqueue(6)
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 2)
	q.Enqueue(7)
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 0)
	q.Enqueue(8)
	assert.Equal(t, q.popI, 1)
	assert.Equal(t, q.pushI, 1)
	assert.Nil(t, q.next)
	assert.True(t, q.full)

	val, _ = q.Dequeue()
	assert.False(t, q.full)
	assert.Equal(t, val, 6)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 7)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 8)
	assert.Nil(t, q.next)

	// test chaining...
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Nil(t, q.next)
	q.Enqueue(11)
	assert.NotNil(t, q.next)
	q.Enqueue(12)
	q.Enqueue(13)
	assert.Nil(t, q.next.next)
	q.Enqueue(21)
	assert.NotNil(t, q.next.next)
	q.Enqueue(22)
	q.Enqueue(23)
	assert.Nil(t, q.next.next.next)
	q.Enqueue(24)
	assert.NotNil(t, q.next.next.next)

	val, _ = q.Dequeue()
	assert.Equal(t, val, 1)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 2)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 3)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 11)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 12)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 13)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 21)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 22)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 23)
	assert.NotNil(t, q.next)
	val, _ = q.Dequeue()
	assert.Equal(t, val, 24)
	assert.Nil(t, q.next)
}
