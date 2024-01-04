package utils

import "fmt"

// Containts functions to handle Linked Lists and Circular Linked Lists
// By default lists are just linked, to make them circular the list needs
// to be passed to ConvertSinglyToCircular

type Node struct {
	info interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (n *Node) GetInfo() interface{} { return n.info }

func (l *LinkedList) Insert(d interface{}) {
	node := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
}

func (l *LinkedList) Values() []interface{} {
	return_list := []interface{}{}
	p := l.head
	for p != nil {
		return_list = append(return_list, p.info)
		p = p.next
	}
	return return_list
}

func ShowCircular(l *LinkedList) {
	p := l.head
	for {
		if p.next == l.head {
			fmt.Printf("-> %v ", p.info)
			break
		}
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
	fmt.Println()
}

func ConvertSinglyToCircular(l *LinkedList) {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = l.head
}

func (l *LinkedList) slideForward(d interface{}, slide int) {
	if slide < 0 {
		panic("Please use slide backward")
	}
	p := l.head
	// Loop through the list until we find that the next value is what we want
	for p.next.info != d {
		p = p.next
	}

	// P is now the previous node
	// Grab the node, and cut it out of the loop
	node_to_slide := p.next
	p.next = p.next.next

	for i := 0; i < slide; i++ {
		p = p.next
	}
	// p is now set to the value before where we want to insert

	node_to_slide.next = p.next
	p.next = node_to_slide
}

func (l *LinkedList) slideBackward(d interface{}, slide int) {
	if slide > 0 {
		panic("Please use slide forward")
	}
	p := l.head
	// Loop through the list until we find that the next value is what we want
	for p.next.info != d {
		p = p.next
	}

	// P is now the previous node
	// Grab the node, and cut it out of the loop
	node_to_slide := p.next
	p.next = p.next.next

	for i := 0; i > slide; i-- {
		p = p.next
	}
	// p is now set to the value before where we want to insert

	node_to_slide.next = p.next
	p.next = node_to_slide
}

func (l *LinkedList) Slide(d interface{}, slide int, len int) {
	if slide > 0 {
		l.slideForward(d, slide)
	} else {
		l.slideForward(d, slide+len-1)
	}
}

func GetPos(l *LinkedList, pos int) interface{} {
	p := l.head

	for i := 0; i <= pos; i++ {
		p = p.next
	}
	return p.info
}

func GetFirst(l *LinkedList) *Node {
	return l.head
}

func GetNext(n *Node) *Node {
	return n.next
}

func GetPosFrom(l *LinkedList, from interface{}, pos int) interface{} {
	p := l.head

	for p.info != from {
		p = p.next
	}
	// now p is out start value
	for i := 0; i < pos; i++ {
		p = p.next
	}
	return p.info
}
