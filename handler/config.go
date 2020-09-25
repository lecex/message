package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/message/proto/config"
	"github.com/lecex/message/service/repository"
)

// Config 消息事件模板结构
type Config struct {
	Repo repository.Config
}

// Get 获取消息事件模板
func (srv *Config) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	nat, err := srv.Repo.Get(req.Config)
	if err != nil {
		return err
	}
	res.Config = nat
	return err
}

// Update 更新消息事件模板
func (srv *Config) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Config)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新消息事件模板失败")
	}
	res.Valid = valid
	return err
}
