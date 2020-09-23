package handler

import (
	"context"
	"strings"

	"github.com/lecex/message/service"
	"github.com/lecex/message/service/repository"
)

// Message 消息服务
type Message struct {
	Repo repository.Template
	Sms  service.Sms
}

// Send 发送
func (srv *Message) Send(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 查找对应模板信息
	templates, err := srv.Repo.Get(&tpd.Template{
		Event: req.Event,
	})

	if err != nil {
		return
	}

	if req.Type == "" {
		req.Type = templates.Type
	}
	Type := strings.Split(req.Type, ",")
	valid := false
	if srv.inSliceString(Type, "sms") {
		valid, err = srv.Sms.Send(req, templates)
		if err != nil {
			log.Log(err)
		}
	}
	res.Valid = valid
	return
}

// inSliceString 判断切片是否包含对应字符串
func (srv *Message) inSliceString(array []string, val string) bool {
	for _, item := range array {
		switch item {
		case val:
			return true
		}
	}
	return false
}
