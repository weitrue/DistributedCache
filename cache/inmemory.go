/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午10:44
 * Description: 缓存对象（内存方式）
 **/

package cache

import "sync"

type InMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

/**
   实现接口方法，加入缓存（内存）
 */
func (c *InMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

/**
   实现接口方法，获取数据
 */
func (c *InMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.c[k], nil
}

/**
   实现接口方法，从缓存中删除（内存）
 */
func (c *InMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

/**
   获取缓存当前状态
 */
func (c *InMemoryCache) GetStat() Stat {
	return c.Stat
}

/**
   初始化缓存对象（内存方式）
 */
func newInMemoryCache() *InMemoryCache {
	//return &InMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
	return &InMemoryCache{
		c: make(map[string][]byte),
		mutex: sync.RWMutex{},
		Stat:  Stat{},
	}
}