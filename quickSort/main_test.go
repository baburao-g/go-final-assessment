package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortSuccess(t *testing.T) {
	assert := assert.New(t)
	list := []int{7, 8, 9, 1, 2, 5, 3, 4}
	want := []int{1, 2, 3, 4, 5, 7, 8, 9}
	QuickSort(list, 0, len(list)-1)
	fmt.Println("Success Test")
	assert.Equal(want, list)
}

func TestQuickSortFailure(t *testing.T) {
	assert := assert.New(t)
	list := []int{7, 8, 9, 1, 2, 5, 3, 4}
	want := []int{7, 8, 9, 1, 2, 5, 3, 4}
	QuickSort(list, 0, len(list)-1)
	fmt.Println("Failure Test")
	assert.NotEqual(want, list)
}
