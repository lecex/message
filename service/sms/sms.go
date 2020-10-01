package sms

import (
	// 公共引入

	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
)

//Sms 短信发送接口
type Sms interface {
	Send(*pb.Request, *tpd.Template) (bool, error)
}
