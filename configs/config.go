package config

import (
	"go-share-books/db"
	"encoding/json"
	"os"
	"github.com/panjiang/golog"
)

// Conf 全局配置
var Conf *Config
//Config 配置文件
type Config struct {
	SiteName string          `json:"site_name"`
	Release  bool            `json:"release"`
	BindPort int             `json:"bind_port"`
	Psql     *db.PsqlConfig  `json:"psql"`
	Redis    *db.RedisConfig `json:"redis"`
	Log      *log.Config     `json:"log"`
	Wx       *wx             `json:"wx"`
}
type wx struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func (c *Config) parse(filename string) error {
	
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		return err
	}

	log.Infof("log config (%s)", c.Log.DebugString())
	err = log.ParseConfig(c.Log)
	if err != nil {
		return err
	}

	return nil
}

// Parse 初始化配置实例
func Parse() {
	filename := "config.json"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	log.Printf("parse config: %s", filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal(err)
	}

	Conf = new(Config)
	err := Conf.parse(filename)
	if err != nil {
		log.Fatal(err)
	}
}
