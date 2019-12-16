package singleList

import (
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%v",n.data)
}