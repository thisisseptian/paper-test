package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"paper-test/constant"
)

// Auth is middleware handler to validate token
func (h *Handler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			h.RenderResponse(w, r, "", http.StatusUnauthorized, "unauthorized")
			return
		}

		user, err := h.Storage.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			h.RenderResponse(w, r, "", http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), constant.CtxUserIDKey, user.UserID)
		next(w, r.WithContext(ctx))
	}
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// RenderResponse is handler to generate response
func (h *Handler) RenderResponse(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int, msg string) {
	response := Response{
		Code:    statusCode,
		Data:    data,
		Message: msg,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("[RenderResponse] error marshalling JSON:", err)
		http.Error(w, "error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
