package investgo

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/jstalex/tingo/proto"
	"github.com/jstalex/tingo/retry"
)

type OrdersStreamClient struct {
	conn     *grpc.ClientConn
	config   Config
	logger   Logger
	ctx      context.Context
	pbClient pb.OrdersStreamServiceClient
}

// TradesStream - Стрим сделок по запрашиваемым аккаунтам
func (o *OrdersStreamClient) TradesStream(accounts []string) (*TradesStream, error) {
	ctx, cancel := context.WithCancel(o.ctx)
	ts := &TradesStream{
		stream:       nil,
		ordersClient: o,
		trades:       make(chan *pb.OrderTrades),
		ctx:          ctx,
		cancel:       cancel,
	}
	stream, err := o.pbClient.TradesStream(ctx, &pb.TradesStreamRequest{
		Accounts: accounts,
	}, retry.WithOnRetryCallback(ts.restart))
	if err != nil {
		cancel()
		return nil, err
	}
	ts.stream = stream
	return ts, nil
}
