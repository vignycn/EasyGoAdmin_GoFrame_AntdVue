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
 * 职级管理-控制器
 * @author 半城风雨
 * @since 2021/5/20
 * @File : level
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Level = new(levelCtl)

type levelCtl struct{}

// 查询数据列表
func (c *levelCtl) List(r *ghttp.Request) {
	// 请求参数
	var req *model.LevelQueryReq
	// 请求验证
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表函数
	list, count, err := service.Level.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果集
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

// 添加职级
func (c *levelCtl) Add(r *ghttp.Request) {
	var req *model.LevelAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	id, err := service.Level.Add(req, utils.Uid(r))
	if err != nil || id <= 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 保存成功
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}

// 更新职级
func (c *levelCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.LevelUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	result, err := service.Level.Update(req, utils.Uid(r))
	if err != nil || result == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *levelCtl) Delete(r *ghttp.Request) {
	var req *model.LevelDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.Level.Delete(req.Ids)
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

func (c *levelCtl) Status(r *ghttp.Request) {
	var req *model.LevelStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	result, err := service.Level.Status(req, utils.Uid(r))
	if err != nil || result == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 保存成功
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}

func (c *levelCtl) GetLevelList(r *ghttp.Request) {
	// 查询职级列表
	list, _ := dao.Level.Where("status=1 and mark=1").Order("sort asc").All()
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}

func (c *levelCtl) ImportExcel(r *ghttp.Request) {
	// 上传文件
	upFile := r.GetUploadFile("file")
	if upFile == nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  "上传文件不能为空",
		})
	}
	// 调用文件上传导入方法
	count, err := service.Level.ImportExcel(upFile, utils.Uid(r))
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  fmt.Sprintf("本次共导入【%d】条数据", count),
	})
}

func (c *levelCtl) ExportExcel(r *ghttp.Request) {
	// 请求参数
	var req *model.LevelQueryReq
	// 请求验证
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表函数
	fileURL, err := service.Level.GetExcelList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  "导出Excel失败",
		})
	}

	// 返回结果集
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "导出成功",
		Data: fileURL,
	})
}

func (c *levelCtl) DownloadExcel(r *ghttp.Request) {
	// 模板地址
	fileURL := utils.GetImageUrl("/level.xlsx")
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "下载成功",
		Data: fileURL,
	})
}
