package cache

import (
	"errors"
	"log"

	"github.com/emirpasic/gods/sets/treeset"
)

var (
	hashTableContainsError = errors.New("Cache: hashTable contains value but queues doesn't it")
)

type DataCell struct {
	Data interface{}
	Time int64
}

type KeyCell struct {
	Key  interface{}
	Time int64
}

type TwoQCache struct {
	logger *log.Logger

	hashTable map[interface{}]DataCell
	in        *treeset.Set
	inCap     int
	out       *treeset.Set
	outCap    int
	hot       *treeset.Set
	hotCap    int
}

func (c *TwoQCache) compare(a, b interface{}) int {
	acell, ok := a.(KeyCell)
	if !ok {
		c.logger.Printf("Cache: got data of type %T but wanted Cell", a)
		return 0
	}

	bcell, ok := b.(KeyCell)
	if !ok {
		c.logger.Printf("Cache: got data of type %T but wanted Cell", b)
		return 0
	}

	if acell.Time < bcell.Time {
		return -1
	} else if acell.Time == bcell.Time {
		return 0
	}
	return 1
}

func (c *TwoQCache) Get(key interface{}, time int64) (interface{}, bool) {
	if cell, ok := c.hashTable[key]; ok {
		oldTime := cell.Time
		if c.in.Contains(KeyCell{key, oldTime}) {
			c.in.Remove(KeyCell{key, oldTime})
			c.in.Add(KeyCell{key, time})
		} else if c.out.Contains(KeyCell{key, oldTime}) {
			c.out.Remove(KeyCell{key, oldTime})

			if c.hot.Size() == c.hotCap {
				it := c.hot.Iterator()
				exist := it.First()
				if !exist {
					return nil, false
				}

				c.hot.Remove(it.Value())
			}
			c.hot.Add(KeyCell{key, time})
		} else if c.hot.Contains(KeyCell{key, oldTime}) {
			c.hot.Remove(KeyCell{key, oldTime})
			c.hot.Add(KeyCell{key, time})
		} else {
			c.logger.Print(hashTableContainsError)
			return nil, false
		}
		c.hashTable[key] = DataCell{cell.Data, time}
		return cell.Data, true
	}
	return nil, false
}

func (c *TwoQCache) Delete(key interface{}) {
	if cell, ok := c.hashTable[key]; ok {
		keyCell := KeyCell{key, cell.Time}
		if c.in.Contains(keyCell) {
			c.in.Remove(keyCell)
		} else if c.out.Contains(keyCell) {
			c.out.Remove(keyCell)
		} else if c.hot.Contains(keyCell) {
			c.hot.Remove(keyCell)
		} else {
			c.logger.Print(hashTableContainsError)
		}
		delete(c.hashTable, key)
	}
}

func (c *TwoQCache) Set(key, data interface{}) {
	if cell, ok := c.hashTable[key]; ok {
		c.hashTable[key] = DataCell{data, cell.Time}
	}
}

func NewTwoQCache(inCap, outCap, hotCap int, logger *log.Logger) *TwoQCache {
	cache := &TwoQCache{
		logger:    logger,
		hashTable: make(map[interface{}]DataCell),
		inCap:     inCap,
		outCap:    outCap,
		hotCap:    hotCap,
	}
	cache.in = treeset.NewWith(cache.compare)
	cache.out = treeset.NewWith(cache.compare)
	cache.hot = treeset.NewWith(cache.compare)

	return cache
}
