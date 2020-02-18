package search

import (
	"fmt"
	"hash/crc32"
)

// HashTable 这个散列表其实有缺陷，首先除留取🐟法除以的最好是最大素数，其次interface类型的比较有可能涉及到不可比较的类型map，slice
type HashTable struct {
	table []interface{}
	count  int
}

func NewHashTable(values []interface{}) *HashTable {
	hashTable := &HashTable{
		count: len(values),
	}
	hashTable.insert(values)
	return hashTable
}

func (ht *HashTable) Search(x interface{}) interface{} {
	hashKey := ht.hash(x)
	for ht.table[hashKey] != x {
		hashKey = (hashKey + 1)%ht.count
		hashKey++
	}
	return ht.table[hashKey]
}

func (ht *HashTable) Removee(x interface{}) interface{} {
	hashKey := ht.hash(x)
	for ht.table[hashKey] != x {
		hashKey = (hashKey + 1) % ht.count
	}
	v := ht.table[hashKey]
	ht.table[hashKey] = nil
	return v
}

func (ht *HashTable) Insert(x interface{}) {
	hashKey := ht.hash(x)
	count := 0
	for ht.table[hashKey] != nil {
		hashKey = (hashKey+1)%ht.count
		count++
	}
	count++
	ht.table[hashKey] = x
	fmt.Printf("尝试插入次数：%d", count)
}

func (ht *HashTable) insert(xes []interface{}) {
	for _, x := range xes {
		ht.Insert(x)
	}
}

func (ht *HashTable) hash(x interface{}) int {
	hashkey := int(crc32.ChecksumIEEE([]byte(fmt.Sprint(x))))
	key := hashkey % len(ht.table)
	return key
}


