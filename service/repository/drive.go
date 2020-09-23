package repository

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-micro/v2/util/log"

	"github.com/jinzhu/gorm"
	"github.com/lecex/core/uitl"
	pb "github.com/lecex/message/proto/drive"
)

//Drive 仓库接口
type Drive interface {
	Create(drive *pb.Drive) (*pb.Drive, error)
	Get(drive *pb.Drive) (*pb.Drive, error)
	List(req *pb.ListQuery) ([]*pb.Drive, error)
	Total(req *pb.ListQuery) (int64, error)
	Update(drive *pb.Drive) (bool, error)
	Delete(drive *pb.Drive) (bool, error)
}

// DriveRepository 模版仓库
type DriveRepository struct {
	DB *gorm.DB
}

// List 获取所有模版信息
func (repo *DriveRepository) List(req *pb.ListQuery) (drives []*pb.Drive, err error) {
	db := repo.DB
	limit, offset := uitl.Page(req.Limit, req.Page) // 分页
	sort := uitl.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&drives).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return drives, nil
}

// Total 获取所有模版查询总量
func (repo *DriveRepository) Total(req *pb.ListQuery) (total int64, err error) {
	drives := []pb.Drive{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&drives).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取模版信息
func (repo *DriveRepository) Get(drive *pb.Drive) (*pb.Drive, error) {
	if err := repo.DB.Where(&drive).Find(&drive).Error; err != nil {
		return nil, err
	}
	return drive, nil
}

// Create 创建模版
// bug 无模版名创建模版可能引起 bug
func (repo *DriveRepository) Create(drive *pb.Drive) (*pb.Drive, error) {
	if exist := repo.Exist(drive); exist == true {
		return drive, fmt.Errorf("注册模版已存在")
	}
	err := repo.DB.Create(drive).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return drive, fmt.Errorf("注册模版失败")
	}
	return drive, nil
}

// Update 更新模版
func (repo *DriveRepository) Update(drive *pb.Drive) (bool, error) {
	if drive.Id == "" {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Drive{
		Id: drive.Id,
	}
	err := repo.DB.Model(id).Updates(drive).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除模版
func (repo *DriveRepository) Delete(drive *pb.Drive) (bool, error) {
	id := &pb.Drive{
		Id: drive.Id,
	}
	err := repo.DB.Delete(id).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
