package repository

import (
	"encoding/json"
	"fmt"

	// 公共引入

	"github.com/jinzhu/gorm"
	pb "github.com/lecex/message/proto/config"
	"github.com/micro/go-micro/v2/util/log"
)

//Config 仓库接口
type Config interface {
	Get() (*pb.Config, error)
	Update(config *pb.Config) (bool, error)
}

// ConfigRepository 模版仓库
type ConfigRepository struct {
	DB *gorm.DB
}

type Configs struct {
	Name  string
	Value string
}

// Get 获取模版信息
func (repo *ConfigRepository) Get() (config *pb.Config, err error) {
	con := &Configs{
		Name: "config",
	}
	if err := repo.DB.Where(&con).Find(&con).Error; err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(con.Value), &config)
	return config, nil
}

// Update 更新模版
func (repo *ConfigRepository) Update(config *pb.Config) (bool, error) {
	c, err := json.Marshal(config)
	if err != nil {
		log.Log(err)
		return false, err
	}
	con := &Configs{
		Name:  "config",
		Value: string(c),
	}
	err = repo.DB.Model(&con).Updates(&con).Error
	fmt.Println(err)
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
