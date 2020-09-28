package handler

import (
	server "github.com/micro/go-micro/v2/server"

	configPB "github.com/lecex/message/proto/config"
	messagePB "github.com/lecex/message/proto/message"
	templatePB "github.com/lecex/message/proto/template"
	db "github.com/lecex/message/providers/database"
	"github.com/lecex/message/service/repository"
)

// Register 注册
func Register(Server server.Server) {
	repo := &repository.TemplateRepository{db.DB}

	configPB.RegisterConfigsHandler(Server, &Config{&repository.ConfigRepository{db.DB}}) //
	messagePB.RegisterMessageHandler(Server, &Message{repo})                              //
	templatePB.RegisterTemplatesHandler(Server, &Template{repo})                          //
}
