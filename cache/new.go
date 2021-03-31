/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午10:11
 * Description: 适配器
 **/

package cache

import (
	"DistributedCache/cache/memory"
	"log"
)

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		// 内存方式实现的缓存
		c = memory.NewInMemoryCache()
	}
	if typ == "rocksdb" {
		// 基于 rocksdb 存储

	}
	if c == nil {
		panic("Unknown Cache Type " + typ)
	}
	log.Println(typ, "ready to server")
	return c
}
