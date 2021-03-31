/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午9:41
 * Description:
 **/

package main

import (
	"DistributedCache/cache"
	"DistributedCache/http"
	"DistributedCache/tcp"
)

func main() {
	c := cache.New("inmemory")

	go tcp.New(c).Listen()

	http.New(c).Listen()
}
