// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/maps"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// Assert Map implementation
var _ maps.BidiMap[int, int] = (*Map[int, int])(nil)

// Map holds the elements in two red-black trees.
type Map[TKey, TValue comparable] struct {
	forwardMap      redblacktree.Tree[TKey, *data[TKey, TValue]]
	inverseMap      redblacktree.Tree[TValue, *data[TKey, TValue]]
	keyComparator   utils.Comparator
	valueComparator utils.Comparator
}

type data[TKey, TValue comparable] struct {
	key   TKey
	value TValue
}

// NewWith instantiates a bidirectional map.
func NewWith[TKey, TValue comparable](keyComparator utils.Comparator, valueComparator utils.Comparator) *Map[TKey, TValue] {
	return &Map[TKey, TValue]{
		forwardMap:      *redblacktree.NewWith[TKey, *data[TKey, TValue]](keyComparator),
		inverseMap:      *redblacktree.NewWith[TValue, *data[TKey, TValue]](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithIntComparators instantiates a bidirectional map with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithIntComparators[TKey, TValue comparable]() *Map[TKey, TValue] {
	return NewWith[TKey, TValue](utils.IntComparator, utils.IntComparator)
}

// NewWithStringComparators instantiates a bidirectional map with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators[TKey, TValue comparable]() *Map[TKey, TValue] {
	return NewWith[TKey, TValue](utils.StringComparator, utils.StringComparator)
}

// Put inserts element into the map.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d.value)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d.key)
	}
	d := &data[TKey, TValue]{key: key, value: value}
	m.forwardMap.Put(key, d)
	m.inverseMap.Put(value, d)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d.value, true
	}
	return *new(TValue), false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[TKey, TValue]) GetKey(value TValue) (key TKey, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d.key, true
	}
	return *new(TKey), false
}

// Remove removes the element from the map by key.
func (m *Map[TKey, TValue]) Remove(key TKey) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d.value)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[TKey, TValue]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (ordered).
func (m *Map[TKey, TValue]) Keys() []TKey {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
func (m *Map[TKey, TValue]) Values() []TValue {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[TKey, TValue]) String() string {
	str := "TreeBidiMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
