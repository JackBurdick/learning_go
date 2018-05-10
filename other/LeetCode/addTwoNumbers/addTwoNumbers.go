/*
from leetcode:
	"You are given two non-empty linked lists representing two non-negative
	integers. The digits are stored in reverse order and each of their nodes
	contain a single digit. Add the two numbers and return it as a linked list.

	Note: You may assume the two numbers do not contain any leading zero,
	except the number 0 itself."

Example solution:
	Given;
		- (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Solution;
		- 7 -> 0 -> 8

Rational;
	2+5 = `7`
	4+6 = 10 --> `0` w/carry
	(3+4)+carry = 7+1 --> `8`
	therefore, the answer = 7 -> 0 -> 8.

*/
package addTwoNumbers

import (
	"bytes"
	"fmt"
)

type listNode struct {
	val  int
	nPtr *listNode
}

// String implements the stringer.
func (l *listNode) String() string {

	// Example output will look similar to `3 -> 5 -> 4`.

	var buf bytes.Buffer
	for l.nPtr != nil {
		buf.WriteString(fmt.Sprintf("%v -> ", l.val))
		l = l.nPtr
	}
	if l.nPtr == nil {
		buf.WriteString(fmt.Sprintf("%v", l.val))
	}
	return buf.String()
}

func addTwoNumbers(l1 *listNode, l2 *listNode) (*listNode, error) {

	c := false

	// Store address of the head of the list to return after completion.
	headN := &listNode{}
	curN := headN
	pN := curN
	p := l1
	q := l2
	if p == nil && q == nil {
		return nil, fmt.Errorf("Niether list contained values - %v, %v", l1, l2)
	}
	sum := 0
	for p != nil || q != nil {
		curN.nPtr = &listNode{}

		// Get values from each list at the current "index" and advance to next
		// value in the list.
		sum = 0
		if p != nil {
			sum += p.val
			p = p.nPtr
		}
		if q != nil {
			sum += q.val
			q = q.nPtr
		}
		if c {
			sum++
		}

		// Determine if a carry will be present for the next iteration.
		c = false
		if sum >= 10 {
			sum -= 10
			c = true
		}

		// Create and advance to next node.
		curN.val = sum
		pN = curN
		curN = curN.nPtr
	}

	// Consider if final carry is present.
	switch c {
	case true:
		curN.val = 1
	default:
		pN.nPtr = nil
	}

	return headN, nil
}
