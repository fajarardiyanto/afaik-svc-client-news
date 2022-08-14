package models

import (
	"context"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/entity"
	pb "github.com/fajarardiyanto/module-proto/go/services/news"
	"github.com/opentracing/opentracing-go"
)

type NewsRepository interface {
	GetNews(context.Context, entity.GetNews, opentracing.Span) (*pb.NewsResponse, error)
}
