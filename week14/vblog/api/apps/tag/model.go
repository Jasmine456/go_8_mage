package tag


type TagSet struct {
	Items []*Tag
}

type Tag struct {
	// 标签ID
	Id int
	// 创建时间: 用于排序
	CreateAt int64
	// Tag的具体数据
	*CreateTagRequest
}

type CreateTagRequest struct {
	// 关联的博客
	BlogId int
	// 标签名称
	Key string
	// 标签的value
	Value string
	// 标签的颜色
	Color string
}
