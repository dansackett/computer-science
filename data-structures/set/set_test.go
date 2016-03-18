package set

import (
	"strings"
	"testing"
)

func TestEmptySet(t *testing.T) {
	s := NewSet()

	if s.Cardinality() != 0 {
		t.Error("Empty set should contain no items")
	}

	if !s.IsEmpty() {
		t.Error("Empty set should be empty")
	}
}

func TestCanGetValues(t *testing.T) {
	s := NewSet()
	s.Add("one")

	if s.Cardinality() != 1 {
		t.Error("Set should have one item")
	}

	if s.IsEmpty() {
		t.Error("Set should not be empty")
	}

	if s.Values()[0] != "one" {
		t.Error("Values were not correct")
	}
}

func TestCanPrintSet(t *testing.T) {
	st := "[two,one]"

	s := NewSet()
	s.Add("one")
	s.Add("two")

	if s.String() != st {
		t.Errorf("Printed string does not match, should be %s but found %s", st, s.String())
	}
}

func TestCanAddItem(t *testing.T) {
	s := NewSet()
	s.Add("one")

	if !s.Contains("one") {
		t.Error("Didn't find 'one' in set")
	}
}

func TestCanDeleteItem(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Remove("one")

	if s.Contains("one") {
		t.Error("Found 'one' in set after deleting it")
	}
}

func TestCanVisitEachItem(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Add("two")
	o := NewSet()

	s.Visit(func(v string) {
		o.Add(strings.ToUpper(v))
	})

	if !o.Contains("ONE") {
		t.Error("Could not find uppercased 'one' in set")
	}

	if !o.Contains("TWO") {
		t.Error("Could not find uppercased 'two' in set")
	}
}

func TestCanGetSubset(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Add("two")
	sub := s.Subset(func(v string) bool { return v != "one" })

	if sub.Contains("one") {
		t.Error("Subset should not contain 'one'")
	}
}

func TestCanGetUnion(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Add("two")

	o := NewSet()
	o.Add("one")
	o.Add("three")

	u := s.Union(o)

	if u.Cardinality() != 3 {
		t.Error("Union should have three members")
	}
}

func TestCanGetIntersection(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Add("two")

	o := NewSet()
	o.Add("one")
	o.Add("three")

	u := s.Intersection(o)

	if u.Cardinality() != 1 {
		t.Error("Intersection should have one member")
	}

	if u.Values()[0] != "one" {
		t.Error("Intersection value should be 'one'")
	}
}

func TestCanGetDifference(t *testing.T) {
	s := NewSet()
	s.Add("one")
	s.Add("two")

	o := NewSet()
	o.Add("one")
	o.Add("three")

	if s.Difference(o).Values()[0] != "two" {
		t.Errorf("Difference of first to second should be 'two', got %s", s.Difference(o).Values()[0])
	}

	if o.Difference(s).Values()[0] != "three" {
		t.Errorf("Difference of second to first should be 'three', got %s", o.Difference(s).Values()[0])
	}
}
