package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BlackRRR/payment-emulator/internal/repository/transaction"
	"github.com/BlackRRR/payment-emulator/test/test_model"
	"io"
	"net/http"
	"testing"
	"time"
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

	c := NewClient("http://localhost:8000/api/transaction/create-new-transaction")
	res, err := c.PostRequest(r)
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
		"Payload: id = ", paymentResponse.Payload.TransactionID,
		" hash = ", paymentResponse.Payload.TransactionHash,
		" error ", paymentResponse.Error)
}

func TestChangeStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusChangeResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	c := NewClient("http://localhost:8000/api/transaction/change-status-transaction/" + "15" + "&" + "38" + "&" + "SCsm3y5")

	res, err := c.PutRequest()
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
		"Payload: status = ", paymentResponse.Payload.TransactionStatus,
		" error ", paymentResponse.Error)
}

func TestCheckPaymentStatusTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentStatusCheckResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	c := NewClient("http://localhost:8000/api/transaction/check-status-transaction/" + "3911412")
	res, err := c.GetRequest()
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
		"Payload: status = ", paymentResponse.Payload.TransactionStatus,
		" error ", paymentResponse.Error)
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

	c := NewClient("http://localhost:8000/api/transaction/get-payments-id/" + "15")
	res, err := c.GetRequest()
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

	fmt.Println(" error ", paymentResponse.Error)
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

	c := NewClient("http://localhost:8000/api/transaction/get-payments-email/" + "15")
	res, err := c.GetRequest()
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

	fmt.Println(" error ", paymentResponse.Error)

}

func TestCancelTransactionService(t *testing.T) {
	paymentResponse := test_model.PaymentCancelResponse{
		Result:  0,
		Payload: &test_model.TransactionStatusPayload{},
		Error:   nil,
	}

	c := NewClient("http://localhost:8000/api/transaction/cancel-transaction/" + "3911412")

	res, err := c.DeleteRequest()
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
		"Payload: status = ", paymentResponse.Payload.TransactionStatus,
		" error ", paymentResponse.Error)
}

func (c Client) PostRequest(r io.Reader) (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, c.url, r)
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

func (c Client) GetRequest() (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, c.url, nil)
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

func (c Client) PutRequest() (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, c.url, nil)
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

func (c Client) DeleteRequest() (*json.Decoder, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodDelete, c.url, nil)
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
