package myredis

import "serverDemo/common/log"

func Publish(channel string, msg interface{}) error {
	_, err := redisdb.Publish(channel, msg).Result()
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
