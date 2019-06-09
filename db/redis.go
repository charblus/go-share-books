package db

import "github.com/go-redis/redis"

var RDS redis.UniversalClient


func newRedisClient(addr string, pwd string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

//IntiRedisClient 初始化redis
func IntiRedisClient(conf *RedisConfig) error {
	client, err := newRedisClient(conf.Addr, conf.Password, conf.DB)

	if err != nil {
		return err
	}
	RDS = client
	return nil
}
