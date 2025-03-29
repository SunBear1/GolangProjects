package main

import (
	"errors"
	"fmt"
)

const StackMaxSize = 10

type Stack[T any] struct {
	stack    [StackMaxSize]T
	stackPtr int
}

func NewStack[T any]() *Stack[T] {
	stack := &Stack[T]{stackPtr: -1}
	return stack
}

func (s *Stack[T]) push(element T) error {
	if s.stackPtr >= StackMaxSize {
		return errors.New("stack is full")
	} else {
		s.stackPtr++
		s.stack[s.stackPtr] = element
		return nil
	}
}

func (s *Stack[T]) pop() (T, error) {
	if s.stackPtr == -1 {
		var nothing T
		return nothing, errors.New("stack is empty")
	} else {
		value := s.stack[s.stackPtr]
		s.stackPtr--
		return value, nil
	}

}

func (s *Stack[T]) peak() (T, error) {
	if s.stackPtr == -1 {
		var nothing T
		return nothing, errors.New("stack is empty")
	} else {
		return s.stack[s.stackPtr], nil
	}

}

func (s *Stack[T]) show() {
	for i := 0; i < s.stackPtr; i++ {
		fmt.Printf("%s,", s.stack[i])
	}
}
