/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/13 上午9:44
 * Description: 对外暴露缓存接口
 **/

package cache

/**
 * 定义缓存接口对外暴露的接口，实现与定义分离,从而达到支持多态的目的
 */
type Cache interface {
	// 设置键值对到缓存
	Set(string, []byte) error
	// 根据key从缓存中获取value
	Get(string) ([]byte, error)
	// 从缓存中删除key
	Del(string) error
	// 获取缓存状态
	GetStat() Stat
}