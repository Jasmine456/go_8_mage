package impl

import (
	"database/sql"
	"vblog/api/apps/conf"
)

func NewImpl() *Impl{
	return &Impl{}
}


//依赖MySQL连接，能与MySQL交互
// 负责实现Blog service
type Impl struct {
	//通过配置文件
	db *sql.DB

}

//当这个对象初始化的，会获取该对象需要的依赖
//需要db这个依赖，从配置文件中获取
func (i *Impl) Init() error{
	i.db = conf.C().Mysql.GetDB()
	return nil
}
