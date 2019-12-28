package util

type Comparable interface {
	IsEqual(a,b interface{}) bool
	//比较接口用于比较特定的两个元素的值是否相等，比如说结构体的某个键相等就返回true。
	CompareTo(a,b interface{}) int
	/* 用于比较两个类型的大小，a大于b为1，a等于b为0，a小于b为-1. */
}
