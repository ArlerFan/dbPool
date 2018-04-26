package dbpool

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

//Config -資料庫相關組態
type MongoConfig struct {
	ConnName string //連線名稱，若沒填寫則為 default
	Host     string
	UserName string //資料庫帳號
	Password string //資料庫密碼
	DBName   string //資料庫名稱
}

//default value
func NewMongoConfig() *MongoConfig {
	return &MongoConfig{

	}
}

func (config *MongoConfig) MongoConn() (*mgo.Session, error) {
	var connStr string
	//未設定帳號密碼
	if config.UserName == "" || config.Password == "" {
		connStr = fmt.Sprintf("%v", config.Host)
	} else {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s/%s", config.UserName, config.Password, config.Host, config.DBName)
	}
	m, err := mgo.Dial(connStr)
	if err != nil {
		return nil, err
	}
	return m, err
}
