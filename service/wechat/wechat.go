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

// NewClient 创建新的连接
func (srv *Wechat) NewClient() (client *wechat.Client) {
	client = wechat.NewClient()
	c := srv.Client.Config
	c.AppId = srv.AppId
	c.Secret = srv.Secret
	c.AccessToken = srv.AccessToken
	return client
}

func (srv *Wechat) Template(req *pb.Request) (err error) {
	request := requests.NewCommonRequest()
	request.Domain = "offiaccount"
	request.ApiName = "message.template"
	request.QueryParams = map[string]interface{}{
		"touser":      req.Addressee,
		"template_id": "ybgOF-ZQsWTr8JS0lGwuRzFPdBKGAsiJiIk5ZX0EaDY",
		"url":         req.QueryParams["url"],
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": req.QueryParams["data.first"],
				"color": "#173177",
			},
			"keyword1": map[string]interface{}{
				"value": req.QueryParams["data.keyword1"],
				"color": "#173177",
			},
			"keyword2": map[string]interface{}{
				"value": req.QueryParams["data.keyword2"],
				"color": "#173177",
			},
			"keyword3": map[string]interface{}{
				"value": req.QueryParams["data.keyword3"],
				"color": "#173177",
			},
			"remark": map[string]interface{}{
				"value": req.QueryParams["data.remark"],
				"color": "#173177",
			},
		},
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
