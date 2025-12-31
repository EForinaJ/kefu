package auth

import (
	"context"

	v1 "kefu-server/api/auth/v1"
	"kefu-server/internal/consts"
	"kefu-server/internal/service"
	utils_error "kefu-server/internal/utils/error"
	utils_lock "kefu-server/internal/utils/lock"
	"kefu-server/internal/utils/response"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//检查是否用户被锁
	if utils_lock.CheckLock(ctx, consts.LoginLock+req.Phone) {
		return nil, utils_error.Err(response.LOGIN_ERROR, "账号已锁定，请30分钟后再试")
	}

	// 获取token
	token, err := service.Auth().Login(ctx, req.Login)
	if err != nil {
		return nil, utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
	}
	res = &v1.LoginRes{
		Token: token.(string),
	}
	return
}
