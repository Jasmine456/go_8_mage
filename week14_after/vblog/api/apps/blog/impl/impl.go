package impl

import (
	"go_8_mage/week14/vblog/api/apps"
	"go_8_mage/week14/vblog/api/apps/tag"
	"go_8_mage/week14/vblog/api/conf"
	"gorm.io/gorm"
)

//有了IOC，去掉实例的构造函数
func NewImpl() *Impl {
	return &Impl{}
}

//依赖MySQL连接，能与MySQL交互
// 负责实现Blog service
type Impl struct {
	//通过配置文件
	//db *sql.DB
	db *gorm.DB

	tag tag.Service
}

func ( i *Impl) Name() string{
	return "blog"
}

func (i *Impl) DB()*gorm.DB{
	return i.db.Table(i.Name())
}

//当这个对象初始化的，会获取该对象需要的依赖
//需要db这个依赖，从配置文件中获取
func (i *Impl) Init() error {
	i.db = conf.C().Mysql.GetOrmDB().Debug()
	i.tag = apps.GetService(tag.AppName).(tag.Service)
	return nil
}

// import
func init(){
	apps.Registry(&Impl{})
}