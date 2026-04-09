package main

import (
	"container/list"
	"fmt"
)

type Data struct {
	key, val, cnt int
}

type LFUCache struct {
	capacity   int
	length     int
	freqRecord []*list.List
	cache      map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		length:     0,
		freqRecord: []*list.List{list.New(), list.New()},
		cache:      make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	if e, exist := this.cache[key]; exist {
		data := e.Value.(Data)

		// remove from old list
		l := this.freqRecord[data.cnt]
		l.Remove(e)

		data.cnt++

		if len(this.freqRecord) == data.cnt {
			this.freqRecord = append(this.freqRecord, list.New())
		}

		next_l := this.freqRecord[data.cnt]
		e = next_l.PushBack(data)

		this.cache[data.key] = e
		return data.val
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if e, exist := this.cache[key]; exist {
		data := e.Value.(Data)

		// remove from old list
		l := this.freqRecord[data.cnt]
		l.Remove(e)

		data.val = value
		data.cnt++

		if len(this.freqRecord) == data.cnt {
			this.freqRecord = append(this.freqRecord, list.New())
		}

		next_l := this.freqRecord[data.cnt]
		e = next_l.PushBack(data)

		this.cache[data.key] = e

		return
	}

	if this.length == this.capacity {
		// Capacity not sufficient, evict the least recently used key

		// find the least frequency used element
		i := 1
		for ; i < len(this.freqRecord); i++ {
			if this.freqRecord[i].Len() != 0 {
				break
			}
		}

		l := this.freqRecord[i]
		e := l.Front()
		delete(this.cache, e.Value.(Data).key)
		this.freqRecord[i].Remove(e)

		e = this.freqRecord[1].PushBack(Data{key, value, 1})
		this.cache[key] = e
	} else {
		// Capacity Sufficient
		l := this.freqRecord[1]
		e := l.PushBack(Data{key, value, 1})

		this.cache[key] = e
		this.length++
	}
}

func main() {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
