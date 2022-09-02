package impl

import (
	"context"
	"go_8_mage/week14/vblog/api/apps/blog"
)

//对象的保存
//1. 采用原生的SQL,INSERT INTO table_name (filed_name,...) VALUES(filed_value,...)
//2. 采用GORM
// context,网络比较慢，用户等待3秒，用户取消，数据库的操作也必须支持需求，必须传入ctx
func (i *Impl)save(ctx context.Context,ins *blog.Blog) error{

	return i.DB().WithContext(ctx).Create(ins).Error
}
