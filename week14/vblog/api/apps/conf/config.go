package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)


func NewDefaultConfig() *Config{
	return &Config{
		App: newDefautApp(),
		Mysql: newDefaultMysql(),
	}
}


//程序所有配置信息都保存在该对象上
//Struct Tag: toml,等下 "github.com/BurntSushi/toml" 解析式，完成配置文件  对象的映射
//Struct Tag: env,等下 "github.com/caarlos0/env/v6"  解析式，完成环境变量  对象的映射
type Config struct {
	App *app `toml:"app"`
	Mysql *mysql `toml:"mysql"`
}

func (c *Config) String() string{
	jd,err:=json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(jd)
}

func newDefautApp() *app{
	return &app{
		Name:"vblog",
		HTTP: newDefaultHttp(),
	}
}

type app struct {
	//应用名称
	Name string `toml:"name" env:"APP_NAME"`
	HTTP *http `toml:"http"`
}

func newDefaultHttp() *http{
	return &http{
		Host: "localhost",
		Port: "7070",
	}
}

type http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port string `toml:"port" env:"HTTP_PORT"`
}

func newDefaultMysql() *mysql{
	return &mysql{
		Host: "localhost",
		Port: "3306",
		Database: "vblog",
		Username: "vblog",
		Password: "123456",
	}
}

type mysql struct {
	Host string `toml:"host" env:"MYSQL_HOST"`
	Port string `toml:"port" env:"MYSQL_PORT"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	Username string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`

	//连接池设置
	//最大连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	//最大的最大闲置连接数
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 连接的有效时间
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	//一个限制的连接多久没使用会被释放
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	lock sync.Mutex
	dbconn *sql.DB
}


//数据连接，需要单列模式
func (m *mysql) GetDB() (*sql.DB){
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.dbconn == nil{
		conn,err:=m.getDB()
		if err != nil {
			panic(err)
		}
		m.dbconn = conn
	}
	return m.dbconn

}

// 通过MySQL配置获取一个连接池
func (m *mysql) getDB() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	// 设置连接池参数
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	//通过Ping来测试当前MySQL服务是否可达
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}