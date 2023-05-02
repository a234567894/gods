// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/a234567894/gods/containers"
	rbt "github.com/a234567894/gods/trees/redblacktree"
)

// Assert Enumerable implementation
var _ containers.EnumerableWithKey[int, int] = (*Map[int, int])(nil)

// Each calls the given function once for each element, passing that element's key and value.
func (m *Map[TKey, TValue]) Each(f func(key TKey, value TValue)) {
	iterator := m.Iterator()
	for iterator.Next() {
		f(iterator.Key(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a container
// containing the values returned by the given function as key/value pairs.
func (m *Map[TKey, TValue]) Map(f func(key1 TKey, value1 TValue) (TKey, TValue)) *Map[TKey, TValue] {
	newMap := &Map[TKey, TValue]{tree: rbt.NewWith[TKey, TValue](m.tree.Comparator)}
	iterator := m.Iterator()
	for iterator.Next() {
		key2, value2 := f(iterator.Key(), iterator.Value())
		newMap.Put(key2, value2)
	}
	return newMap
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (m *Map[TKey, TValue]) Select(f func(key TKey, value TValue) bool) *Map[TKey, TValue] {
	newMap := &Map[TKey, TValue]{tree: rbt.NewWith[TKey, TValue](m.tree.Comparator)}
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			newMap.Put(iterator.Key(), iterator.Value())
		}
	}
	return newMap
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (m *Map[TKey, TValue]) Any(f func(key TKey, value TValue) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (m *Map[TKey, TValue]) All(f func(key TKey, value TValue) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if !f(iterator.Key(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil,nil otherwise if no element
// matches the criteria.
func (m *Map[TKey, TValue]) Find(f func(key TKey, value TValue) bool) (TKey, TValue) {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return iterator.Key(), iterator.Value()
		}
	}
	return *new(TKey), *new(TValue)
}
