package models

import (
	"context"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/entity"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/repository/models"
	tr "github.com/fajarardiyanto/flt-go-tracer/lib/jaeger"
	"github.com/fajarardiyanto/flt-go-utils/pagination"
	"github.com/fajarardiyanto/flt-go-utils/parser"
	pb "github.com/fajarardiyanto/module-proto/go/services/news"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
)

type news struct {
	client pb.NewsServiceClient
	sync.RWMutex
}

func NewNewsRepository(c *grpc.ClientConn) models.NewsRepository {
	return &news{
		client: pb.NewNewsServiceClient(c),
	}
}

func (c *news) GetNews(ctx context.Context, req entity.GetNews, span opentracing.Span) (*pb.NewsResponse, error) {
	sp := tr.CreateSubChildSpan(span, "Get News Service Rest")
	defer sp.Finish()

	tr.LogRequest(sp, req)

	var wg sync.WaitGroup

	var errs []error
	total := make(chan *pb.TotalNews)
	result := make(chan *pb.NewsResponse)

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		res, err := c.client.GetNews(ctx, &pb.GetNewsRequest{Limit: req.Limit, Offset: req.Offset})
		if err != nil {
			c.Lock()
			errs = append(errs, err)
			c.Unlock()
		}

		c.Lock()
		result <- res
		c.Unlock()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		t, err := c.client.GetTotalNews(ctx, &emptypb.Empty{})
		if err != nil {
			c.Lock()
			errs = append(errs, err)
			c.Unlock()
		}

		c.Lock()
		total <- t
		c.Unlock()
	}(&wg)

	if len(errs) != 0 {
		return nil, errs[0]
	}

	res := <-result
	totalData := <-total
	results := &pb.NewsResponse{
		News:  res.News,
		Total: pagination.TotalPage(parser.StrToInt64(totalData.Total), req.Limit),
		Page:  req.Offset,
	}

	tr.LogRequest(sp, results)

	return results, nil
}
