package cirDoubleList

import "fmt"

type doubleNode struct {
	data	interface{}
	prev,next	*doubleNode
}

func (dn *doubleNode) String() string {
	return fmt.Sprintf("%v",dn.data)
}