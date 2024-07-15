package handler

import (
	"log"
	"net/http"

	"paper-test/constant"
	"paper-test/model"
)

// UserDetail is handler to get user detail
func (h *Handler) UserDetail(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.CtxUserIDKey).(int64)
	if !ok {
		log.Println("[UserDetail] error get user id from context")
		h.RenderResponse(w, r, "", http.StatusInternalServerError, "internal server error")
		return
	}

	user, err := h.Storage.GetUserByUserID(userID)
	if err != nil {
		log.Printf("[UserDetail][UserID %d] failed get user by user id with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	wallet, err := h.Storage.GetWalletByWalletID(user.WalletID)
	if err != nil {
		log.Printf("[UserDetail][UserID %d] failed get wallet by wallet id with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	response := model.UserDetail{
		User:   user,
		Wallet: wallet,
	}

	h.RenderResponse(w, r, response, http.StatusOK, "")
}
