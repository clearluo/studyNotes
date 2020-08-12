package myredis

import (
	"errors"
	"fmt"
	"serverDemo/common/log"
	"time"
)

func Set(key string, value interface{}, timeout time.Duration) error {
	if len(key) < 1 || timeout < 0 || value == nil {
		err := fmt.Errorf("redis.Set args err:%v,%v,%v", key, value, timeout)
		return err
	}
	err := redisdb.Set(key, value, timeout).Err()
	if err != nil {
		log.Warn(err)
		return fmt.Errorf("redis.Set err:%v", err)
	}
	return nil
}

func GetString(key string) (string, error) {
	if len(key) < 1 {
		err := errors.New("redis.Get args err")
		return "", err
	}
	return redisdb.Get(key).Result()
}

func GetInt(key string) (int, error) {
	if len(key) < 1 {
		err := errors.New("redis.Get args err")
		return 0, err
	}
	return redisdb.Get(key).Int()
}
