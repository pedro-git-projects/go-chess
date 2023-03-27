package utils

import "errors"

// CircularLinkedCoords represents a circular linked list
// of chess coordinate letters
type LinkedCoords struct {
	current *Node
	a       *Node
	b       *Node
	c       *Node
	d       *Node
	e       *Node
	f       *Node
	g       *Node
	h       *Node
}

// NewCircularCoordList creates a new circular linked list of chess
// letter coordinates starting at the intial letter parameter
func NewCoordList(initial rune) (*LinkedCoords, error) {
	clc := LinkedCoords{
		a: NewNode('a'),
		b: NewNode('b'),
		c: NewNode('c'),
		d: NewNode('d'),
		e: NewNode('e'),
		f: NewNode('f'),
		g: NewNode('g'),
		h: NewNode('h'),
	}

	clc.a.next = clc.b
	clc.b.next = clc.c
	clc.c.next = clc.d
	clc.d.next = clc.e
	clc.e.next = clc.f
	clc.f.next = clc.g
	clc.g.next = clc.h
	clc.h.next = nil

	clc.a.previous = nil
	clc.b.previous = clc.a
	clc.c.previous = clc.b
	clc.d.previous = clc.c
	clc.e.previous = clc.d
	clc.f.previous = clc.e
	clc.g.previous = clc.f
	clc.h.previous = clc.a

	switch initial {
	case 'a':
		clc.current = clc.a
	case 'b':
		clc.current = clc.b
	case 'c':
		clc.current = clc.c
	case 'd':
		clc.current = clc.d
	case 'e':
		clc.current = clc.e
	case 'f':
		clc.current = clc.f
	case 'g':
		clc.current = clc.g
	case 'h':
		clc.current = clc.h
	default:
		err := errors.New("invalid initial value")
		return nil, err
	}

	return &clc, nil
}

// CurrentValue is an accessor for the current string value
// on the list
func (c LinkedCoords) CurrentValue() rune {
	switch c.current {
	case c.a:
		return 'a'
	case c.b:
		return 'b'
	case c.c:
		return 'c'
	case c.d:
		return 'd'
	case c.e:
		return 'e'
	case c.f:
		return 'f'
	case c.g:
		return 'g'
	case c.h:
		return 'h'
	}
	return *new(rune)
}

func (c LinkedCoords) IsNextNil() bool {
	if c.current.next == nil {
		return true
	}
	return false
}

func (c LinkedCoords) IsPrevNil() bool {
	if c.current.previous == nil {
		return true
	}
	return false
}

func (c *LinkedCoords) Current() *Node {
	return c.current
}

// Next moves the current to the next node
func (c *LinkedCoords) MoveToNext() {
	c.current = c.current.next
}

// Next moves the current to the previous node
func (c *LinkedCoords) MoveToPrev() {
	c.current = c.current.previous
}
