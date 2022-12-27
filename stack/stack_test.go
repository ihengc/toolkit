package stack

import "testing"

/*
* @author: Chen Chiheng
* @date: 2022/12/27 0027 10:23
* @version: 1.0
* @description:
**/

func TestStack(t *testing.T) {
	stack := New()
	testStackIsEmpty(stack, true, t)
	testStackLen(stack, 0, t)
	num := 10
	for i := 0; i < num; i++ {
		stack.Push(i)
		testStackLen(stack, i+1, t)
		testStackIsEmpty(stack, false, t)
	}
	for i := num - 1; i >= 0; i-- {
		val := stack.Pop()
		testStackPop(val.(int), i, t)
		testStackLen(stack, i, t)
		if i == 0 {
			testStackIsEmpty(stack, true, t)
		} else {
			testStackIsEmpty(stack, false, t)
		}
	}
}

func testStackLen(stack *Stack, dest int, t *testing.T) {
	if stack.Len() != dest {
		t.Fatal("Stack Len Error")
	}
}

func testStackPop(ret int, dest int, t *testing.T) {
	if ret != dest {
		t.Fatal("Stack Pop Error")
	}
}

func testStackIsEmpty(stack *Stack, ret bool, t *testing.T) {
	if stack.IsEmpty() != ret {
		t.Fatal("Stack IsEmpty Error")
	}
}
