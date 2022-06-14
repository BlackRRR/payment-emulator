package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/BlackRRR/payment-emulator/test/test_model"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

var Transaction *TransactionEntities

type TransactionEntities struct {
	TransactionID   int64
	TransactionHash string
}

func TestCreatePaymentTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentResponse{
		Result:  0,
		Payload: &test_model.TransactionHashPayload{},
		Error:   nil,
	}

	marshal, err := json.Marshal(&test_model.PaymentRequest{
		UserID:   "15",
		Email:    "15",
		Amount:   "15",
		Currency: "rub",
	})
	if err != nil {
		t.Errorf("failed marshal payment request %v", err)
	}

	r := bytes.NewReader(marshal)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/create-new-transaction")
	res, err := c.PostRequest(r)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload.TransactionHash == "" {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	ts := &TransactionEntities{
		TransactionID:   paymentResponse.Payload.TransactionID,
		TransactionHash: paymentResponse.Payload.TransactionHash,
	}

	Transaction = ts

	fmt.Println(string(res))
}

func TestChangeStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusChangeResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/change-status-transaction/" + "15" + "?" + "38" + "?" + "SCsm3y5")

	res, err := c.PutRequest(svr)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println(string(res))
}

func TestCheckPaymentStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusCheckResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/check-status-transaction/" + "3911412")
	res, err := c.GetRequest(handler)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println(string(res))
}

func TestGetAllPaymentsByIDTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentGetFromIDResponse{
		Result: 0,
		Payload: &test_model.Payments{Payments: []*transaction.Transaction{{
			TransactionID:    0,
			UserID:           "",
			Email:            "",
			Amount:           "",
			Currency:         "",
			DateOfCreation:   time.Time{},
			DateOfLastChange: time.Time{},
			Status:           "",
		},
		}},
		Error: nil,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/get-payments-id/" + "15")
	res, err := c.GetRequest(handler)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println(string(res))
}

func TestGetAllPaymentsByEmailTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentGetFromEmailResponse{
		Result: 0,
		Payload: &test_model.Payments{Payments: []*transaction.Transaction{{
			TransactionID:    0,
			UserID:           "",
			Email:            "",
			Amount:           "",
			Currency:         "",
			DateOfCreation:   time.Time{},
			DateOfLastChange: time.Time{},
			Status:           "",
		},
		}},
		Error: nil,
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/get-payments-email/" + "15")
	res, err := c.GetRequest(handler)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println(string(res))
}

func TestCancelTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentCancelResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	marshal, err := json.Marshal(&test_model.PaymentCancelRequest{
		TransactionID: Transaction.TransactionID,
	})

	if err != nil {
		t.Errorf("failed marshal payment request %v", err)
	}

	r := bytes.NewReader(marshal)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer svr.Close()

	c := NewClient("http://localhost:8000/api/transaction/cancel-transaction")

	res, err := c.DeleteRequest(r)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.Unmarshal(res, &paymentResponse)
	if err != nil {
		t.Errorf("unable to unmarshal got %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println(string(res))
}

func (c Client) PostRequest(r io.Reader) ([]byte, error) {
	res, err := http.Post(c.url, "application/json", r)
	if err != nil {
		return nil, errors.Wrap(err, "unable to complete POST request")
	}

	res.Close = true

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response data")
	}

	return out, nil
}

func (c Client) GetRequest(handler http.HandlerFunc) ([]byte, error) {
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, c.url, nil)
	handler.ServeHTTP(rr, request)

	out, err := io.ReadAll(rr.Body)
	if err != nil {
		return nil, err
	}

	//res, err := http.Get(c.url)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to complete POST request")
	//}
	//
	//res.Close = true
	//
	//defer res.Body.Close()
	//out, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to read response data")
	//}

	return out, nil
}

func (c Client) PutRequest(server *httptest.Server) ([]byte, error) {
	rr := httptest.NewRecorder()

	res, err := http.NewRequest(http.MethodPut, c.url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to complete POST request")
	}

	server.Config.Handler.ServeHTTP(rr, res)

	check := rr.Body.String()

	fmt.Println(check)

	out, err := io.ReadAll(rr.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c Client) DeleteRequest(r io.Reader) ([]byte, error) {
	res, err := http.NewRequest(http.MethodDelete, c.url, r)
	if err != nil {
		return nil, errors.Wrap(err, "unable to complete POST request")
	}

	res.Close = true

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response data")
	}

	return out, nil
}
