/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午10:11
 * Description: 适配器
 **/

package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		// 内存方式实现的缓存
        c = newInMemoryCache()
	}
	if c == nil {
		panic("Unknown Cache Type "+ typ)
	}
	log.Println(typ, "ready to server")
	return c
}