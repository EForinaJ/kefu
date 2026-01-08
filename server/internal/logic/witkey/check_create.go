package witkey

import (
	"context"
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
		return utils_error.Err(response.FAILD, "手机号已存在")
	}
	return
}
