package main

import (
	"fmt"
	"cmp"
)

type Node[T comparable] struct {
	val	T		
	next	*Node[T]
}

type LinkedList[T comparable] struct {
	head	*Node[T]
}

func (list *LinkedList[T]) Add(value T) {
	if list.head == nil {
		list.head = &Node[T]{val: value}
		return
	}
	
	current := list.head
	for current.next != nil {
		current = current.next
	}
	
	current.next = &Node[T]{val: value}
}

func (list *LinkedList[T]) Insert(value T, position int) {
	if position == 0 {
		newNode := &Node[T]{val: value, next: list.head}
		list.head = newNode
		return
	}
	
	current := list.head
	for i := 0; i < position - 1; i++ {
		if current == nil {
			return
		}
		current = current.next
	}
	
	newNode := &Node[T]{val: value, next: current.next}
	current.next = newNode
}

func (list *LinkedList[T]) Index(value T) int {
	current := list.head
	index := 0
	
	for current != nil {
		if current.val == value {
			return index
		}
		current = current.next
		index++
	}
	
	return -1
}
