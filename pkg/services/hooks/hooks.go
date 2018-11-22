package hooks

import (
	"github.com/aergoio/grafana/pkg/api/dtos"
	"github.com/aergoio/grafana/pkg/registry"
)

type IndexDataHook func(indexData *dtos.IndexViewData)

type HooksService struct {
	indexDataHooks []IndexDataHook
}

func init() {
	registry.RegisterService(&HooksService{})
}

func (srv *HooksService) Init() error {
	return nil
}

func (srv *HooksService) AddIndexDataHook(hook IndexDataHook) {
	srv.indexDataHooks = append(srv.indexDataHooks, hook)
}

func (srv *HooksService) RunIndexDataHooks(indexData *dtos.IndexViewData) {
	for _, hook := range srv.indexDataHooks {
		hook(indexData)
	}
}
