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
	obj, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, req.Id).
		Fields(dao.SysWitkey.Columns().GameId, dao.SysWitkey.Columns().TitleId).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	_, err = dao.SysWitkeyTitle.Ctx(ctx).
		Where(dao.SysWitkeyTitle.Columns().Id, req.Id).
		Data(g.Map{
			dao.SysWitkeyTitle.Columns().ManageId:    ctx.Value("userId"),
			dao.SysWitkeyTitle.Columns().OldGameId:   obj.GMap().Get(dao.SysWitkey.Columns().GameId),
			dao.SysWitkeyTitle.Columns().OldTitleId:  obj.GMap().Get(dao.SysWitkey.Columns().TitleId),
			dao.SysWitkeyTitle.Columns().NewGameId:   req.GameId,
			dao.SysWitkeyTitle.Columns().NewTitleId:  req.TitleId,
			dao.SysWitkeyTitle.Columns().WitkeyId:    req.Id,
			dao.SysWitkeyTitle.Columns().Type:        req.Type,
			dao.SysWitkeyTitle.Columns().Description: req.Description,
			dao.SysWitkeyTitle.Columns().Status:      consts.StatusApply,
			dao.SysWitkeyTitle.Columns().CreateTime:  gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}
