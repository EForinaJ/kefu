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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Create implements service.IWitkey.
func (s *sWitkey) Create(ctx context.Context, req *dto_witkey.Create) (err error) {
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

	var entity *do.SysWitkey
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
	entity.UpdateTime = gtime.Now()

	//  获取配置
	userInfo, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.UserSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	userJosn := gjson.New(userInfo)
	//  获取默认头像和封面
	if req.Avatar == "" {
		entity.Avatar = userJosn.Get("avatar")
	}
	entity.Status = req.Status
	_, err = tx.Model(dao.SysWitkey.Table()).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	// rid, err := rs.LastInsertId()
	// if err != nil {
	// 	return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	// }

	//  添加日志
	// _, err = tx.Model(dao.SysWitkeyLog.Table()).Data(g.Map{
	// 	dao.SysWitkeyLog.Columns().CreateTime: gtime.Now(),
	// 	dao.SysWitkeyLog.Columns().Content:    "客服添加威客",
	// 	dao.SysWitkeyLog.Columns().ManageId:   ctx.Value("userId"),
	// 	dao.SysWitkeyLog.Columns().WitkeyId:   rid,
	// 	dao.SysWitkeyLog.Columns().Type:       consts.WitkeyLogTypeCreate,
	// }).Insert()
	// if err != nil {
	// 	return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	// }
	return
}
