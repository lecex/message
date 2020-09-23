package hander

import (
	"context"
	"fmt"

	pb "github.com/lecex/message/proto/drive"
	"github.com/lecex/message/service/repository"
)

// Drive 消息事件模板结构
type Drive struct {
	Repo repository.Drive
}

// List 获取所有消息事件模板
func (srv *Drive) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	drives, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.Drives = drives
	res.Total = total
	return err
}

// Get 获取消息事件模板
func (srv *Drive) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	nat, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Drive = nat
	return err
}

// Create 创建消息事件模板
func (srv *Drive) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加消息事件模板失败")
	}
	res.Valid = true
	return err
}

// Update 更新消息事件模板
func (srv *Drive) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新消息事件模板失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除消息事件模板
func (srv *Drive) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除消息事件模板失败")
	}
	res.Valid = valid
	return err
}
