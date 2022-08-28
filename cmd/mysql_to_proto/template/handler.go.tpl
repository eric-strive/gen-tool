package api

import (
	"context"
	"gitlab.intsig.net/cs-server2/services/core_lib/app_proto/app_recharge/proto/recharge"
    "gitlab.intsig.net/cs-server2/services/core_lib/service"
    "gitlab.intsig.net/cs-server2/services/purchase/app_recharge/bis"
    "gitlab.intsig.net/cs-server2/services/purchase/app_recharge/config"
    "gitlab.intsig.net/cs-server2/services/purchase/app_recharge/errno"
)

type {{.ModelName}} struct{}

// 生成curd代码
func (s *{{.ModelName}}) Create(ctx context.Context, req *proto.{{.ModelName}}Create, response *proto.{{.ModelName}}Response) (error) {
	err := ins.BisRegister.Recharge.Create(ctx,req)
	rsp.Code, rsp.Msg = errno.ErrHandle(err)
	return nil
}


func (s *{{.ModelName}}) Delete(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Delete{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) DeleteById(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.DeleteById{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Update(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Update{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Find(ctx context.Context, req *proto.{{.ModelName}}) (*proto.Response, error) {
	data, err := service.Find{{.ModelName}}(req)
	return response.SuccessAny(data), err
}

func (s *{{.ModelName}}) Lists(ctx context.Context, req *proto.Request) (*proto.Responses, error) {
    data, total, err := service.GetList{{.ModelName}}(req)

	var any = make([]*anypb.Any, len(data))
	for k, r := range data {
		any[k], err = ptypes.MarshalAny(r)
	}

	return response.SuccesssAny(any, total), err
}


