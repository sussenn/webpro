package tool

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

//redis 连接
type RedisStore struct {
	client *redis.Client
}

var RedisIn RedisStore

//redis初始化
func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	RedisIn = RedisStore{
		client: client,
	}
	//TODO 将图形验证码存入缓存...
	return &RedisIn
}

//插入redis
func (rs *RedisStore) Set(id string, value string) {
	err := rs.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

//获取
func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		//取出缓存并删除
		err := rs.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
