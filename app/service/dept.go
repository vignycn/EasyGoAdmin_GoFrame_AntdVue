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
 * 部门管理-服务类
 * @author 半城风雨
 * @since 2021/7/16
 * @File : dept
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"reflect"
)

// 中间件管理服务
var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req *model.DeptPageReq) []model.Dept {
	// 创建查询对象
	query := dao.Dept.Where("mark=1")
	// 查询条件
	if req != nil {
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.Dept
	query.Structs(&list)
	return list
}

func (s *deptService) Add(req *model.DeptAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.Dept.Insert(entity)
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
func (s *deptService) Update(req *model.DeptUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Dept.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置参数
	info.Id = req.Id
	info.Name = req.Name
	info.Code = req.Code
	info.Fullname = req.Fullname
	info.Type = req.Type
	info.Pid = req.Pid
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.Dept.Save(info)
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *deptService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.Dept.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// 获取子级菜单
func (s *deptService) GetDeptTreeList() ([]*model.DeptTreeNode, error) {
	var deptNode model.DeptTreeNode
	// 查询列表
	list, _ := dao.Dept.Where("mark=1").Fields("id,name,pid").Order("sort asc").FindAll()
	makeDeptTree(list, &deptNode)
	return deptNode.Children, nil
}

// 递归生成分类列表
func makeDeptTree(cate []*model.Dept, tn *model.DeptTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &model.DeptTreeNode{}
			child.Dept = *c
			tn.Children = append(tn.Children, child)
			makeDeptTree(cate, child)
		}
	}
}

// 数据源转换
func (s *deptService) MakeList(data []*model.DeptTreeNode) map[int]string {
	deptList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			deptList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				deptList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					deptList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return deptList
}
