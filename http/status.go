/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午11:48
 * Description:
 **/

package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusHandler struct {
	*Server
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 序列化查询到的缓存状态数据
	//{"Count":0,"KeySize":0,"ValueSize":0}
	b, e := json.Marshal(h.GetStat())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *Server) statusHandler() http.Handler {
	return &StatusHandler{s}
}
