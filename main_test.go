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
					AppID:        "8a48b551506fd26f01509405471a6db8",
					AccountSid:   "aaf98f895069246a01506a9770ea0268",
					AccountToken: "3fd8b18597d346c48631821abc00b138",
				},
			},
			Wechat: &conPB.Wechat{
				Appid:  "8a48b551506fd26f01509405471a6db8",
				Secret: "aaf98f895069246a01506a9770ea0268",
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
		Event:     "register_verify",
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
		Event:     "register_verify",
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
