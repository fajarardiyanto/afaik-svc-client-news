package connection

import (
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/config"
	"github.com/fajarardiyanto/flt-go-listener/interfaces"
	"github.com/fajarardiyanto/flt-go-listener/lib/client"
	"google.golang.org/grpc"
)

type connection struct {
	conf interfaces.Client
}

func NewConnection(conf interfaces.Client) *connection {
	return &connection{conf}
}

func (c *connection) InitConn() *grpc.ClientConn {
	return client.NewListenerClient(config.GetLogger(), c.conf).Init()
}
