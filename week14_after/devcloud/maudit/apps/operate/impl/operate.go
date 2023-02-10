package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/apps/operate"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (s *service) SaveOperateLog(ctx context.Context, req *operate.OperateLog) (*operate.SaveOperateLogResponse, error) {

	// 1. 校验用户的请求合法
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2. 业务逻辑的处理
	// 2.1 构造出需要的对象
	if req.Time == 0 {
		req.Time = time.Now().UnixMicro()
	}
	// 保存入库
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}

	return &operate.SaveOperateLogResponse{Message: "ok"}, nil
}
func (s *service) QueryOperateLog(ctx context.Context, req *operate.QueryOperateLogRequest) (*operate.OperateLogSet, error) {
	// 1.校验参数
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2. 构造过滤条件
	filter:=bson.M{}
	if req.Username !=""{
		filter["username"]=req.Username
	}

	// 3. 构造分页
	pageSize := int64(req.Page.PageSize)
	skip := int64(req.Page.PageSize) * int64(req.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort: bson.D{
			{Key: "time", Value: -1},
		},
		Limit: &pageSize,
		Skip:  &skip,
	}

	// 4. 执行mongodb 查询
	cursor,err:=s.col.Find(ctx, filter, opt)
	if err != nil {
		return nil,err
	}
	// 4.1 处理数据库查询出来的结果
	set:=operate.NewOperateLogSet()
	for cursor.Next(ctx){
		ins:=operate.NewOperateLog()
		if err:=cursor.Decode(ins);err!=nil{
			return nil,exception.NewInternalServerError("decode operate log error,error is %s",err)
		}
		set.Add(ins)
	}
	// 4.2 统计总数据 COUNT(*) WHERE ...
	count,err:=s.col.CountDocuments(ctx,filter)
	if err != nil {
		return nil,exception.NewInternalServerError("get operate log  count error,error is %s",err)
	}
	set.Total = count
	// 5. 返回结果
	return set, nil
}
