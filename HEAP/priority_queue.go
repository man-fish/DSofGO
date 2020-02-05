/*
Package heap physical datastruct
*/
package heap

//PriorityQueue DataStruct of Logic
type PriorityQueue interface {
	Insert(interface{})  //插入
	Max() interface{}    //返回最大值
	DelMax() interface{} //删除最大元素
	IsEmpty() bool       //是否为空
	Size() int           //实际容量
}
