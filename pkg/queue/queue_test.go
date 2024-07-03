package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	queue := New[int]()
	assert.Equal(t, 0, queue.count)
	assert.Equal(t, []int{}, queue.elements)
}

func TestQueuePushMultipleItems(t *testing.T) {
	queue := New[int]()
	queue.Push(2)
	queue.Push(3)

	assert.Equal(t, 2, queue.count)
	assert.Equal(t, []int{2, 3}, queue.elements)
}

func TestQueuePopWithSingleValue(t *testing.T) {
	queue := New[int]()
	queue.Push(2)

	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, 0, queue.count)
	assert.Equal(t, []int{}, queue.elements)
}

func TestQueuePopWithMultipleValues(t *testing.T) {
	queue := New[int]()
	queue.Push(2)
	queue.Push(3)

	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, 1, queue.count)
	assert.Equal(t, []int{3}, queue.elements)

	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, 0, queue.count)
	assert.Equal(t, []int{}, queue.elements)
}

func TestIsEmptyIsTrueWhenNoItemsExist(t *testing.T) {
    queue := New[int]()
    assert.True(t, queue.IsEmpty())

    queue.Push(2)
    queue.Pop()
    assert.True(t, queue.IsEmpty())
}

func TestIsEmptyIsFalseWhenItemsExist(t *testing.T) {
    queue := New[int]()
    queue.Push(2)
    assert.False(t, queue.IsEmpty())
}

