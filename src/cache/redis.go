// redis.go
// author:昌维 [github.com/cw1997]
// date:2017-05-08 00:45:14
package cache

import (
	"errors"
	//	"fmt"
	"log"
	"strconv"
	//	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	"config"
)

var (
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

func ConnectPool() error {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = config.Get("cache.ip") + ":" + config.Get("cache.port")
	REDIS_DB, dbConfigErr := strconv.Atoi(config.Get("cache.db"))
	// 建立连接池
	maxIdle, maxIdleConfigErr := strconv.Atoi(config.Get("cache.maxIdle"))
	maxActive, maxActiveConfigErr := strconv.Atoi(config.Get("cache.maxActive"))
	idleTimeout, idleTimeoutConfigErr := strconv.Atoi(config.Get("cache.idleTimeout"))
	if dbConfigErr != nil || maxIdleConfigErr != nil || maxActiveConfigErr != nil || idleTimeoutConfigErr != nil {
		return errors.New("config error")
	}
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			if password := config.Get("cache.password"); password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	_, err := RedisClient.Dial()
	return err
}

type CacheConfig struct {
	ip       string
	port     string
	username string
	password string
}

// 废弃，改用redis连接池
/*func connect() (redis.Conn, error) {
	cacheConfig := &CacheConfig{
		config.Get("cache.ip"),
		config.Get("cache.port"),
		config.Get("cache.username"),
		config.Get("cache.password"),
	}
	uri := cacheConfig.ip + ":" + cacheConfig.port
	c, err := redis.Dial("tcp", uri)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, nil
	//	defer c.Close()
}*/
func connect() redis.Conn {
	// 从池里获取连接
	return RedisClient.Get()
}

func Set(key string, value string, expire int) (bool, error) {
	c := connect()
	// 用完后将连接放回连接池
	defer c.Close()
	_, errSet := c.Do("SET", key, value)
	if errSet != nil {
		log.Printf("set cache key-value error: %v\n", errSet)
		return false, errSet
	}
	_, errExpire := c.Do("EXPIRE", key, expire)
	if errExpire != nil {
		log.Printf("set cache expire error: %v\n", errExpire)
		return false, errExpire
	}
	//	fmt.Println(v)
	return true, nil
}

func Get(key string) (string, error) {
	c := connect()
	// 用完后将连接放回连接池
	defer c.Close()
	v, err := redis.String(c.Do("GET", key))
	return v, err
}

func GetTTL(key string) (interface{}, error) {
	c := connect()
	// 用完后将连接放回连接池
	defer c.Close()
	v, err := c.Do("TTL", key)
	if err != nil {
		log.Printf("get the redis' ket TTL error: %v\n", err)
	}
	return v, nil
}

func SetShortUrlCache(shortUrl string, longUrl string) (bool, error) {
	expire, errConv := strconv.Atoi(config.Get("cache.expire"))
	if errConv != nil {
		return true, errConv
	}
	key := config.Get("cache.prefix") + shortUrl
	if ok, errCache := Set(key, longUrl, expire); ok == true {
		return true, nil
	} else {
		return false, errCache
	}
}
