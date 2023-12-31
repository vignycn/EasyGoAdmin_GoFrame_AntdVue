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
 * 登录日志
 * @author 半城风雨
 * @since 2021/8/14
 * @File : loginlog
 */
package middleware

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/mssola/user_agent"
)

func LoginLog(r *ghttp.Request) {
	// 后置中间件
	r.Middleware.Next()
	// 中间件处理逻辑
	fmt.Println("登录日志中间件")

	// 记录登录、退出日志
	urlItem := []string{"/login", "/logout"}
	if utils.InStringArray(r.RequestURI, urlItem) {

		// 获取浏览器信息
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)

		// 实例化对象
		var entity model.LoginLog
		entity.Method = r.Method
		entity.OperUrl = r.URL.String()
		entity.OperIp = r.GetClientIp()
		entity.OperLocation = utils.GetIpCity(entity.OperIp)
		entity.RequestParam = string(r.GetBody())
		entity.Status = 0
		// 操作系统
		entity.Os = ua.OS()
		entity.Browser, _ = ua.Browser()
		entity.UserAgent = r.UserAgent()
		entity.CreateTime = gtime.Now()
		entity.Mark = 1
		if r.RequestURI == "/login" {
			// 登录成功
			var jsonObj map[string]interface{}
			json.Unmarshal(r.GetBody(), &jsonObj)
			// 获取用户信息
			user, err := dao.User.Where("username=?", jsonObj["username"]).FindOne()
			if err != nil {
				return
			}
			entity.Type = 1
			entity.Username = user.Username
			entity.CreateUser = user.Id
		} else if r.RequestURI == "/logout" {
			// 注销成功
			entity.Type = 3
			entity.Username = utils.UInfo(r).Username
			entity.CreateUser = utils.Uid(r)
		}
		// 插入记录
		dao.LoginLog.Insert(entity)
	}
}
