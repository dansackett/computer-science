package linkedlist

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	Convey("Given an empty Singly Linked List", t, func() {
		empty := NewSingleList()
		Convey("it has size of zero", func() {
			So(empty.Size(), ShouldEqual, 0)
		})
		Convey("it is empty", func() {
			So(empty.IsEmpty(), ShouldBeTrue)
		})
	})
	Convey("Given a Singly Linked List with one node", t, func() {
		single := NewSingleList()
		single.Append(7)
		Convey("it has size of one", func() {
			So(single.Size(), ShouldEqual, 1)
		})
		Convey("it is not empty", func() {
			So(single.IsEmpty(), ShouldBeFalse)
		})
		Convey("it contains the node", func() {
			So(single.Contains(7), ShouldBeTrue)
		})
		Convey("it has a root of 7", func() {
			So(single.GetFirst().Value, ShouldEqual, 7)
		})
		Convey("it is also the last node", func() {
			So(single.GetLast().Value, ShouldEqual, 7)
		})
		Convey("its first node is at index 0", func() {
			So(single.Get(0).Value, ShouldEqual, 7)
			So(single.IndexOf(7), ShouldEqual, 0)
			So(single.LastIndexOf(7), ShouldEqual, 0)
		})
		Convey("its value can be updated", func() {
			single.Set(0, 10)
			So(single.Get(0).Value, ShouldEqual, 10)
			So(single.IndexOf(10), ShouldEqual, 0)
			So(single.IndexOf(7), ShouldEqual, -1)
			So(single.LastIndexOf(7), ShouldEqual, -1)
		})
		Convey("it can be removed", func() {
			single.Remove(10)
			So(single.IndexOf(10), ShouldEqual, -1)
			So(single.GetFirst(), ShouldEqual, nil)
		})
	})
	Convey("Given a Singly Linked List with multiple nodes", t, func() {
		multiple := NewSingleList()

		data := []int{1, 3, 5, 7, 9}
		vals := make([]interface{}, len(data))

		for i, d := range data {
			vals[i] = d
		}

		multiple.AppendMultiple(vals)

		Convey("it has size of 5", func() {
			So(multiple.Size(), ShouldEqual, 5)
		})
		Convey("its root is 1", func() {
			So(multiple.GetFirst().Value, ShouldEqual, 1)
		})
		Convey("its tail is 9", func() {
			So(multiple.GetLast().Value, ShouldEqual, 9)
		})
		Convey("it can be prepended", func() {
			multiple.Prepend(9)
			So(multiple.GetFirst().Value, ShouldEqual, 9)
			So(multiple.Size(), ShouldEqual, 6)
		})
		Convey("it contains our values", func() {
			So(multiple.Contains(1), ShouldBeTrue)
			So(multiple.Contains(3), ShouldBeTrue)
			So(multiple.Contains(5), ShouldBeTrue)
			So(multiple.Contains(7), ShouldBeTrue)
			So(multiple.Contains(9), ShouldBeTrue)
		})
		Convey("it can determine indexes correctly", func() {
			So(multiple.IndexOf(9), ShouldEqual, 0)
			So(multiple.IndexOf(1), ShouldEqual, 1)
			So(multiple.IndexOf(3), ShouldEqual, 2)
			So(multiple.IndexOf(5), ShouldEqual, 3)
			So(multiple.IndexOf(7), ShouldEqual, 4)
			So(multiple.LastIndexOf(9), ShouldEqual, 5)
		})
		Convey("it can remove values", func() {
			So(multiple.Pop().Value, ShouldEqual, 9)
			So(multiple.GetLast().Value, ShouldEqual, 7)
			So(multiple.Contains(9), ShouldBeTrue)
			So(multiple.LeftPop().Value, ShouldEqual, 9)
			So(multiple.GetFirst().Value, ShouldEqual, 1)
			So(multiple.Contains(9), ShouldBeFalse)
			So(multiple.Remove(5).Value, ShouldEqual, 5)
			So(multiple.Contains(5), ShouldBeFalse)
		})
		Convey("it can set values", func() {
			multiple.Set(0, 99)
			multiple.Set(1, 999)
			multiple.Set(2, 9999)
			So(multiple.GetFirst().Value, ShouldEqual, 99)
			So(multiple.Get(1).Value, ShouldEqual, 999)
			So(multiple.GetLast().Value, ShouldEqual, 9999)
		})
		Convey("it can be converted to a slice", func() {
			s := multiple.ToSlice()

			data := []int{99, 999, 9999}
			vals := make([]interface{}, len(data))

			for i, d := range data {
				vals[i] = d
			}

			So(s, ShouldEqual, vals)
		})
	})
}
