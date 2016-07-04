package observer

import "fmt"

type Observer interface {
	Update()
}

type Subject struct {
	observers []Observer
	state     int
}

func NewSubject() *Subject {
	var observers []Observer
	return &Subject{observers: observers}
}

func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAll()
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) NotifyAll() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

type BinaryObserver struct {
	subject *Subject
}

func NewBinaryObserver(subject *Subject) {
	subject.Attach(&BinaryObserver{subject: subject})
}

func (o *BinaryObserver) Update() {
	fmt.Printf("%b", o.subject.GetState())
	fmt.Println()
}

type HexObserver struct {
	subject *Subject
}

func NewHexObserver(subject *Subject) {
	subject.Attach(&HexObserver{subject: subject})
}

func (o *HexObserver) Update() {
	fmt.Printf("%x", o.subject.GetState())
	fmt.Println()
}
