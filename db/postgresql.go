package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" 
)

//DB 全局DB
var DB *gorm.DB
//PsqlConfig  用于解析PostgreSQL
type PsqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db"`
}
// RedisConfig 用于解析redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func newPsqlClient(host string, port int, user string, dbName string, password string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.DB().SetConnMaxLifetime(time.Second * 20)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(30)
  
	return db, nil
}

//InitPsqlClient 初始化mysql 客户端
func InitPsqlClient(conf *PsqlConfig) error {
	client, err := newPsqlClient(conf.Host, conf.Port, conf.User, conf.DBName, conf.Password)
	if err != nil {

		return err
	}

	DB = client
	return nil
}

