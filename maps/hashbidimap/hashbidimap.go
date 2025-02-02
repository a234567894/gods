// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashbidimap implements a bidirectional map backed by two hashmaps.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package hashbidimap

import (
	"fmt"

	"github.com/a234567894/gods/maps"
	"github.com/a234567894/gods/maps/hashmap"
)

// Assert Map implementation
var _ maps.BidiMap[int, int] = (*Map[int, int])(nil)

// Map holds the elements in two hashmaps.
type Map[TKey, TValue comparable] struct {
	forwardMap hashmap.Map[TKey, TValue]
	inverseMap hashmap.Map[TValue, TKey]
}

// New instantiates a bidirectional map.
func New[TKey, TValue comparable]() *Map[TKey, TValue] {
	return &Map[TKey, TValue]{*hashmap.New[TKey, TValue](), *hashmap.New[TValue, TKey]()}
}

// Put inserts element into the map.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	if valueByKey, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(valueByKey)
	}
	if keyByValue, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(keyByValue)
	}
	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	return m.forwardMap.Get(key)
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[TKey, TValue]) GetKey(value TValue) (key TKey, found bool) {
	return m.inverseMap.Get(value)
}

// Remove removes the element from the map by key.
func (m *Map[TKey, TValue]) Remove(key TKey) {
	if value, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(value)
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

// Keys returns all keys (random order).
func (m *Map[TKey, TValue]) Keys() []TKey {
	return m.forwardMap.Keys()
}

// Values returns all values (random order).
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
	str := "HashBidiMap\n"
	str += fmt.Sprintf("%v", m.forwardMap)
	return str
}
