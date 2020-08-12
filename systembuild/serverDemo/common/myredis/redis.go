package myredis

import (
	"fmt"
	"serverDemo/common/basic"
	"serverDemo/common/log"
	"serverDemo/common/util"
	"time"

	"github.com/go-redis/redis"
)

const (
	INFINITE time.Duration = 0
	SECOND   time.Duration = time.Second
	MINUTE   time.Duration = time.Minute
	HOUR     time.Duration = time.Hour
	DAY      time.Duration = 24 * HOUR
	WEEK     time.Duration = 7 * DAY
)

var redisdb *redis.Client

func InitRedis() {
	log.Info("redis.init:", util.AssertMarshal(basic.Redis))
	addr := fmt.Sprintf("%v:%v", basic.Redis.Host, basic.Redis.Port)
	redisdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: basic.Redis.Password,
		DB:       0,
	})
	_ = redisdb.Ping().String()
	//log.Info("redis pingResult:", pingResult)
}

func GetDb() *redis.Client {
	return redisdb
}

// 匹配获取key
func GetKeyPattern(pattern string) (ret []string, err error) {
	if len(pattern) < 1 {
		err = fmt.Errorf("pattern err")
		return
	}
	return redisdb.Keys(pattern).Result()
}

// 匹配删除key
func DelKeyPattern(pattern string) (int64, error) {
	keySlice, _ := GetKeyPattern(pattern)
	if len(keySlice) < 1 {
		return 0, nil
	}
	return redisdb.Del(keySlice...).Result()
}

func GetLock(key string, value interface{}, expiration time.Duration) bool {
	if len(key) < 1 {
		return false
	}
	ok, _ := redisdb.SetNX(key, value, expiration).Result()
	return ok
}

func GetLockKey(key string) string {
	return fmt.Sprintf("lock:%v", key)
}
