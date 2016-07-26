package utils

import (
	beego_utils "github.com/astaxie/beego/utils"
)

const (
	DEF_CACHE_ITEM_EXPIRE_TIME = 300  //300 秒， 5分钟
	DEF_CACHE_CLEAR_INTERVAL   = 3600 //3600 秒， 1小时
)

type CacheItem struct {
	time_set_sec        int64 //该数据被设置的时间
	time_expire_len_sec int64 //数据过期时间
	data                interface{}
}

type CacheManager struct {
	cache_map *beego_utils.BeeMap // string -> CacheItem
}

var g_cache_manager CacheManager
var g_last_clear_cache_time int64

func GetCacheManager() *CacheManager {
	if g_cache_manager.cache_map == nil {
		g_cache_manager.cache_map = beego_utils.NewBeeMap()
		g_last_clear_cache_time = GetNowTimeInSec()
	}

	return &g_cache_manager
}

// Get from maps return the k's value
func (m *CacheManager) Get(k string) interface{} {
	val, _ := m.cache_map.Get(k).(*CacheItem)
	if val == nil {
		return val
	}

	time_now := GetNowTimeInSec()
	if (time_now - val.time_set_sec) >= val.time_expire_len_sec {
		m.cache_map.Delete(k)
		return nil
	}

	if time_now-g_last_clear_cache_time > DEF_CACHE_CLEAR_INTERVAL {
		m.clearExpireItems()
	}

	return nil
}

func (m *CacheManager) Set(k string, data interface{}) bool {
	return m.SetWithExpireTime(k, data, DEF_CACHE_ITEM_EXPIRE_TIME)
}

func (m *CacheManager) SetWithExpireTime(k string, data interface{}, expire_interval int64) bool {
	cache_item := CacheItem{time_set_sec: GetNowTimeInSec(),
		time_expire_len_sec: expire_interval,
		data:                data}

	return m.cache_map.Set(k, &cache_item)
}

func (m *CacheManager) Delete(k string) {
	m.cache_map.Delete(k)
}

func (m *CacheManager) clearExpireItems() {
	//简单粗暴, 将来要改
	m.cache_map = beego_utils.NewBeeMap()
	g_last_clear_cache_time = GetNowTimeInSec()
}
