package repository

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-micro/v2/util/log"

	"github.com/jinzhu/gorm"
	pb "github.com/lecex/message/proto/config"
)

//Config 仓库接口
type Config interface {
	Get(config *pb.Config) (*pb.Config, error)
	Update(config *pb.Config) (bool, error)
}

// ConfigRepository 模版仓库
type ConfigRepository struct {
	DB *gorm.DB
}

// Get 获取模版信息
func (repo *ConfigRepository) Get(config *pb.Config) (*pb.Config, error) {
	if err := repo.DB.Where(&config).Find(&config).Error; err != nil {
		return nil, err
	}
	return config, nil
}

// Create 创建模版
// bug 无模版名创建模版可能引起 bug
func (repo *ConfigRepository) Create(config *pb.Config) (*pb.Config, error) {
	err := repo.DB.Create(config).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return config, fmt.Errorf("注册模版失败")
	}
	return config, nil
}

// Update 更新模版
func (repo *ConfigRepository) Update(config *pb.Config) (bool, error) {
	err := repo.DB.Updates(config).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
