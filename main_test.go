package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/message/handler"
	db "github.com/lecex/message/providers/database"
	"github.com/lecex/message/service/repository"

	conPB "github.com/lecex/message/proto/config"
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
					AppID:        "1",
					AccountSid:   "2",
					AccountToken: "3",
				},
			},
		},
	}
	res := &conPB.Response{}
	h := handler.Config{&repository.ConfigRepository{db.DB}}
	err := h.Update(context.TODO(), req, res)
	t.Log(req, res, err)
}
