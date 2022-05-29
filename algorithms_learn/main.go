package main

import "./lru"

type String string

func (d String) Len() int {
	return len(d)
}

func main() {
	LRU := lru.New(int64(0), nil)
	LRU.Add("key1", String("1234"))
	value1, ok1 := LRU.Get("key1")
	if ok1 == true {
		print(value1.(String), ok1)
	}
	value2, ok2 := LRU.Get("key2")
	if ok2 == true {
		print(value2.(String), ok2)
	}
}
