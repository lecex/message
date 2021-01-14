package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/micro/go-micro/v2/util/log"

	conPB "github.com/lecex/message/proto/config"
	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
	db "github.com/lecex/message/providers/database"
	"github.com/lecex/message/service/repository"
	"github.com/lecex/message/service/sms"
	"github.com/lecex/message/service/wechat"
)

// Message 消息服务
type Message struct {
	Repo repository.Template
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
		sms, err := srv.sms()
		if err != nil {
			log.Log(err)
			return err
		}
		valid, err = sms.Send(req, templates)
		if err != nil {
			log.Log(err)
			return err
		}
	}
	if srv.inSliceString(Type, "wechat") {
		wechat, err := srv.wechat()
		if err != nil {
			log.Log(err)
			return err
		}
		valid, err = wechat.Template(req, templates)
		if err != nil {
			log.Log(err)
			return err
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

// sms 构建 sms 短信服务结构
func (srv *Message) sms() (h sms.Sms, err error) {
	con, err := srv.getConfig()
	if err != nil {
		return h, err
	}
	switch con.Sms.Drive {
	case "aliyun":
		h = &sms.Aliyun{
			RegionID:        "default",
			AccessKeyID:     con.Sms.Aliyun.AccessKeyID,
			AccessKeySecret: con.Sms.Aliyun.AccessKeySecret,
			SignName:        con.Sms.Aliyun.SignName,
		}
	case "cloopen":
		h = &sms.Cloopen{
			AppID:        con.Sms.Cloopen.AppID,
			AccountSid:   con.Sms.Cloopen.AccountSid,
			AccountToken: con.Sms.Cloopen.AccountToken,
		}
	default:
		return nil, fmt.Errorf("未找 %s SMS 驱动", con.Sms.Drive)
	}
	return h, err
}

// wechat 构建 wechat 模板消息
func (srv *Message) wechat() (h *wechat.Wechat, err error) {
	con, err := srv.getConfig()
	if err != nil {
		return h, err
	}
	h = &wechat.Wechat{
		AppId:       con.Wechat.Appid,
		Secret:      con.Wechat.Secret,
		AccessToken: con.Wechat.AccessToken,
	}
	return h, err
}

// config 初始化配置等
func (srv *Message) getConfig() (*conPB.Config, error) {
	res := &conPB.Response{}
	h := Config{&repository.ConfigRepository{db.DB}}
	err := h.Get(context.TODO(), &conPB.Request{}, res)
	return res.Config, err
}
