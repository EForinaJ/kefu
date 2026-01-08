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
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	obj, err := tx.Model(dao.SysWitkey.Table()).
		Where(dao.SysWitkey.Columns().Id, req.Id).
		Fields(dao.SysWitkey.Columns().GameId, dao.SysWitkey.Columns().TitleId).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	oldGame, err := tx.Model(dao.SysGame.Table()).
		Where(dao.SysGame.Columns().Id, obj.GMap().Get(dao.SysWitkey.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	oldTitle, err := tx.Model(dao.SysTitle.Table()).
		Where(dao.SysTitle.Columns().Id, obj.GMap().Get(dao.SysWitkey.Columns().TitleId)).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	newGame, err := tx.Model(dao.SysGame.Table()).
		Where(dao.SysGame.Columns().Id, req.GameId).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	newTitle, err := tx.Model(dao.SysTitle.Table()).
		Where(dao.SysTitle.Columns().Id, req.TitleId).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	_, err = tx.Model(dao.SysWitkeyLog.Table()).
		Data(g.Map{
			dao.SysWitkeyLog.Columns().ManageId:   ctx.Value("userId"),
			dao.SysWitkeyLog.Columns().WitkeyId:   req.Id,
			dao.SysWitkeyLog.Columns().Type:       consts.WitkeyLogTypeChangeTitle,
			dao.SysWitkeyLog.Columns().Content:    "从" + oldGame.String() + "领域中的" + oldTitle.String() + "头衔，变更到" + newGame.String() + "领域中" + newTitle.String() + "头衔",
			dao.SysWitkeyLog.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = tx.Model(dao.SysWitkey.Table()).
		Where(dao.SysWitkey.Columns().Id, req.Id).
		Data(g.Map{
			dao.SysWitkey.Columns().GameId:  req.GameId,
			dao.SysWitkey.Columns().TitleId: req.TitleId,
		}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}
