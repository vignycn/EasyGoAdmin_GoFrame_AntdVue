// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
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
