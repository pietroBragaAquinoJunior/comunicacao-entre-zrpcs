// Code generated by goctl. DO NOT EDIT.
// Source: zrpc-um.proto

package zrpcumservice

import (
	"context"

	"zrpc-aula-2/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ZrpcUmMethodRequest  = __.ZrpcUmMethodRequest
	ZrpcUmMethodResponse = __.ZrpcUmMethodResponse

	ZrpcUmService interface {
		ZrpcUmMethod(ctx context.Context, in *ZrpcUmMethodRequest, opts ...grpc.CallOption) (*ZrpcUmMethodResponse, error)
	}

	defaultZrpcUmService struct {
		cli zrpc.Client
	}
)

func NewZrpcUmService(cli zrpc.Client) ZrpcUmService {
	return &defaultZrpcUmService{
		cli: cli,
	}
}

func (m *defaultZrpcUmService) ZrpcUmMethod(ctx context.Context, in *ZrpcUmMethodRequest, opts ...grpc.CallOption) (*ZrpcUmMethodResponse, error) {
	client := __.NewZrpcUmServiceClient(m.cli.Conn())
	return client.ZrpcUmMethod(ctx, in, opts...)
}
