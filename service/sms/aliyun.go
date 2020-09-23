package sms

import (
	// 公共引入
	"encoding/json"
	"errors"

	"github.com/micro/go-micro/util/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"

	pb "github.com/lecex/message/proto/message"
	tpd "github.com/lecex/message/proto/template"
)

// Aliyun 阿里云驱动
type Aliyun struct {
	RegionID        string
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
}

// Send 获取所有消息事件信息
func (srv *Aliyun) Send(req *pb.Request, t *tpd.Template) (valid bool, err error) {
	// 创建连接
	client, err := srv.client()
	if err != nil {
		log.Log(err)
		return false, err
	}
	// 请求参数构建
	request, err := srv.request(req, t)
	if err != nil {
		log.Log(err)
		return false, err
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)

	if err != nil {
		log.Log(err)
		return false, err
	}

	// 返回数据处理
	valid, err = srv.response(response)
	if err != nil {
		log.Log(err)
	}
	return valid, err
}

// client 创建阿里云连接
func (srv *Aliyun) client() (client *sdk.Client, err error) {
	// 创建连接
	return sdk.NewClientWithAccessKey(
		srv.RegionID,
		srv.AccessKeyID,
		srv.AccessKeySecret,
	)
}

// request 构建阿里云请求参数
func (srv *Aliyun) request(req *pb.Request, t *tpd.Template) (request *requests.CommonRequest, err error) {
	// 配置参数
	request = requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["PhoneNumbers"] = req.Addressee
	request.QueryParams["SignName"] = srv.SignName
	request.QueryParams["TemplateCode"] = t.TemplateCode
	queryParams, err := json.Marshal(req.QueryParams)
	request.QueryParams["TemplateParam"] = string(queryParams)
	return
}

// response 返回数据处理
func (srv *Aliyun) response(response *responses.CommonResponse) (valid bool, err error) {
	// res 返回请求
	res := map[string]string{}
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &res)
	if err != nil {
		return false, err
	}
	if res["Code"] != "OK" {
		return false, errors.New(res["Message"])
	}
	return true, err
}
