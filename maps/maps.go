// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps provides an abstract Map interface.
//
// In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears just once in the collection.
//
// Operations associated with this data type allow:
// - the addition of a pair to the collection
// - the removal of a pair from the collection
// - the modification of an existing pair
// - the lookup of a value associated with a particular key
//
// Reference: https://en.wikipedia.org/wiki/Associative_array
package maps

import "github.com/a234567894/gods/containers"

// Map interface that all maps implement
type Map[TKey, TValue comparable] interface {
	Put(key TKey, value TValue)
	Get(key TKey) (value TValue, found bool)
	Remove(key TKey)
	Keys() []TKey

	containers.Container[TValue]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[TKey, TValue comparable] interface {
	GetKey(value TKey) (key TKey, found bool)

	Map[TKey, TValue]
}
