package repository

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-micro/v2/util/log"

	"github.com/jinzhu/gorm"
	"github.com/lecex/core/util"
	pb "github.com/lecex/message/proto/template"
)

//Template 仓库接口
type Template interface {
	Create(template *pb.Template) (*pb.Template, error)
	Get(template *pb.Template) (*pb.Template, error)
	List(req *pb.ListQuery) ([]*pb.Template, error)
	Total(req *pb.ListQuery) (int64, error)
	Update(template *pb.Template) (bool, error)
	Delete(template *pb.Template) (bool, error)
}

// TemplateRepository 模版仓库
type TemplateRepository struct {
	DB *gorm.DB
}

// List 获取所有模版信息
func (repo *TemplateRepository) List(req *pb.ListQuery) (templates []*pb.Template, err error) {
	db := repo.DB
	limit, offset := util.Page(req.Limit, req.Page) // 分页
	sort := util.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&templates).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return templates, nil
}

// Total 获取所有模版查询总量
func (repo *TemplateRepository) Total(req *pb.ListQuery) (total int64, err error) {
	templates := []pb.Template{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&templates).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取模版信息
func (repo *TemplateRepository) Get(template *pb.Template) (*pb.Template, error) {
	if err := repo.DB.Where(&template).Find(&template).Error; err != nil {
		return nil, err
	}
	return template, nil
}

// Create 创建模版
// bug 无模版名创建模版可能引起 bug
func (repo *TemplateRepository) Create(template *pb.Template) (*pb.Template, error) {
	err := repo.DB.Create(template).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return template, fmt.Errorf("注册模版失败")
	}
	return template, nil
}

// Update 更新模版
func (repo *TemplateRepository) Update(template *pb.Template) (bool, error) {
	if template.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Template{
		Id: template.Id,
	}
	err := repo.DB.Model(id).Updates(template).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除模版
func (repo *TemplateRepository) Delete(template *pb.Template) (bool, error) {
	id := &pb.Template{
		Id: template.Id,
	}
	err := repo.DB.Delete(id).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
