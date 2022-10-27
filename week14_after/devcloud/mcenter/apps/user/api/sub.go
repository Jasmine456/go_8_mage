package api

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/apps/user"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	subHandler = &sub{}
)

//子账号管理
type sub struct {
	service user.Service
	log     logger.Logger
}

func (h *sub) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetGrpcApp(user.AppName).(user.Service)
	return nil
}

func (h *sub) Name() string {
	return "user/sub"
}

func (h *sub) Version() string {
	return "v1"
}

func (h *sub) Registry(ws *restful.WebService) {
	//tags := []string{"子账号管理"}

}
