package witkey

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IWitkey.
func (s *sWitkey) Create(ctx context.Context, req *dto_witkey.Create) (err error) {

	var entity *do.SysWitkeyOnboarding
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	if req.Password == "" {
		req.Password = "123456"
	}
	newSalt := grand.S(6)
	newToken := consts.SYSTEMNAME + req.Password + newSalt
	newToken = gmd5.MustEncryptString(newToken)
	entity.Salt = newSalt
	entity.Password = newToken
	if req.Name == "" {
		req.Name = "新用户" + grand.S(6)
	}
	// 转换为 gtime.Time
	if req.Birthday == 0 {
		req.Birthday = 631123200000
	}
	birthday := gtime.NewFromTimeStamp(req.Birthday / 1000).UTC()
	entity.Birthday = birthday
	entity.CreateTime = gtime.Now()
	entity.Status = consts.StatusApply
	entity.ManageId = ctx.Value("userId")
	_, err = dao.SysWitkeyOnboarding.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
