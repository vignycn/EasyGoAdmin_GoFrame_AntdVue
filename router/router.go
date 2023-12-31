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
 * 系统路由
 * @author 半城风雨
 * @since 2021/7/26
 * @File : submit
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 跨域处理
	s.Use(middleware.CORS)
	// 登录验证中间件
	s.Use(middleware.CheckLogin)
	// 操作日志中间件
	s.Use(middleware.OperLog)
	// 登录日志中间件
	s.Use(middleware.LoginLog)

	/* 文件上传 */
	s.Group("/upload", func(group *ghttp.RouterGroup) {
		// 上传图片
		group.POST("/uploadImage", controller.Upload.UploadImage)
	})

	/* 登录注册 */
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", controller.Login.Login)
		group.ALL("/login", controller.Login.Login)
		group.GET("/captcha", controller.Login.Captcha)
		group.ALL("/updateUserInfo", controller.Index.UpdateUserInfo)
		group.ALL("/updatePwd", controller.Index.UpdatePwd)
		group.GET("/logout", controller.Index.Logout)
	})

	s.Group("index", func(group *ghttp.RouterGroup) {
		group.GET("/menu", controller.Index.Menu)
		group.GET("/user", controller.Index.User)
	})

	/* 用户管理 */
	s.Group("user", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.User.List)
		group.GET("/detail", controller.User.Detail)
		group.POST("/add", controller.User.Add)
		group.PUT("/update", controller.User.Update)
		group.DELETE("/delete/:ids", controller.User.Delete)
		group.PUT("/status", controller.User.Status)
		group.PUT("/resetPwd", controller.User.ResetPwd)
		group.GET("/checkUser", controller.User.CheckUser)
	})

	/* 职级管理 */
	s.Group("level", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Level.List)
		group.POST("/add", controller.Level.Add)
		group.PUT("/update", controller.Level.Update)
		group.DELETE("/delete/:ids", controller.Level.Delete)
		group.PUT("/status", controller.Level.Status)
		group.GET("/getLevelList", controller.Level.GetLevelList)
		group.GET("/exportExcel", controller.Level.ExportExcel)
		group.POST("/importExcel", controller.Level.ImportExcel)
		group.GET("/downloadExcel", controller.Level.DownloadExcel)
	})

	/* 岗位路由 */
	s.Group("position", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Position.List)
		group.POST("/add", controller.Position.Add)
		group.PUT("/update", controller.Position.Update)
		group.DELETE("/delete/:ids", controller.Position.Delete)
		group.PUT("/status", controller.Position.Status)
		group.GET("/getPositionList", controller.Position.GetPositionList)
	})

	/* 角色路由 */
	s.Group("role", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Role.List)
		group.POST("/add", controller.Role.Add)
		group.PUT("/update", controller.Role.Update)
		group.DELETE("/delete/:ids", controller.Role.Delete)
		group.PUT("/status", controller.Role.Status)
		group.GET("/getRoleList", controller.Role.GetRoleList)
	})

	/* 角色菜单权限 */
	s.Group("rolemenu", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.RoleMenu.Index)
		group.POST("/save", controller.RoleMenu.Save)
	})

	/* 部门管理 */
	s.Group("dept", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Dept.List)
		group.POST("/add", controller.Dept.Add)
		group.PUT("/update", controller.Dept.Update)
		group.DELETE("/delete/:ids", controller.Dept.Delete)
		group.GET("/getDeptList", controller.Dept.GetDeptList)
	})

	/* 菜单管理 */
	s.Group("menu", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Menu.List)
		group.GET("/detail", controller.Menu.Detail)
		group.POST("/add", controller.Menu.Add)
		group.PUT("/update", controller.Menu.Update)
		group.DELETE("/delete/:ids", controller.Menu.Delete)
	})

	/* 操作日志 */
	s.Group("operlog", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.OperLog.List)
	})

	/* 登录日志 */
	s.Group("loginlog", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.LoginLog.List)
		group.DELETE("/delete/:ids", controller.LoginLog.Delete)
	})

	/* 城市管理 */
	s.Group("city", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.City.List)
		group.POST("/add", controller.City.Add)
		group.PUT("/update", controller.City.Update)
		group.DELETE("/delete/:ids", controller.City.Delete)
		group.POST("/getChilds", controller.City.GetChilds)
	})

	/* 字典管理 */
	s.Group("dict", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Dict.List)
		group.POST("/add", controller.Dict.Add)
		group.PUT("/update", controller.Dict.Update)
		group.DELETE("/delete/:ids", controller.Dict.Delete)
	})

	/* 字典项管理 */
	s.Group("dictdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.DictData.List)
		group.POST("/add", controller.DictData.Add)
		group.PUT("/update", controller.DictData.Update)
		group.DELETE("/delete/:ids", controller.DictData.Delete)
	})

	/* 配置管理 */
	s.Group("config", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Config.List)
		group.POST("/add", controller.Config.Add)
		group.PUT("/update", controller.Config.Update)
		group.DELETE("/delete/:ids", controller.Config.Delete)
	})

	/* 配置项管理 */
	s.Group("configdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ConfigData.List)
		group.POST("/add", controller.ConfigData.Add)
		group.PUT("/update", controller.ConfigData.Update)
		group.DELETE("/delete/:ids", controller.ConfigData.Delete)
		group.PUT("/status", controller.ConfigData.Status)
	})

	/* 友链管理 */
	s.Group("link", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Link.List)
		group.POST("/add", controller.Link.Add)
		group.PUT("/update", controller.Link.Update)
		group.DELETE("/delete/:ids", controller.Link.Delete)
		group.PUT("/status", controller.Link.Status)
	})

	/* 站点管理 */
	s.Group("item", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Item.List)
		group.POST("/add", controller.Item.Add)
		group.PUT("/update", controller.Item.Update)
		group.DELETE("/delete/:ids", controller.Item.Delete)
		group.PUT("/status", controller.Item.Status)
		group.GET("/getItemList", controller.Item.GetItemList)
	})

	/* 栏目管理 */
	s.Group("itemcate", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ItemCate.List)
		group.POST("/add", controller.ItemCate.Add)
		group.PUT("/update", controller.ItemCate.Update)
		group.DELETE("/delete/:ids", controller.ItemCate.Delete)
		//group.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
		group.GET("/getCateList", controller.ItemCate.GetCateList)
	})

	/* 广告位管理 */
	s.Group("adsort", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.AdSort.List)
		group.POST("/add", controller.AdSort.Add)
		group.PUT("/update", controller.AdSort.Update)
		group.DELETE("/delete/:ids", controller.AdSort.Delete)
		group.GET("/getAdSortList", controller.AdSort.GetAdSortList)
	})

	/* 广告管理 */
	s.Group("ad", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Ad.List)
		group.POST("/add", controller.Ad.Add)
		group.PUT("/update", controller.Ad.Update)
		group.DELETE("/delete/:ids", controller.Ad.Delete)
		group.PUT("/status", controller.Ad.Status)
	})

	/* 通知管理 */
	s.Group("notice", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Notice.List)
		group.POST("/add", controller.Notice.Add)
		group.PUT("/update", controller.Notice.Update)
		group.DELETE("/delete/:ids", controller.Notice.Delete)
		group.PUT("/status", controller.Notice.Status)
	})

	/* 网站设置 */
	s.Group("configweb", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.ConfigWeb.Index)
		group.PUT("/save", controller.ConfigWeb.Save)
	})

	/* 会员等级 */
	s.Group("memberlevel", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.MemberLevel.List)
		group.POST("/add", controller.MemberLevel.Add)
		group.PUT("/update", controller.MemberLevel.Update)
		group.DELETE("/delete/:ids", controller.MemberLevel.Delete)
		group.GET("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	})

	/* 会员管理 */
	s.Group("member", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Member.List)
		group.POST("/add", controller.Member.Add)
		group.PUT("/update", controller.Member.Update)
		group.DELETE("/delete/:ids", controller.Member.Delete)
		group.PUT("/status", controller.Member.Status)
	})

	/* 统计分析 */
	s.Group("analysis", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Analysis.Index)
	})

	/* 代码生成器 */
	s.Group("generate", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Generate.List)
		group.POST("/generate", controller.Generate.Generate)
		group.POST("/batchGenerate", controller.Generate.BatchGenerate)
	})

}
