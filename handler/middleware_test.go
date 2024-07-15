package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"paper-test/model"
	"paper-test/storage/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth(t *testing.T) {
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
			name:         "error - unauthorized",
			expectedCode: http.StatusUnauthorized,
			mocks:        func() {},
		},
		{
			name:         "error - user not found",
			expectedCode: http.StatusUnauthorized,
			mocks: func() {
				mockStorage.On("GetUserByUsernameAndPassword", mock.Anything, mock.Anything).Return(model.User{}, errors.New("fail")).Once()
			},
		},
		{
			name:         "success",
			expectedCode: http.StatusOK,
			mocks: func() {
				mockStorage.On("GetUserByUsernameAndPassword", mock.Anything, mock.Anything).Return(model.User{UserID: 1, Username: "Septian", Password: []byte{}, WalletID: 1}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			if tt.name != "error - unauthorized" {
				r.SetBasicAuth("testuser", "testpass")
			}

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// main func
			authHandler := mockHandler.Auth(nextHandler)
			authHandler.ServeHTTP(w, r)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestRenderResponse(t *testing.T) {
	mockHandler := &Handler{} // Assuming no need for mocks in RenderResponse

	tests := []struct {
		name         string
		data         interface{}
		statusCode   int
		expectedCode int
	}{
		{
			name:         "success",
			data:         "test data",
			statusCode:   http.StatusOK,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			w := httptest.NewRecorder()

			// main func
			mockHandler.RenderResponse(w, r, tt.data, tt.statusCode, "")

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
