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
 * 职级管理-服务类
 * @author 半城风雨
 * @since 2021/7/13
 * @File : level
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 中间件管理服务
var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req *model.LevelQueryReq) ([]model.Level, int, error) {
	query := dao.Level.Clone()
	query = query.Where("mark=1")
	if req != nil {
		// 职级名称查询
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
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
	var list []model.Level
	query.Structs(&list)
	return list, count, nil
}

func (s *levelService) Add(req *model.LevelAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Level
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1
	// 插入数据
	result, err := dao.Level.Insert(entity)
	if err != nil {
		return 0, err
	}
	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

func (s *levelService) Update(req *model.LevelUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Level.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	info.Name = req.Name
	info.Status = req.Status
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()
	result, err := dao.Level.Save(info)
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

// 删除
func (s *levelService) Delete(Ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	idsArr := convert.ToInt64Array(Ids, ",")
	result, err := dao.Level.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响行数
	rows, err := result.RowsAffected()
	return rows, nil
}

func (s *levelService) Status(req *model.LevelStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Level.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Level.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Level.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *levelService) ImportExcel(uploadFile *ghttp.UploadFile, userId int) (int, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 临时存储目录
	savePath := utils.TempPath() + "/" + gtime.Now().Format("Ymd")
	// 上传文件
	fileName, err := uploadFile.Save(savePath, true)
	if err != nil {
		return 0, errors.New("文件上传失败")
	}
	// 文件绝对路径
	filePath := filepath.Join(savePath, "/", fileName)
	// 读取Excel文件
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, errors.New("excel文件读取失败")
	}
	// 读取第一张Sheet表
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return 0, errors.New("excel文件读取失败")
	}
	// 计数器
	totalNum := 0
	// Excel文件头，此处必须与Excel模板头保持一致
	excelHeader := []string{"职级名称", "职级状态", "显示顺序"}
	// 循环遍历读取的数据源
	for rows.Next() {
		// Excel列对象
		item, err2 := rows.Columns()
		if err2 != nil {
			return 0, errors.New("excel文件解析异常")
		}
		// 读取的列数与Excel头列数不等则跳过读取下一条
		if len(item) != len(excelHeader) {
			continue
		}
		// 如果是标题栏则跳过
		if item[1] == "职级状态" {
			continue
		}
		// 职级名称
		name := item[0]
		// 职级状态
		status := 1
		if item[1] == "正常" {
			status = 1
		} else {
			status = 2
		}
		// 显示顺序
		sort, _ := strconv.Atoi(item[2])
		// 实例化职级导入对象
		level := model.Level{
			Name:       name,
			Status:     status,
			Sort:       sort,
			CreateUser: userId,
			CreateTime: gtime.Now(),
			UpdateUser: userId,
			UpdateTime: gtime.Now(),
			Mark:       1,
		}
		// 插入职级数据
		if _, err := dao.Level.Insert(level); err != nil {
			return 0, err
		}
		// 计数器+1
		totalNum++
	}
	return totalNum, nil
}

func (s *levelService) GetExcelList(req *model.LevelQueryReq) (string, error) {
	query := dao.Level.Clone()
	query = query.Where("mark=1")
	if req != nil {
		// 职级名称查询
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 查询记录总数
	_, err := query.Count()
	if err != nil {
		return "", err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.Level
	query.Structs(&list)

	// 循环遍历处理数据源
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "职级名称", "职级状态", "排序", "创建时间"})
	for i, v := range list {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			v.Id,
			v.Name,
			v.Status,
			v.Sort,
			v.CreateTime,
		})
	}
	// 定义文件名称
	fileName := fmt.Sprintf("%s.xlsx", time.Now().Format("20060102150405"))
	// 设置Excel保存路径
	filePath := filepath.Join(utils.TempPath(), "/", fileName)
	err2 := excel.SaveAs(filePath)
	// 获取文件URL地址
	fileURL := utils.GetImageUrl(strings.ReplaceAll(filePath, utils.UploadPath(), ""))
	// 返回结果
	return fileURL, err2
}
