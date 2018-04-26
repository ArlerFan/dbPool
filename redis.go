package dbpool

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"strconv"
)

//Config -資料庫相關組態
type RedisConfig struct {
	ConnName string //連線名稱，若沒填寫則為 default
	Host     string
	UserName string //資料庫帳號
	Password string //資料庫密碼
	DBName   string //資料庫名稱
}

//default value
func NewRedisConfig() *RedisConfig {
	return &RedisConfig{

	}
}

func (config *RedisConfig) RedisConn() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 100,
		//MaxActive:   12000, // max number of connections
		MaxActive:   0, // max number of connections
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			var option []redis.DialOption

			if config.Password != "" {
				option = []redis.DialOption{redis.DialPassword(config.Password)}
			}
			r, err := redis.Dial("tcp", config.Host, option...)
			if err != nil {
				return nil, err
			}
			if config.DBName != "" {
				dbNum, err := strconv.ParseInt(config.DBName, 10, 32)
				if err != nil {
					return nil, err
				}
				_, err = r.Do("SELECT", dbNum)

				if err != nil {
					r.Close()
					return nil, err
				}
			}
			return r, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) (err error) {
			_, err = c.Do("PING")
			return err
		},
	}
}

