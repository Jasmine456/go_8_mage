package common

//默认情况下，GORM使用ID作为主键，使用结构体名的蛇形复数作为表名，字段名的蛇形作为列名
type User struct {
	Id       int
	Uid      uint32
	Keywords string
	Degree   string
	Gender   string
	City     string
}

//使用TableName() 来修改默认的表名
func (User) TableName() string {
	return "user"
}

type UserCollection []User

type SocketRequest struct {
	MustKeys, ShouldKeys []string
	OnFlag, OffFlag      uint32
	OrFlags              []uint32
}

//搜索请求
type SearchRequest struct {
	MustKeys   []string `form:"must_keys"`   //必须同时包含这些词
	ShouldKeys []string `form:"should_keys"` //只需要包含这些词中的任意一个或几个
	Gender     string   `form:"gender"`      //必须是指定的性别
	City       string   `form:"city"`        //必须在指定的城市
	Degrees    []string `form:"degrees"`     //学历需要是指定集合中的一个

}
