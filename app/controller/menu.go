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
 * 菜单管理-控制器
 * @author 半城风雨
 * @since 2021/7/19
 * @File : menu
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gutil"
)

var Menu = new(menuCtl)

type menuCtl struct{}

func (c *menuCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询方法
	list := service.Menu.GetList(req)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *menuCtl) Detail(r *ghttp.Request) {
	// 查询记录
	id := r.GetQueryInt64("id")
	if id > 0 {
		info, err := dao.Menu.FindOne("id=?", id)
		if err != nil || info == nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 菜单信息
		menu := model.MenuInfoVo{}
		menu.Menu = *info
		// 获取权限节点
		if info.Type == 0 {
			// 获取角色菜单权限列表
			var menuList []model.Menu
			dao.Menu.Where("parent_id=?", info.Id).Where("type=1").Where("mark=1").Structs(&menuList)
			checkedList := gutil.ListItemValuesUnique(&menuList, "Sort")
			if len(checkedList) > 0 {
				menu.CheckedList = checkedList
			} else {
				menu.CheckedList = make([]interface{}, 0)
			}
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "查询成功",
			Data: menu,
		})
	}
}

func (c *menuCtl) Add(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用添加方法
	id, err := service.Menu.Add(req, utils.Uid(r))
	if err != nil || id == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回成功提示
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *menuCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := service.Menu.Update(req, utils.Uid(r))
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回成功提示
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *menuCtl) Delete(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.Menu.Delete(req.Ids)
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}
