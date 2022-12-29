// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	"encoding/json"

	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/utils"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Tree[int, int])(nil)
var _ containers.JSONDeserializer = (*Tree[int, int])(nil)

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[TKey, TValue]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[TKey, TValue]) FromJSON(data []byte) error {
	elements := make(map[TKey]TValue)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (tree *Tree[TKey, TValue]) UnmarshalJSON(bytes []byte) error {
	return tree.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (tree *Tree[TKey, TValue]) MarshalJSON() ([]byte, error) {
	return tree.ToJSON()
}
