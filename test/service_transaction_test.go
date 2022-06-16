package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BlackRRR/payment-emulator/internal/model"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/BlackRRR/payment-emulator/test/test_model"
	"io"
	"net/http"
	"strconv"
	"testing"
)

const (
	userID          = 8040554
	transactionID   = 4433257
	transactionHash = "hai0gjO"
	amount          = 5555
	email           = "bl@er.tu"
	currency        = "RUB"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func TestCreatePaymentTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentResponse{
		Result:  0,
		Payload: &test_model.TransactionHashPayload{},
		Error:   &model.ServerError{},
	}

	marshal, err := json.Marshal(&test_model.PaymentRequest{
		UserID:   userID,
		Email:    email,
		Amount:   amount,
		Currency: currency,
	})
	if err != nil {
		t.Errorf("failed marshal payment request %v", err)
	}

	r := bytes.NewReader(marshal)

	c := NewClient("http://localhost:8000/api/transaction/create-new-transaction")
	res, err := c.NewRequest(r, http.MethodPost)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result: ", paymentResponse.Result,
		"Payload: id =", paymentResponse.Payload.TransactionID,
		" hash =", paymentResponse.Payload.TransactionHash,
		" error =", paymentResponse.Error)
}

func TestChangeStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusChangeResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   &model.ServerError{},
	}

	uID := strconv.FormatInt(userID, 10)
	tID := strconv.FormatInt(transactionID, 10)

	c := NewClient("http://localhost:8000/api/transaction/change-status-transaction/" + uID + "&" + tID + "&" + transactionHash)

	res, err := c.NewRequest(nil, http.MethodPut)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result:", paymentResponse.Result,
		"Payload: status =", paymentResponse.Payload.TransactionStatus,
		" error =", paymentResponse.Error)
}

func TestCheckPaymentStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusCheckResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	tID := strconv.FormatInt(transactionID, 10)

	c := NewClient("http://localhost:8000/api/transaction/check-status-transaction/" + tID)
	res, err := c.NewRequest(nil, http.MethodGet)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result:", paymentResponse.Result,
		"Payload: status =", paymentResponse.Payload.TransactionStatus,
		" error =", paymentResponse.Error)
}

func TestGetAllPaymentsByIDTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentGetFromIDResponse{
		Result:  0,
		Payload: &test_model.Payments{Payments: []*transaction.Payment{{}}},
		Error:   &model.ServerError{},
	}

	uID := strconv.FormatInt(userID, 10)

	c := NewClient("http://localhost:8000/api/transaction/get-payments-id/" + uID)
	res, err := c.NewRequest(nil, http.MethodGet)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result: ", paymentResponse.Result)

	for _, value := range paymentResponse.Payload.Payments {
		fmt.Print(value, " ")
	}

	fmt.Println(" error =", paymentResponse.Error)
}

func TestGetAllPaymentsByEmailTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentGetFromEmailResponse{
		Result:  0,
		Payload: &test_model.Payments{Payments: []*transaction.Payment{{}}},
		Error:   &model.ServerError{},
	}

	c := NewClient("http://localhost:8000/api/transaction/get-payments-email/" + email)
	res, err := c.NewRequest(nil, http.MethodGet)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result:", paymentResponse.Result)

	for _, value := range paymentResponse.Payload.Payments {
		fmt.Print(value, " ")
	}

	fmt.Println(" error =", paymentResponse.Error)

}

func TestCancelTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentCancelResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   &model.ServerError{},
	}

	tID := strconv.FormatInt(transactionID, 10)

	c := NewClient("http://localhost:8000/api/transaction/cancel-transaction/" + tID)

	res, err := c.NewRequest(nil, http.MethodDelete)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = res.Decode(&paymentResponse)
	if err != nil {
		t.Errorf("failed to decode response %v", err)
	}

	if paymentResponse.Payload == nil {
		t.Errorf("expected request payload to be not nil got %v", err)
	}

	fmt.Println("Result: ", paymentResponse.Result,
		"Payload: status =", paymentResponse.Payload.TransactionStatus,
		" error =", paymentResponse.Error)
}

func (c Client) NewRequest(r io.Reader, method string) (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, c.url, r)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	decode := json.NewDecoder(resp.Body)

	return decode, nil
}
