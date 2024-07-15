package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"paper-test/constant"
	"paper-test/model"
	"paper-test/storage/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDisburseList(t *testing.T) {
	mockStorage := new(mocks.IStorage)
	mockHandler := &Handler{
		Storage: mockStorage,
	}

	tests := []struct {
		name         string
		expectedCode int
		mocks        func()
	}{
		{
			name:         "error - get context",
			expectedCode: http.StatusInternalServerError,
			mocks:        func() {},
		},
		{
			name:         "error - get user by user id",
			expectedCode: http.StatusNotFound,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{}, errors.New("fail")).Once()
			},
		},
		{
			name:         "error - get disbursements",
			expectedCode: http.StatusNotFound,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetDisbursements", mock.Anything).Return([]model.Disbursement{}, errors.New("fail")).Once()
			},
		},
		{
			name:         "success",
			expectedCode: http.StatusOK,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetDisbursements", mock.Anything).Return([]model.Disbursement{{DisbursementID: 1, FromWalletID: 1}}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()

			r := httptest.NewRequest(http.MethodGet, "/disburse_list", nil)
			w := httptest.NewRecorder()

			if tt.name != "error - get context" {
				userID := int64(1)
				ctx := context.WithValue(r.Context(), constant.CtxUserIDKey, userID)
				r = r.WithContext(ctx)
			}

			// main func
			mockHandler.DisburseList(w, r)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockStorage.AssertExpectations(t)
		})
	}
}

func TestDisburse(t *testing.T) {
	mockStorage := new(mocks.IStorage)
	mockHandler := &Handler{
		Storage: mockStorage,
	}

	tests := []struct {
		name         string
		requestBody  interface{}
		expectedCode int
		mocks        func()
	}{
		{
			name:         "error - get context",
			requestBody:  "",
			expectedCode: http.StatusInternalServerError,
			mocks:        func() {},
		},
		{
			name:         "error - get request body",
			requestBody:  "invalid body",
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - payload amount zero",
			requestBody: model.DisbursementRequest{
				Amount: 0,
			},
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - payload bank name empty",
			requestBody: model.DisbursementRequest{
				Amount:   500000,
				BankName: "",
			},
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - payload bank name invalid",
			requestBody: model.DisbursementRequest{
				Amount:   500000,
				BankName: "Bang salah",
			},
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - bank account name empty",
			requestBody: model.DisbursementRequest{
				Amount:          500000,
				BankName:        "bca",
				BankAccountName: "",
			},
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - bank account number empty",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "",
			},
			expectedCode: http.StatusBadRequest,
			mocks:        func() {},
		},
		{
			name: "error - get user by user id",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusNotFound,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{}, errors.New("fail")).Once()
			},
		},
		{
			name: "error - get wallet by wallet id",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusNotFound,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{}, errors.New("fail")).Once()
			},
		},
		{
			name: "error - insufficient balance",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusBadRequest,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{WalletID: 1, Balance: 100000}, nil).Once()
			},
		},
		{
			name: "error - insert disbursement",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusInternalServerError,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{WalletID: 1, Balance: 1000000}, nil).Once()
				mockStorage.On("InsertDisbursement", mock.Anything).Return(int64(1), errors.New("fail")).Once()
			},
		},
		{
			name: "error - update wallet balance by wallet id",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusInternalServerError,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{WalletID: 1, Balance: 1000000}, nil).Once()
				mockStorage.On("InsertDisbursement", mock.Anything).Return(int64(1), nil).Once()
				mockStorage.On("UpdateWalletBalanceByWalletID", mock.Anything).Return(errors.New("fail")).Once()
			},
		},
		{
			name: "success",
			requestBody: model.DisbursementRequest{
				Amount:            500000,
				BankName:          "bca",
				BankAccountName:   "Septian",
				BankAccountNumber: "12345678",
			},
			expectedCode: http.StatusOK,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{WalletID: 1, Balance: 1000000}, nil).Once()
				mockStorage.On("InsertDisbursement", mock.Anything).Return(int64(1), nil).Once()
				mockStorage.On("UpdateWalletBalanceByWalletID", mock.Anything).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()

			var requestBody bytes.Buffer
			if err := json.NewEncoder(&requestBody).Encode(tt.requestBody); err != nil {
				t.Fatalf("could not encode request body: %+v", err)
			}

			r := httptest.NewRequest(http.MethodPost, "/disburse", &requestBody)
			w := httptest.NewRecorder()

			if tt.name != "error - get context" {
				userID := int64(1)
				ctx := context.WithValue(r.Context(), constant.CtxUserIDKey, userID)
				r = r.WithContext(ctx)
			}

			// main func
			mockHandler.Disburse(w, r)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockStorage.AssertExpectations(t)
		})
	}
}
