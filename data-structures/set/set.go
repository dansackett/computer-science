package set

import (
	"fmt"
	"strings"

	hashtable "../hash-table"
)

// Set is a wrapper around a hash table containing the actual items. To make
// it easy, items must be strings in this example so it conforms to my simple
// hash table structure.
type Set struct {
	items *hashtable.HashTable
}

// NewSet creates an empty Set instance.
func NewSet() *Set {
	return &Set{items: hashtable.NewHashTable()}
}

// Values returns the values from the hash table. In reality it iterates the
// hash table for the keys.
func (s *Set) Values() []string {
	var ret []string

	if s.IsEmpty() {
		return ret
	}

	for _, l := range s.items.Buckets {
		if l != nil {
			for cur := l.GetHead(); cur != nil; cur = cur.Next() {
				ret = append(ret, cur.GetKey())
			}
		}
	}
	return ret
}

// String fullfills the Stringer interface allowing easy set printing.
func (s *Set) String() string {
	return fmt.Sprintf("[%s]", strings.Join(s.Values(), ","))
}

// Cardinality returns the number of items in the set.
func (s *Set) Cardinality() uint64 {
	return s.items.NumEntries
}

// IsEmpty returns an answer based on the current cardinality.
func (s *Set) IsEmpty() bool {
	if s.Cardinality() == 0 {
		return true
	}
	return false
}

// Add inserts a new item into the set.
func (s *Set) Add(key string) {
	s.items.Put(key, true)
}

// Remove attempts to delete an item from the set.
func (s *Set) Remove(key string) {
	s.items.Del(key)
}

// Contains checks if an items exists within the set.
func (s *Set) Contains(key string) bool {
	if s.items.Get(key) != nil {
		return true
	}
	return false
}

// Visit is a method allowing a function to be applied to each item in the set.
func (s *Set) Visit(fn func(string)) {
	for _, v := range s.Values() {
		fn(v)
	}
}

// Subset builds a subset based on a conditional function.
func (s *Set) Subset(condition func(string) bool) *Set {
	newSet := NewSet()
	s.Visit(func(v string) {
		if condition(v) {
			newSet.Add(v)
		}
	})
	return newSet
}

// Union returns a new set containing all of the values from both sets.
func (s *Set) Union(other *Set) *Set {
	newSet := NewSet()
	s.Visit(func(v string) { newSet.Add(v) })
	other.Visit(func(v string) { newSet.Add(v) })
	return newSet
}

// Intersection returns a new set containing the values that overlap between
// both sets.
func (s *Set) Intersection(other *Set) *Set {
	if s.Cardinality() > other.Cardinality() {
		return s.Subset(func(v string) bool { return other.Contains(v) })
	} else {
		return other.Subset(func(v string) bool { return s.Contains(v) })
	}
}

// Difference returns a new set containing its items that the other set does
// not contain.
func (s *Set) Difference(other *Set) *Set {
	return s.Subset(func(v string) bool { return !other.Contains(v) })
}
