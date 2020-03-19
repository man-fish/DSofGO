package main

import (
	"fmt"
	"go_dataStruct/QUEUE"
	"go_dataStruct/STACK"
)

func generateParenthesis(n int) []string {
	answers := make([]string, 0)
	helper(n, n, "", &answers)
	return answers
}

func helperOne(l, r int, str string, store *[]string) bool {
	if l == 0 && r == 0 {
		*store = append(*store, str)

	}
	if l > 0 && helperOne(l-1, r, str+"(", store) {
		return true
	} else if r > 0 && r > l && helperOne(l, r-1, str+")", store) {
		return true
	} else {
		return false
	}
}

func helper(l, r int, str string, store *[]string) {
	if l == 0 && r == 0 {
		*store = append(*store, str)
	}

	if l > r {
		return
	}

	if l > 0 {
		helper(l-1, r, str+"(", store)
	}

	if r > 0 {
		helper(l, r-1, str+")", store)
	}
}

type node struct {
	left  int
	right int
	str   string
}

func generateParenthesisBFS(n int) []string {
	var (
		queue *QUEUE.LinkedQueue
		ans   []string
	)

	ans = make([]string, 0)
	queue = QUEUE.NewLinkedQueue()
	queue.Add(&node{
		left:  n,
		right: n,
		str:   "",
	})
	for !queue.IsEmpty() {
		size := queue.Size()
		for size > 0 {
			size--
			cur := queue.Poll().(*node)

			if cur.left == 0 && cur.right == 0 {
				ans = append(ans, cur.str)
				continue
			}
			if cur.left > 0 {
				queue.Add(&node{
					left:  cur.left - 1,
					right: cur.right,
					str:   cur.str + "(",
				})
			}
			if cur.right > cur.left {
				queue.Add(&node{
					left:  cur.left,
					right: cur.right - 1,
					str:   cur.str + ")",
				})
			}
		}
	}
	return ans
}

func generateParenthesisDFS(n int) []string {
	stk := STACK.NewSingleStack()
	stk.Push(&node{
		left:  n,
		right: n,
		str:   "",
	})

	ans := make([]string, 0)
	for !stk.IsEmpty() {
		cur := stk.Pop().(*node)
		if cur.left == 0 && cur.right == 0 {
			ans = append(ans, cur.str)
			continue
		}

		if cur.left > 0 {
			stk.Push(&node{
				left:  cur.left - 1,
				right: cur.right,
				str:   cur.str + "(",
			})
		}

		if cur.right > cur.left {
			stk.Push(&node{
				left:  cur.left,
				right: cur.right - 1,
				str:   cur.str + ")",
			})
		}
	}
	return ans
}

func main() {
	anses := generateParenthesis(2)
	fmt.Println(anses)
	anses = generateParenthesisBFS(2)
	fmt.Println(anses)
	anses = generateParenthesisDFS(2)
	fmt.Println(anses)
}
