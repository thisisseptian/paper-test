package handler

import (
	"context"
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

func TestUserDetail(t *testing.T) {
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
			name:         "error - get wallet by wallet id",
			expectedCode: http.StatusNotFound,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{}, errors.New("fail")).Once()
			},
		},
		{
			name:         "success",
			expectedCode: http.StatusOK,
			mocks: func() {
				mockStorage.On("GetUserByUserID", mock.Anything).Return(model.User{UserID: 1, Username: "Septian", WalletID: 1}, nil).Once()
				mockStorage.On("GetWalletByWalletID", mock.Anything).Return(model.Wallet{WalletID: 1, Balance: 10000000}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()

			r := httptest.NewRequest(http.MethodGet, "/user_detail", nil)
			w := httptest.NewRecorder()

			if tt.name != "error - get context" {
				userID := int64(1)
				ctx := context.WithValue(r.Context(), constant.CtxUserIDKey, userID)
				r = r.WithContext(ctx)
			}

			// main func
			mockHandler.UserDetail(w, r)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockStorage.AssertExpectations(t)
		})
	}
}
