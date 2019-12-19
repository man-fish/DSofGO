package LIST

import "fmt"

type DoubleNode struct {
	Data	interface{}
	Prev,Next	*DoubleNode
}

func (dn *DoubleNode) String() string {
	return fmt.Sprintf("%v",dn.Data)
}