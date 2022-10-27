package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (i *impl) save(ctx context.Context, u *user.User) error {

	if _, err := i.col.InsertOne(ctx, u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Id, err)
	}
	return nil
}

func newQueryRequest(r *user.QueryUserRequest) *queryRequest {
	return &queryRequest{
		r,
	}
}

// 构建mongodb查询参数
type queryRequest struct {
	*user.QueryUserRequest
}

// 分页
func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		// DESC -1
		Sort: bson.D{
			{Key: "create_at", Value: -1},
		},
		// LIMIT offset,limit
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

// 过滤条件
func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Domain != "" {
		filter["spec.domain"] = r.Domain
	}
	if r.Provider != nil {
		filter["spec.provider"] = r.Provider
	}
	if r.Type != nil {
		filter["spec.type"] = r.Type
	}
	// IN ()
	if len(r.UserIds) > 0 {
		filter["_id"] = bson.M{"$in": r.UserIds}
	}

	return filter
}
