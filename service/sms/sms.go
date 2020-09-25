package sms

import (
	// 公共引入

	"fmt"

	"github.com/lecex/core/env"

	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
)

//Sms 短信发送接口
type Sms interface {
	Send(*pb.Request, *tpd.Template) (bool, error)
}

// SmsHandler sms 驱动
type SmsHandler struct {
	Drive string
}

// NewHandler 使用对应驱动使用 sms
func (s *SmsHandler) NewHandler() (handler Sms, err error) {
	switch s.Drive {
	case "aliyun":
		handler = &Aliyun{
			RegionID:        "default",
			AccessKeyID:     env.Getenv("SMS_KEY_ID", ""),
			AccessKeySecret: env.Getenv("SMS_KEY_SECRET", ""),
			SignName:        env.Getenv("SMS_SIGN_NAME", "阿里云短信测试专用"),
		}
	case "cloopen":
		handler = &Cloopen{
			AccountSid:   env.Getenv("SMS_CLOOPEN_SID", "aaf98f895069246a01506a9770ea0268"),
			AppID:        env.Getenv("SMS_CLOOPEN_APP_ID", "8a48b551506fd26f01509405471a6db8"),
			AccountToken: env.Getenv("SMS_CLOOPEN_ACCOUNT_TOKEN", "3fd8b18597d346c48631821abc00b138"),
		}
	default:
		return handler, fmt.Errorf("未找 %s SMS 驱动", s.Drive)
	}
	return handler, err
}
