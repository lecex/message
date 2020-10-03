package wechat

import (
	"github.com/bigrocs/wechat"
	pb "github.com/lecex/message/proto/message"
)

type Wechat struct {
	AppId       string
	Secret      string
	AccessToken string
}

type QueryParams struct {
	Mobile       string
	TemplateCode string
	Url          string
	Miniprogram  string `json:"miniprogram"`
	Data         string `json:"data"`
}

// NewClient 创建新的连接
func (srv *Wechat) NewClient() (client *wechat.Client) {
	client = wechat.NewClient()
	c := srv.Client.Config
	c.AppId = srv.AppId
	c.Secret = srv.Secret
	c.AccessToken = srv.AccessToken
	return client
}

func (srv *Wechat) Template(req *pb.Request, t *tpd.Template) (err error) {
	queryParams := QueryParams{}
	err = json.Unmarshal([]byte(req.QueryParams), &queryParams)
	if req.Addressee != "" {
		queryParams.Mobile = req.Addressee
	}
	if t.TemplateCode != "" {
		queryParams.TemplateCode = t.TemplateCode
	}
	request := requests.NewCommonRequest()
	request.Domain = "offiaccount"
	request.ApiName = "message.template"
	request.QueryParams = map[string]interface{}{
		"touser":      queryParams.Addressee,
		"template_id": queryParams.TemplateCode,
		"url":         queryParams.Url,
		"miniprogram": queryParams.Miniprogram,
		"data":        queryParams.Data,
	}
	return srv.request(request)
}
func (srv *Wechat) request(request *requests.CommonRequest) (req mxj.Map, err error) {
	client := srv.NewClient()
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return req, err
	}
	req, err = response.GetHttpContentMap()
	if err != nil {
		return req, err
	}
	return req, err
}
