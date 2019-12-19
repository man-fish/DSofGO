package LIST

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%v",n.Data)
}