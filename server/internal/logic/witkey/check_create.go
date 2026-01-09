package witkey

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
)

// CheckExist implements service.IWitkey.
func (s *sWitkey) CheckCreate(ctx context.Context, req *dto_witkey.Create) (err error) {
	res, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Phone, req.Phone).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if res {
		return utils_error.Err(response.FAILD, "该手机号已入职")
	}

	res, err = dao.SysWitkeyOnboarding.Ctx(ctx).
		Where(dao.SysWitkeyOnboarding.Columns().Phone, req.Phone).
		Where(dao.SysWitkeyOnboarding.Columns().Status, consts.StatusApply).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if res {
		return utils_error.Err(response.FAILD, "该手机号入职正在审核中，请不要重复提交")
	}
	return
}
