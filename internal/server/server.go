package server

import (
	"context"
	"encoding/json"
	"github.com/BlackRRR/payment-emulator/internal/services"
	"github.com/BlackRRR/payment-emulator/internal/services/transaction"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	trans *transaction.TransactionService
}

func InitServer(services *services.Services) *Server {
	server := &Server{
		trans: services.Trans,
	}

	return server
}

func MakeHTTPHandler(s *Server) http.Handler {
	h := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(NewEncodeError()),
	}

	initTransactionHandlers(h, s, opts)

	return h
}

func newEncodeResponse() httptransport.EncodeResponseFunc {
	return func(_ context.Context, w http.ResponseWriter, response interface{}) error {
		w.WriteHeader(http.StatusOK)

		return json.NewEncoder(w).Encode(response)
	}
}

func NewEncodeError() httptransport.ErrorEncoder {
	return func(_ context.Context, err error, w http.ResponseWriter) {

		w.WriteHeader(http.StatusInternalServerError)

		_ = json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
}
