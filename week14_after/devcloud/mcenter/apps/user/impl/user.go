package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/apps/user"
)

// 创建用户
// 管理员创建主账号
// 也有可能是主账号，创建子账号
func (i *impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	// 如果是管理员创建的账号需要用户自己重置密码
	if req.CreateBy.IsIn(user.CREATE_BY_ADMIN) {
		u.Password.SetNeedReset("admin created user need reset when first login")
	}

	// 持久化
	if err := i.save(ctx, u); err != nil {
		return nil, err
	}

	// 为了安全, 密码对象清除掉
	u.Password = nil
	u.Spec.Password = ""
	return u, nil
}

// 查询用户列表
func (i *impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	r := newQueryRequest(req)
	resp, err := i.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	set := user.NewUserSet()
	// 循环, 读取数据
	if !req.SkipItems {
		for resp.Next(ctx) {
			ins := user.NewDefaultUser()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode user error, error is %s", err)
			}
			// 为了防止user有敏感信息
			ins.Desensitize()
			set.Add(ins)
		}
	}

	// count
	count, err := i.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get user count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 查询用户详情
func (i *impl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	// 构建mongodb过滤器
	filter := bson.M{}
	switch req.DescribeBy {
	case user.DESCRIBE_BY_USER_ID:
		filter["_id"] = req.Id
	case user.DESCRIBE_BY_USER_NAME:
		filter["spec.username"] = req.Username
	default:
		return nil, exception.NewBadRequest("unknow desribe by %s", req.DescribeBy)
	}

	ins := user.NewDefaultUser()
	if err := i.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("user %s error, %s", req, err)
	}
	return ins, nil
}

//   更新用户
func (i *impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return nil, nil
}

//    删除用户
func (i *impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	return nil, nil
}

// 修改用户密码, 用户需要知道原先密码
func (i *impl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.Password, error) {
	return nil, nil
}

// 重置密码, 无需知道原先密码, 主账号执行
func (i *impl) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) (*user.Password, error) {
	return nil, nil
}
