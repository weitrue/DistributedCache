/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午11:09
 * Description: 实现Http服务
 **/

package http

import (
	"DistributedCache/cache"
	"log"
	"net/http"
)

type Server struct {
	cache.Cache  // 内嵌Cache接口，意味着Server也实现了Cache接口
}

/**
   监听方法
 */
func (s *Server) Listen()  {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status/", s.statusHandler())
	http.Handle("/status", s.statusHandler())
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/**
   初始化方法
 */
func New(c cache.Cache) *Server {
	return &Server{c}
}
