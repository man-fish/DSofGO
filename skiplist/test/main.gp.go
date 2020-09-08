package main

import "go_dataStruct/skiplist"

type comparabelInt int

func (comparabelInt) CompareTo(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func (comparabelInt) IsEqual(a, b interface{}) bool {
	return a.(int) == b.(int)
}

func main() {
	var key1 comparabelInt = 1
	var key2 comparabelInt = 2
	var key3 comparabelInt = 3
	var key4 comparabelInt = 4

	list := skiplist.New()
	list.Put(key1, 1)
	list.Put(key2, 1)
	list.Put(key3, 1)
	list.Put(key4, 1)

}
