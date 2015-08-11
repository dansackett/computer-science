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
}
