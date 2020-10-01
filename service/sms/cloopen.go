package sms

//package main
import (
	"encoding/json"

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

type QueryParams struct {
	Mobile       string
	TemplateCode string
	Datas        []string `json:"datas"`
}

// Send 获取所有消息事件信息
func (srv *Cloopen) Send(req *pb.Request, t *tpd.Template) (valid bool, err error) {
	queryParams := QueryParams{}
	err = json.Unmarshal([]byte(req.QueryParams), &queryParams)
	client := srv.NewClient()
	if req.Addressee != "" {
		queryParams.Mobile = req.Addressee
	}
	if t.TemplateCode != "" {
		queryParams.TemplateCode = t.TemplateCode
	}
	request := &sdk.Request{
		Mobile:       req.Addressee,
		TemplateCode: t.TemplateCode,
		Datas:        queryParams.Datas,
	}
	valid, err = client.Send(request)
	return valid, err
}

// Client 初始化连接
func (srv *Cloopen) NewClient() *sdk.Cloopen {
	return &sdk.Cloopen{
		AccountSid:   srv.AccountSid,
		AppID:        srv.AppID,
		AccountToken: srv.AccountToken,
	}
}
