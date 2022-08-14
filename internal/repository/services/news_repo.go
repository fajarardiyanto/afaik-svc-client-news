package services

import (
	"net/http"
)

type NewsClientRepo interface {
	Get(w http.ResponseWriter, r *http.Request)
}
