package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/message/proto/template"
	"github.com/lecex/message/service/repository"
)

// Template 消息事件模板结构
type Template struct {
	Repo repository.Template
}

// List 获取所有消息事件模板
func (srv *Template) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	templates, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Templates = templates
	res.Total = total
	return err
}

// Get 获取消息事件模板
func (srv *Template) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	nat, err := srv.Repo.Get(req.Template)
	if err != nil {
		return err
	}
	res.Template = nat
	return err
}

// Create 创建消息事件模板
func (srv *Template) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Template)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加消息事件模板失败")
	}
	res.Valid = true
	return err
}

// Update 更新消息事件模板
func (srv *Template) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Template)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新消息事件模板失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除消息事件模板
func (srv *Template) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Template)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除消息事件模板失败")
	}
	res.Valid = valid
	return err
}
