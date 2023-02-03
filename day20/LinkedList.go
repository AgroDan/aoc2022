package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	coord int
	prev  *LinkedList
	next  *LinkedList
}

func (l *LinkedList) NextStep() (*LinkedList, error) {
	if l.next == nil {
		return nil, errors.New("dangling list entry")
	}

	return l.next, nil
}

func NewLinkedList(c int) LinkedList {
	return LinkedList{
		coord: c,
		prev:  nil,
		next:  nil,
	}
}

func UpdateAfter(setNode, movedNode *LinkedList) error {
	// this function UPDATES the movedNode item
	// to the new place in the chain AFTER setNode

	if setNode.next == nil || setNode.prev == nil || movedNode.next == nil || movedNode.prev == nil {
		return errors.New("cannot update dangling nodes")
	}

	// first, bind the old chains
	movedNode.prev.next = movedNode.next
	movedNode.next.prev = movedNode.prev

	// Now set the movedNode to the new chain
	movedNode.prev = setNode
	movedNode.next = setNode.next

	// now sever and bind the new chains
	setNode.next.prev = movedNode
	setNode.next = movedNode

	return nil
}

func UpdateBefore(setNode, movedNode *LinkedList) error {
	// similar to update after, this moves the movedNode to
	// before the setnode

	if setNode.next == nil || setNode.prev == nil || movedNode.next == nil || movedNode.prev == nil {
		return errors.New("cannot update dangling nodes")
	}

	// first bind the old chains
	movedNode.prev.next = movedNode.next
	movedNode.next.prev = movedNode.prev

	// now set the movedNode to the new chain
	movedNode.prev = setNode.prev
	movedNode.next = setNode

	// now sever and bind the new chains
	setNode.prev.next = movedNode
	setNode.prev = movedNode

	return nil
}

func InsertAfter(orig, newMember *LinkedList) {
	// this function adds a linked list item to the
	// list AFTER the given listMember address.
	if orig == newMember {
		return
	}

	if orig.next == nil {
		// this is a current list, not a single item. If we
		// got here, the list is dangling
		newMember.prev = orig
		newMember.next = orig

		orig.prev = newMember
		orig.next = newMember
	} else {
		// otherwise the list item is dangling or something,
		// signifying that this is a list of one item or it's
		// just incorrectly created
		newMember.prev = orig
		newMember.next = orig.next
		orig.next.prev = newMember
		orig.next = newMember
	}
}

func InsertBefore(orig, newMember *LinkedList) {
	// as the name implies, this inserts an item
	// BEFORE the chosen listMember
	if orig == newMember {
		return
	}
	if orig.prev == nil {
		newMember.prev = orig
		newMember.next = orig

		orig.next = newMember
		orig.prev = newMember
	} else {
		newMember.prev = orig.prev
		newMember.next = orig

		orig.prev.next = newMember
		orig.prev = newMember

	}
}

func FindCoord(c int, l *LinkedList) *LinkedList {
	// returns the address of the specific linked list
	// item for indexing. This will return the first
	// instance it finds.

	current := l

	for current.next != l {
		if current.coord == c {
			return current
		}
		current = current.next
	}
	if current.coord == c {
		return current
	}
	return nil
}

func StepTraverse(howMany int, l *LinkedList) *LinkedList {
	// this function will ignore the "step over itself" line
	// in the Traverse function and simply move howMany items
	// in the proposed direction. This is necessary for part 2
	// since we are obtaining a modulus to determine the proper
	// direction, which itself accounts for the "moving" item
	current := l
	if howMany == 0 {
		return current
	}
	for i := 0; i < howMany; i++ {
		current = current.next
	}

	return current
}

func Traverse(howMany int, l *LinkedList) *LinkedList {
	// moves through the linked list howMany times and
	// returns the item it lands on. If you are using
	// this function to find out where to insert something,
	// it is up to the caller to insert after or insert before.
	// you can determine that based on whether the provided coord
	// is positive or negative. pos = after, neg = before.
	current := l
	if howMany > 0 {
		for i := 0; i < howMany; i++ {
			if current.next == l {
				// if the next item has wrapped around the entire list back
				// to itself, skip over it.
				current = current.next
			}
			current = current.next
		}
	} else if howMany < 0 {
		for i := 0; i > howMany; i-- {
			if current.prev == l {
				current = current.prev
			}
			current = current.prev
		}
	}
	return current
}

// formula: take current number and position as
// (Num + Pos) % len(nums-1)

func PrintLL(l *LinkedList, debug bool) {
	fmt.Printf("[ ")
	ptr := l
	for ptr.next != l {
		if debug {
			fmt.Printf("Item: %d, Addr: %p, Contents: %#v \n", ptr.coord, ptr, ptr)
		} else {
			fmt.Printf("<-> %d ", ptr.coord)
		}
		ptr = ptr.next
	}
	if debug {
		fmt.Printf("Item: %d, Addr: %p, Contents: %#v \n", ptr.coord, ptr, ptr)
	} else {
		fmt.Printf("<-> %d ", ptr.coord)
	}
	fmt.Printf("]\n")
}

func GetAddresses(l *LinkedList) []*LinkedList {
	// this function returns a slice of the LinkedList objects
	// in the order at the time of this function call. This will
	// allow me to "freeze" the order that the list is in at this
	// time so I can sequentially call every single item within
	// the list.

	lSlice := make([]*LinkedList, 0)
	current := l
	for current.next != l {
		lSlice = append(lSlice, current)
		current = current.next
	}

	// finally get the last one
	lSlice = append(lSlice, current)
	return lSlice
}
