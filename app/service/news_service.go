package service

import (
	"context"
	"encoding/json"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/entity"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/repository/models"
	"github.com/fajarardiyanto/afaik-svc-client-news/internal/repository/services"
	logger "github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-tracer/lib/jaeger"
	"github.com/fajarardiyanto/flt-go-utils/response"
	"net/http"
	"sync"
	"time"
)

type NewsClient struct {
	log   logger.Logger
	model models.NewsRepository
	sync.RWMutex
}

func NewNewsClient(log logger.Logger, model models.NewsRepository) services.NewsClientRepo {
	return &NewsClient{
		log:   log,
		model: model,
	}
}

func (n *NewsClient) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	span, _ := jaeger.CreateRootSpan(ctx, "Get Handler News")
	defer span.Finish()

	//for k, v := range r.Header {
	//	fmt.Println("dsadasd", k, v)
	//}
	//span.LogFields(log.Object("header", parser.Stringify(w.Header())))

	decoder := json.NewDecoder(r.Body)
	var req entity.GetNews

	if err := decoder.Decode(&req); err != nil {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Parsing Error failed "+err.Error())
		return
	}

	results, err := n.model.GetNews(ctx, req, span)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Parsing Error failed "+err.Error())
		return
	}

	response.WriteResponse(w, http.StatusOK, results)
}
