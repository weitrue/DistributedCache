/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午9:50
 * Description: 缓存状态对象
 **/

package cache

type Stat struct {
	// 缓存中保存键值对的数量
	Count      int64
	// key占据的字节数
	KeySize    int64
	// value占据的字节数
	ValueSize  int64
}

/**
   新增键值对时更新缓存状态
 */
func (s *Stat) add(k string, v []byte)  {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

/**
   删除键值对时更新缓存状态
 */
func (s *Stat) del(k string, v []byte)  {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
