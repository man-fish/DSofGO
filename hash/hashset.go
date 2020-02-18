package hash

import (
	"errors"
	"reflect"
)

// HashSet datastruct hashset realized by HashMap
type HashSet struct {
	data map[interface{}]interface{}
	dataType string
	size int
}

// NewHashSet
func NewHashSet(dataType string) *HashSet {
	return &HashSet{
		data: make(map[interface{}]interface{}),
		dataType: string,
		size: 0
	}
}

// Size 获取散列大小
func (hs *HashSet) Size() int {
	return hs.size
}

// DataType 获取散列类型
func (hs *HashSet) DataType() string {
	return hs.dataType
}

// Members 获取散列所有元素
func (hs *HashSet) Members() []interface{} {
	members := make([]interface{}, 0)
	for key, _ := range hs.data {
		members = append(members, key)
	}
	return members
}

// Add 向散列中添加元素
func (hs *HashSet) Add(key interface{}) error {
	err := hs.checkData(key)
	if err != nil {
		return err
	}
	_, ok := hs.data[key]
	if ok {
		return errors.New("ElementExisted")
	}
	hs.data[key] = nil
	hs.size++
	return nil
} 
// Remove 从散列中移除元素
func (hs *HashSet) Remove(key interface{}) {
	err := hs.checkData(key)
	if err != nil {
		return err
	}
	dataKey, ok := hs.data[key]
	if ok {
		delete(hs.data, dataKey)
		hs.size--
		return nil
	}
	return errors.New("KeyNotExisted")
}

// Contain 散列中是否包含元素
func (hs *HashSet) Contain(key interface{}) (bool, error) {
	err := hs.checkData(key)
	if err != nil {
		return false, err
	}
	_, ok := hs.data[key]
	if ok {
		return true, nil
	} else {
		return false, nil
	}
}

// Clear 清空散列
func (hs *HashSet) Clear() {
	hs.data = make(map[interface{}]interface{})
	hs.size = 0
}

func (hs *HashSet) checkData(data interface{}) error {
	if data == nil {
		return errors.New("dataIsNil")
	}

	dataTypeof := reflect.TypeOf(data).String()
	if dataTypeof != hs.dataType {
		return errors.New("UnSupportedType")
	}
	return nil
}