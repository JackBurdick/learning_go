package addTwoNumbers

import (
	"fmt"
	"testing"
)

func createList(vals []int) (*listNode, error) {
	lH := &listNode{}
	pN := lH
	cN := lH
	for _, val := range vals {
		nN := &listNode{}
		cN.val = val
		cN.nPtr = nN
		pN = cN
		cN = nN
	}
	pN.nPtr = nil
	return lH, nil
}

func compareTwoLinkLists(want *listNode, got *listNode) (bool, error) {
	wN := want
	gN := got
	for wN != nil && gN != nil {
		same := wN.val == gN.val
		switch same {
		case false:
			return false, nil
		default:
			wN = wN.nPtr
			gN = gN.nPtr
		}
	}
	if wN != nil || gN != nil {
		return false, fmt.Errorf("One list is longer than the other")
	}
	return true, nil
}

func Test_twoSum(t *testing.T) {
	var cases = []struct {
		target []int
		l1     []int
		l2     []int
	}{
		{
			// no carry, equal len, same sum len
			[]int{2, 4, 6},
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			// carry, equal len, same sum len
			[]int{0, 5, 6},
			[]int{9, 2, 3},
			[]int{1, 2, 3},
		},
		{
			// no carry, non-equal len, same sum len
			[]int{2, 3, 4, 1, 1},
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3},
		},
		{
			// carry, non-equal len, same sum len
			[]int{0, 4, 4, 1, 1},
			[]int{1, 1, 1, 1, 1},
			[]int{9, 2, 3},
		},
		{
			// carry, non-equal len, > sum len
			[]int{2, 2, 0, 1},
			[]int{1, 1, 9},
			[]int{1, 1, 1},
		},
		{
			// overflow carry
			[]int{0, 1, 2, 1},
			[]int{9, 9, 9},
			[]int{1, 1, 2},
		},
		{
			// l1 empty
			[]int{0, 1, 2},
			[]int{},
			[]int{0, 1, 2},
		},
		{
			// l2 empty
			[]int{3, 4, 5},
			[]int{3, 4, 5},
			[]int{},
		},
	}
	for _, c := range cases {
		l1, err := createList(c.l1)
		if err != nil {
			t.Errorf("Error creating l1: %v", err)
		}
		l2, err := createList(c.l2)
		if err != nil {
			t.Errorf("Error creating l2: %v", err)
		}
		want, err := createList(c.target)
		if err != nil {
			t.Errorf("Error creating want: %v", err)
		}

		got, err := addTwoNumbers(l1, l2)
		if err != nil {
			t.Errorf("Error adding two lists: %v\n l1: %v\n, l2 %v", err, c.l1, c.l2)
		}
		ok, err := compareTwoLinkLists(got, want)
		if err != nil {
			t.Errorf("Error when comparing two lists:%v\n l1: %v\n, l2 %v", err, c.l1, c.l2)
		}
		if !ok {
			t.Errorf(`Sum and target value do not match:\n 
					  > l1  : %v\n, 
					  > l2  : %v\n,
					  > want: %v\n,
					  > sum : %v`, l1, l2, want, got)
		}
	}
}
