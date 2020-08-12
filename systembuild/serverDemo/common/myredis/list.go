package myredis

import (
	"errors"
	"fmt"
	"serverDemo/common/log"
	"time"
)

func Rpush(key string, value ...interface{}) error {
	if len(key) < 1 || len(value) < 1 {
		err := fmt.Errorf("redis.Rpush args err:%v,%v,%v", key, value)
		return err
	}
	err := redisdb.RPush(key, value...).Err()
	if err != nil {
		log.Warn(err)
		return fmt.Errorf("redis.Rpush err:%v", err)
	}
	return nil
}

func Blpop(key string, timeout time.Duration) ([]string, error) {
	if len(key) < 1 {
		err := errors.New("redis.Blpop args err")
		return nil, err
	}
	return redisdb.BLPop(timeout, key).Result()
}

func Llen(key string) (int64, error) {
	return redisdb.LLen(key).Result()
}
