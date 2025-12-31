package site

import (
	"context"

	v1 "kefu-server/api/site/v1"
	"kefu-server/internal/service"
)

func (c *ControllerV1) GetSite(ctx context.Context, req *v1.GetSiteReq) (res *v1.GetSiteRes, err error) {
	detail, err := service.Site().GetInfo(ctx)
	res = &v1.GetSiteRes{
		Detail: detail,
	}
	return
}
