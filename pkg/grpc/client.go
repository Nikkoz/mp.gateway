package grpc

import (
	"context"
	"fmt"
	innerContext "github.com/Nikkoz/mp.gateway/pkg/types/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	appNameHeader    = "x-app-name"
	appVersionHeader = "x-app-version"
)

type Client struct {
	Host       string
	Port       uint16
	Ctx        innerContext.Context
	AppName    string
	AppVersion string
}

func New(ctx innerContext.Context, host string, port uint16, appName, appVersion string) *Client {
	return &Client{
		Host:       host,
		Port:       port,
		Ctx:        ctx,
		AppName:    appName,
		AppVersion: appVersion,
	}
}

func (client *Client) GetConnection() (*grpc.ClientConn, error) {
	return grpc.DialContext(
		client.Ctx,
		fmt.Sprintf("%s:%d", client.Host, client.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(client.addAppInfoUnary),
	)
}

// AddAppInfoUnary добавляет в единичные запросы информацию о клиенте.
func (client *Client) addAppInfoUnary(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, appNameHeader, client.AppName)
	ctx = metadata.AppendToOutgoingContext(ctx, appVersionHeader, client.AppVersion)

	return invoker(ctx, method, req, reply, cc, opts...)
}
