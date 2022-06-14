package server

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type request struct {
	weight  int
	payload []byte
}

func initTransactionHandlers(h *mux.Router, s *Server, opts []httptransport.ServerOption) {
	createNewTransactionEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentRequest)
		return s.CreatePayment(ctx, r)
	}

	changeStatusTransactionEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentStatusChangeRequest)
		return s.ChangePaymentStatus(ctx, r)
	}

	checkStatusTransactionEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentStatusCheckRequest)
		return s.CheckPaymentStatus(ctx, r)
	}

	getAllPaymentsByIDEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentGetFromIDRequest)
		return s.GetAllPaymentsByID(ctx, r)
	}

	getAllPaymentsByEmailEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentGetFromEmailRequest)
		return s.GetAllPaymentsByEmail(ctx, r)
	}

	cancelTransactionEndpoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(*PaymentCancelRequest)
		return s.CancelTransaction(ctx, r)
	}

	h.Methods("POST").Path("/api/transaction/create-new-transaction").Handler(httptransport.NewServer(
		createNewTransactionEndpoint,
		decodeCreateNewTransactionRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("PUT").Path("/api/transaction/change-status-transaction/{id}&{transaction_id}&{secret_key}").Handler(httptransport.NewServer(
		changeStatusTransactionEndpoint,
		decodeChangeStatusTransactionRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("GET").Path("/api/transaction/check-status-transaction/{id}").Handler(httptransport.NewServer(
		checkStatusTransactionEndpoint,
		decodeCheckStatusTransactionRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("GET").Path("/api/transaction/get-payments-id/{id}").Handler(httptransport.NewServer(
		getAllPaymentsByIDEndpoint,
		decodeGetAllPaymentsByIDRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("GET").Path("/api/transaction/get-payments-email/{email}").Handler(httptransport.NewServer(
		getAllPaymentsByEmailEndpoint,
		decodeGetAllPaymentsByEmailRequest,
		newEncodeResponse(),
		opts...))

	h.Methods("DELETE").Path("/api/transaction/cancel-transaction").Handler(httptransport.NewServer(
		cancelTransactionEndpoint,
		decodeCancelTransactionRequest,
		newEncodeResponse(),
		opts...))

}

func decodeCreateNewTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	incomeReq, err := readRequestBody(r)
	if err != nil {
		log.Printf("failed read http request body: %s", err.Error())
		return &PaymentRequest{}, nil
	}

	req := PaymentRequest{}
	if err := json.Unmarshal(incomeReq.payload, &req); err != nil {
		return nil, errors.Wrap(err, "unmarshal request")
	}

	return &req, nil
}

func decodeChangeStatusTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	trid := mux.Vars(r)["transaction_id"]
	hash := mux.Vars(r)["secret_key"]
	idInt, err := strconv.ParseInt(trid, 10, 64)
	if err != nil {
		log.Printf("failed read http request body: %s", err.Error())
		return &PaymentStatusCheckRequest{}, nil
	}

	req := &PaymentStatusChangeRequest{
		UserID:          id,
		TransactionID:   idInt,
		TransactionHash: hash,
	}

	return &req, nil
}

func decodeCheckStatusTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("failed read http request body: %s", err.Error())
		return &PaymentStatusCheckRequest{}, nil
	}

	req := PaymentStatusCheckRequest{
		TransactionID: idInt,
	}

	return &req, nil
}

func decodeGetAllPaymentsByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]

	req := PaymentGetFromIDRequest{
		UserID: id,
	}

	return &req, nil
}

func decodeGetAllPaymentsByEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["email"]

	req := PaymentGetFromEmailRequest{
		Email: id,
	}

	return &req, nil
}

func decodeCancelTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	incomeReq, err := readRequestBody(r)
	if err != nil {
		log.Printf("failed read http request body: %s", err.Error())
		return &PaymentCancelRequest{}, nil
	}

	req := PaymentCancelRequest{}
	if err := json.Unmarshal(incomeReq.payload, &req); err != nil {
		return nil, errors.Wrap(err, "unmarshal request")
	}

	return &req, nil
}

func readRequestBody(r *http.Request) (*request, error) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read from body")
	}

	return &request{
		weight:  len(req),
		payload: req,
	}, nil
}
