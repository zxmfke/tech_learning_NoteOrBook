package main

import (
	"fmt"
)

type Iterator interface {
	HasNext() bool

	Next() (interface{}, error)
}

type SliceIterator struct {
	Slice []interface{}
	index int
}

func NewSliceIterator(slice []interface{}) *SliceIterator {
	return &SliceIterator{
		Slice: slice,
		index: 0,
	}
}

func (s *SliceIterator) HasNext() bool {
	return s.index != len(s.Slice)
}

func (s *SliceIterator) Next() (interface{}, error) {
	if s.index > len(s.Slice) {
		return nil, fmt.Errorf("out of range")
	}

	element := s.Slice[s.index]
	s.index++

	return element, nil
}

type List struct {
	Value interface{}
	Next  *List
}

type ListIterator struct {
	List    *List
	current *List
}

func NewList(values []interface{}) *List {
	if len(values) == 0 {
		return nil
	}

	root := new(List)
	temp := new(List)
	for index, v := range values {
		list := &List{
			Value: v,
			Next:  nil,
		}

		if index == 0 {
			temp = list
			root = list
		} else {
			temp.Next = list
			temp = temp.Next
		}
	}

	return root
}

func NewListIterator(list *List) (*ListIterator, error) {
	if list == nil {
		return nil, fmt.Errorf("list is nil")
	}
	return &ListIterator{
		List:    list,
		current: list,
	}, nil
}

func (l *ListIterator) HasNext() bool {
	return l.current != nil
}

func (l *ListIterator) Next() (interface{}, error) {
	if l.current== nil {
		return 0, fmt.Errorf("out of range")
	}

	value := l.current.Value
	l.current = l.current.Next
	return value, nil
}
