package binarysearchtree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBST(t *testing.T) {
	Convey("Given an empty BST", t, func() {
		empty := NewBST()
		Convey("it has size of zero", func() {
			So(empty.Size(), ShouldEqual, 0)
		})
		Convey("it can insert many elements", func() {
			empty.MassInsert([]int{3, 7, 1, 19, 10})
			So(empty.Size(), ShouldEqual, 5)
			So(empty.Contains(19), ShouldBeTrue)
		})
	})
	Convey("Given a BST with one element", t, func() {
		single := NewBST()
		single.Insert(16)
		Convey("it has size of 1", func() {
			So(single.Size(), ShouldEqual, 1)
		})
		Convey("it contains the element", func() {
			So(single.Contains(16), ShouldBeTrue)
		})
	})
	Convey("Given a BST with many elements", t, func() {
		multiple := NewBST()
		multiple.MassInsert([]int{3, 7, 1, 19, 10})
		Convey("it has size greater than 1", func() {
			So(multiple.Size(), ShouldBeGreaterThan, 1)
		})
		Convey("it contains the elements", func() {
			So(multiple.Contains(3), ShouldBeTrue)
			So(multiple.Contains(7), ShouldBeTrue)
			So(multiple.Contains(1), ShouldBeTrue)
			So(multiple.Contains(19), ShouldBeTrue)
			So(multiple.Contains(10), ShouldBeTrue)
		})
		Convey("it can be balanced", func() {
			multiple.Balance()
			So(multiple.Root.Key, ShouldEqual, 7)
		})
		Convey("it can delete elements", func() {
			multiple.Remove(19)
			So(multiple.Contains(19), ShouldBeFalse)
			So(multiple.Size(), ShouldEqual, 4)
		})
	})
}
