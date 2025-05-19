package svc

import (
	"github.com/zeromicro/go-zero/tools/goctl/demo/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
