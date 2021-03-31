/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午11:18
 * Description:
 **/

package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CacheHandler struct {
	* Server
}

/**
 * Go标准包中net/http有Handler的定义
 * type Handler interface{
 *     ServeHTTP(ResponseWriter, *Request)
 * }
 * 这里复写标准包中的ServeHTTP
 * 通过HTTP协议中的PUT GET DELETE实现对应缓存的Set Get Del
 */
func (h *CacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	method := r.Method
	if method == http.MethodPut {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if len(b) != 0 {
			err = h.Set(key, b)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}
	if method == http.MethodGet {
		b, err := h.Get(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write(b)
		return
	}
	if method == http.MethodDelete {
		err := h.Del(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) cacheHandler() http.Handler {
	return &CacheHandler{s}
}
