package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/johnson7543/pricefetcher/service"
	"github.com/johnson7543/pricefetcher/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        service.PriceService
}

type contextKey string

func NewJSONAPIServer(listenAddr string, svc service.PriceService) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	const uuidKey contextKey = "uuid"
	ctx := context.Background()
	ctx = context.WithValue(ctx, uuidKey, uuid.New().String())

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			uuid, _ := ctx.Value(uuidKey).(string)
			timestamp := time.Now().Format(time.RFC3339Nano)

			Resp := types.TemplateResponse{
				UUID:       uuid,
				Message:    err.Error(),
				Timestamp:  timestamp,
				StatusCode: http.StatusBadRequest,
			}

			writeJSON(w, http.StatusBadRequest, &Resp)
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	const uuidKey contextKey = "uuid"
	uuid, _ := ctx.Value(uuidKey).(string)

	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	timestamp := time.Now().Format(time.RFC3339Nano)

	Resp := types.TemplateResponse{
		UUID:       uuid,
		Data:       priceResp,
		Timestamp:  timestamp,
		StatusCode: http.StatusOK,
	}

	return writeJSON(w, http.StatusOK, &Resp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
