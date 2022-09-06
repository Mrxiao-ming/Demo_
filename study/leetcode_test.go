package study

import (
	"fmt"
	"testing"
)

func TestFun21(t *testing.T) {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}
	res := Fun21(&list1, &list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestFun206(t *testing.T) {
	res := Fun206(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	})
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestFun1480(t *testing.T) {
	fmt.Println(Fun1480([]int{1, 2, 3, 4, 5}))
}

func TestFun205(t *testing.T) {
	fmt.Println(Fun205("paper", "title"))
}

func TestFun392(t *testing.T) {
	fmt.Println(Fun392("", "ahbgdc"))
}

func TestFun142_2(t *testing.T) {
	nodeCycle2 := &ListNode{
		Val: 6,
	}
	nodeCycle1 := &ListNode{
		Val: 100,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nodeCycle2,
					},
				},
			},
		},
	}
	nodeCycle2.Next = nodeCycle1

	node := &ListNode{
		Val:  1,
		Next: nodeCycle1}

	res := Fun142_2(node)
	fmt.Println(res)

}

func TestFun733(t *testing.T) {
	fmt.Println(Fun733([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}

func TestFun509(t *testing.T) {
	fmt.Println(Fun509(7))
}

func TestFun135(t *testing.T) {
	fmt.Println(Fun135([]int{1, 0, 2}))
}
