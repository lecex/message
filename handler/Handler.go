package handler

import (
	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/lecex/core/env"
	messagePB "github.com/lecex/message/proto/message"
	templatePB "github.com/lecex/message/proto/template"
	db "github.com/lecex/message/providers/database"
	"github.com/lecex/message/service/repository"
	"github.com/lecex/message/service/sms"
)

// Register 注册
func Register(Server server.Server) {
	repo := &repository.TemplateRepository{db.DB}
	smsHander := &sms.SmsHandler{
		env.Getenv("SMS_DRIVE", "aliyun"),
	}
	sms, err := smsHander.NewHandler()
	if err != nil {
		log.Log(err)
	}
	messagePB.RegisterMessageHandler(Server, &Message{repo, sms}) // 用户服务实现
	templatePB.RegisterTemplatesHandler(Server, &Template{repo})  // 用户服务实现
	templatePB.RegisterConfigsHandler(Server, &Config{repo})      // 用户服务实现
}
