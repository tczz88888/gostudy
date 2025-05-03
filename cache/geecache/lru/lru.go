package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element

	OnEviced func(key string, value Value) //某条记录被移除时的回调
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache { //初始化cache
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		OnEviced: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) //强转成entry类型的指针
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() { //删除最久未使用的元素
	ele := c.ll.Back() //front是最新的，back是最旧的
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEviced != nil {
			c.OnEviced(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { //已经存在则修改
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) //kv是指向list中一个ele的指针，所以修改kv指向的元素，list中的ele的内容也会被修改
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.nbytes += int64(len(key)) + int64(value.Len())
		c.cache[key] = ele
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
