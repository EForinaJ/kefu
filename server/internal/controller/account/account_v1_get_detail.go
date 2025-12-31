package account

import (
	"context"

	v1 "kefu-server/api/account/v1"
	"kefu-server/internal/service"
)

func (c *ControllerV1) GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error) {
	detail, err := service.Account().GetDetail(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetDetailRes{
		Detail: detail,
	}
	return
}
