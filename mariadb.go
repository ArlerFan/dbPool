package dbpool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

//Config -資料庫相關組態
type MariaConfig struct {
	ConnName string //連線名稱，若沒填寫則為 default
	Host     string
	UserName string //資料庫帳號
	Password string //資料庫密碼
	DBName   string //資料庫名稱
}

//default value
func NewMariaConfig() *MariaConfig {
	return &MariaConfig{

	}
}

func (config *MariaConfig) MariaConn() (db *gorm.DB, err error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s", config.UserName, config.Password, config.Host, config.DBName)
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(50)
	db.DB().SetMaxIdleConns(25)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	return db, err
}
