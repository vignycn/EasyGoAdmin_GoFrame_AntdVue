// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2023 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: https://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 配置数据-服务类
 * @author 半城风雨
 * @since 2021/7/21
 * @File : config_data
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var ConfigData = new(configDataService)

type configDataService struct{}

func (s *configDataService) GetList(req *model.ConfigDataPageReq) ([]model.ConfigDataVo, int, error) {
	// 创建查询对象
	query := dao.ConfigData.Where("config_id=?", req.ConfigId).Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典项名称
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
		}
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.ConfigData
	query.Structs(&list)

	var result = make([]model.ConfigDataVo, 0)
	for _, v := range list {
		typeName, ok := common.CONFIG_DATA_TYPE_LIST[v.Type]
		item := model.ConfigDataVo{}
		item.ConfigData = v
		if ok {
			item.TypeName = typeName
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *configDataService) Add(req *model.ConfigDataAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.ConfigData
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.ConfigData.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *configDataService) Update(req *model.ConfigDataUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.ConfigData.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Title = req.Title
	info.Code = req.Code
	info.Value = req.Value
	info.Options = req.Options
	info.ConfigId = req.ConfigId
	info.Type = req.Type
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.ConfigData.Save(info)
	if err != nil {
		return 0, err
	}

	// 获取受影响函数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *configDataService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.ConfigData.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响函数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *configDataService) Status(req *model.ConfigDataStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.ConfigData.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.ConfigData.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.ConfigData.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
