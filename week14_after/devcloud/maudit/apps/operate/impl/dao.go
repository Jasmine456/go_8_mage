package impl

import (
	"context"
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/maudit/apps/operate"
)

func (s *service) save(ctx context.Context,ins *operate.OperateLog)error{
	_,err:=s.col.InsertOne(ctx,ins)
	if err != nil {
		return err
	}
	return nil
}
