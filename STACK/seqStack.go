package STACK

type SeqStack struct {
	element []interface{}
	maxSize int
	top     int
}

func NewSeqStack(length int) *SeqStack {
	return &SeqStack{
		element: make([]interface{}, length),
		maxSize: length,
		top:     0,
	}
}

func (stk *SeqStack) IsEmpty() bool {
	return stk.top == 0
}

func (stk *SeqStack) Peek() interface{} {
	if stk.top == 0 {
		return nil
	}
	return stk.element[stk.top-1]
}

func (stk *SeqStack) Pop() interface{} {
	if stk.top == 0 {
		panic("栈内空空如野")
	}

	defer func() {
		stk.top--
	}()

	return stk.element[stk.top-1]
}

func (stk *SeqStack) Push(x interface{}) {
	if stk.top == stk.maxSize {
		panic("满员了，栈溢出")
	}
	stk.element[stk.top] = x
	stk.top++
}
