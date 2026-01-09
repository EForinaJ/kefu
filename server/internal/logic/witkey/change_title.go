package witkey

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_witkey "server/internal/type/witkey/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChangeTitle implements service.IWitkey.
func (s *sWitkey) ChangeTitle(ctx context.Context, req *dto_witkey.ChangeTitle) (err error) {

	witkey, err := dao.SysWitkey.Ctx(ctx).Where(dao.SysWitkey.Columns().Id, req.Id).
		Fields(dao.SysWitkey.Columns().GameId, dao.SysWitkey.Columns().TitleId).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	_, err = dao.SysWitkeyChangeTitle.Ctx(ctx).
		Data(g.Map{
			dao.SysWitkeyChangeTitle.Columns().ManageId:    ctx.Value("userId"),
			dao.SysWitkeyChangeTitle.Columns().WitkeyId:    req.Id,
			dao.SysWitkeyChangeTitle.Columns().AfterTitle:  witkey.GMap().Get(dao.SysWitkey.Columns().TitleId),
			dao.SysWitkeyChangeTitle.Columns().AfterGame:   witkey.GMap().Get(dao.SysWitkey.Columns().GameId),
			dao.SysWitkeyChangeTitle.Columns().BeforeTitle: req.TitleId,
			dao.SysWitkeyChangeTitle.Columns().BeforeGame:  req.GameId,
			dao.SysWitkeyChangeTitle.Columns().Description: req.Description,
			dao.SysWitkeyChangeTitle.Columns().CreateTime:  gtime.Now(),
			dao.SysWitkeyChangeTitle.Columns().Status:      consts.StatusApply,
		}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
