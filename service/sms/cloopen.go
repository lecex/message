package sms

//package main
import (
	sdk "github.com/bigrocs/cloopen"
	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
)

// Cloopen 创建容联云通讯类
type Cloopen struct {
	AccountSid   string
	AccountToken string
	AppID        string
}

// Send 获取所有消息事件信息
func (srv *Cloopen) Send(req *pb.Request, t *tpd.Template) (valid bool, err error) {
	client := srv.Client()
	request := &sdk.Request{
		Mobile:       req.Addressee,
		TemplateCode: t.TemplateCode,
		Datas: []string{
			req.QueryParams["code"],
			req.QueryParams["time"],
		},
	}
	valid, err = client.Send(request)
	return valid, err
}

// Client 初始化连接
func (srv *Cloopen) Client() *sdk.Cloopen {
	return &sdk.Cloopen{
		AccountSid:   srv.AccountSid,
		AppID:        srv.AppID,
		AccountToken: srv.AccountToken,
	}
}
