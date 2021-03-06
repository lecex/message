package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/message/handler"
	db "github.com/lecex/message/providers/database"
	"github.com/lecex/message/service/repository"

	conPB "github.com/lecex/message/proto/config"
	mesPB "github.com/lecex/message/proto/message"
)

func TestMessageConfigGet(t *testing.T) {
	req := &conPB.Request{}
	res := &conPB.Response{}
	h := handler.Config{&repository.ConfigRepository{db.DB}}
	err := h.Get(context.TODO(), req, res)
	fmt.Println("--------", req, res, err)
}

func TestMessageConfigUpdate(t *testing.T) {
	req := &conPB.Request{
		Config: &conPB.Config{
			Sms: &conPB.Sms{
				Drive: "cloopen",
				Cloopen: &conPB.Cloopen{
					AppID:        "",
					AccountSid:   "",
					AccountToken: "",
				},
			},
			Wechat: &conPB.Wechat{
				Appid:  "",
				Secret: "",
			},
		},
	}
	res := &conPB.Response{}
	h := handler.Config{&repository.ConfigRepository{db.DB}}
	err := h.Update(context.TODO(), req, res)
	t.Log(req, res, err)
}

func TestMessageSend(t *testing.T) {
	repo := &repository.TemplateRepository{db.DB}
	req := &mesPB.Request{
		Addressee: "13954386521",
		Event:     "sms_captcha",
		Type:      "wechat",
		QueryParams: `
			{
				"datas":[
					"654321",
					"5"
				]
			}
		`,
	}
	res := &mesPB.Response{}
	h := handler.Message{repo}
	err := h.Send(context.TODO(), req, res)
	t.Log(req, res, err)
}

func TestSmsMessageSend(t *testing.T) {
	repo := &repository.TemplateRepository{db.DB}
	req := &mesPB.Request{
		Addressee: "13954386521",
		Event:     "sms_captcha",
		Type:      "sms",
		QueryParams: `
			{
				"datas":[
					"654321",
					"5"
				]
			}
		`,
	}
	res := &mesPB.Response{}
	h := handler.Message{repo}
	err := h.Send(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
