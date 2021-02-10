package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/bigrocs/wechat"
	"github.com/bigrocs/wechat/requests"
	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
	"github.com/micro/go-micro/v2/util/log"
)

type Wechat struct {
	AppId       string
	Secret      string
	AccessToken string
}

type QueryParams struct {
	Addressee    string `json:"addressee"`
	TemplateCode string `json:"template_code"`
	Url          string `json:"url"`
	Miniprogram  string `json:"miniprogram"`
	Data         string `json:"data"`
}

// NewClient 创建新的连接
func (srv *Wechat) NewClient() (client *wechat.Client) {
	client = wechat.NewClient()
	c := client.Config
	c.AppId = srv.AppId
	c.Secret = srv.Secret
	c.AccessToken = srv.AccessToken
	return client
}

func (srv *Wechat) Template(req *pb.Request, t *tpd.Template) (valid bool, err error) {
	queryParams := QueryParams{}
	err = json.Unmarshal([]byte(req.QueryParams), &queryParams)
	if req.Addressee != "" {
		queryParams.Addressee = req.Addressee
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
	valid, err = srv.request(request)
	if err != nil {
		log.Log(err)
	}
	return valid, err
}
func (srv *Wechat) request(request *requests.CommonRequest) (valid bool, err error) {
	client := srv.NewClient()
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return false, err
	}
	req, err := response.GetHttpContentMap()
	fmt.Println(req) //debug
	if err != nil {
		return false, err
	}
	return true, err
}
