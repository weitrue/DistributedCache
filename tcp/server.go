/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/21 下午9:47
 * Description: 基于HTTP/REST协议实现的缓存性能较差，只有Redis的四分之一，Redis使用的序列化协议规范是基于TCP实现的
                因此，实现基于TCP的缓存服务
 **/

package tcp

import (
	"DistributedCache/cache"
	"bufio"
	"io"
	"log"
	"net"
)


type Server struct {
	cache.Cache
}

func (s *Server) Listen()  {
	listen, err := net.Listen("tcp", ":9099")
	if err != nil {
		panic(err)
	}

	for  {
		c, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(c)
	}
}

/**
 * 处理请求
 */
func (s *Server) process(conn net.Conn)  {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("Close connection due to error:", err)
			}
			return
		}
		if op == 'S' {
			err = s.set(conn, r)
		} else if op == 'G' {
			err = s.get(conn, r)
		} else if op == 'D' {
			err = s.del(conn, r)
		} else {
			log.Println("Close connection due to invalid operation:", op)
			return
		}
		if err != nil {
			log.Println("Close connection due to error:", err)
			return
		}
	}
}

/**
 * 反序列化 传入的字节数组中的key
 */
func (s *Server) readKey(r *bufio.Reader) (string, error) {
	keyLen, err := readLen(r)
	if err != nil {
		return "", err
	}
	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

/**
 * 反序列化 传入的字节数组中的key和value
 */
func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	keyLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	valLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	key := make([]byte, keyLen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", nil, err
	}
	val := make([]byte, valLen)
	_, err = io.ReadFull(r, val)
	if err != nil {
		return "", nil, err
	}
	return string(key), val, nil
}

/**
 * 获取缓存中key的值
 */
func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	val, err := s.Get(key)
	return sendResponse(val, err, conn)
}

/**
 * 往缓存中设置值
 */
func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, val, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Set(key, val), conn)
}

/**
 * 从缓存中删除key
 */
func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Del(key), conn)
}

/**
 * 构造器
 */
func New(c cache.Cache) *Server {
	return &Server{c}
}