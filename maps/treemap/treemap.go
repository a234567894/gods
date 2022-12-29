// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/maps"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// Assert Map implementation
var _ maps.Map[int, int] = (*Map[int, int])(nil)

// Map holds the elements in a red-black tree
type Map[TKey, TValue comparable] struct {
	tree *rbt.Tree[TKey, TValue]
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[TKey, TValue comparable](comparator utils.Comparator) *Map[TKey, TValue] {
	return &Map[TKey, TValue]{tree: rbt.NewWith[TKey, TValue](comparator)}
}

// NewWithIntComparator instantiates a tree map with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[TKey, TValue comparable]() *Map[TKey, TValue] {
	return &Map[TKey, TValue]{tree: rbt.NewWithIntComparator[TKey, TValue]()}
}

// NewWithStringComparator instantiates a tree map with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[TKey, TValue comparable]() *Map[TKey, TValue] {
	return &Map[TKey, TValue]{tree: rbt.NewWithStringComparator[TKey, TValue]()}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Remove(key TKey) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[TKey, TValue]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[TKey, TValue]) Keys() []TKey {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[TKey, TValue]) Values() []TValue {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[TKey, TValue]) Min() (key interface{}, value interface{}) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	return nil, nil
}

// Max returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[TKey, TValue]) Max() (key interface{}, value interface{}) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	return nil, nil
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Floor(key TKey) (foundKey TKey, foundValue TValue) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value
	}
	return *new(TKey), *new(TValue)
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Ceiling(key TKey) (foundKey TKey, foundValue TValue) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value
	}
	return *new(TKey), *new(TValue)
}

// String returns a string representation of container
func (m *Map[TKey, TValue]) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}
