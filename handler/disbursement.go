package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"paper-test/constant"
	"paper-test/helper"
	"paper-test/model"
)

// DisburseList is handler to get list of user disbursement
func (h *Handler) DisburseList(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.CtxUserIDKey).(int64)
	if !ok {
		log.Println("[DisburseList] error get user id from context")
		h.RenderResponse(w, r, "", http.StatusInternalServerError, "internal server error")
		return
	}

	user, err := h.Storage.GetUserByUserID(userID)
	if err != nil {
		log.Printf("[DisburseList][UserID %d] failed get user by user id with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	disbursements, err := h.Storage.GetDisbursements(user.WalletID)
	if err != nil {
		log.Printf("[DisburseList][UserID %d] failed get users with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	h.RenderResponse(w, r, disbursements, http.StatusOK, "")
}

// Disburse is handler to disburse balance from wallet to bank account
func (h *Handler) Disburse(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constant.CtxUserIDKey).(int64)
	if !ok {
		log.Println("[Disburse] error get user id from context")
		h.RenderResponse(w, r, "", http.StatusInternalServerError, "internal server error")
		return
	}

	var payload model.DisbursementRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("[Disburse][UserID %d] failed decode body with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "bad Request")
		return
	}

	if payload.Amount == 0 {
		log.Printf("[Disburse][UserID %d] requested amount cannot be 0", userID)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "requested amount cannot be 0")
		return
	}

	if payload.BankName == "" {
		log.Printf("[Disburse][UserID %d] requested bank name cannot be empty", userID)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "requested bank name cannot be empty")
		return
	}

	if !helper.IsValidBankName(payload.BankName) {
		log.Printf("[Disburse][UserID %d] bank name invalid", userID)
		h.RenderResponse(w, r, "", http.StatusBadRequest, `bank name invalid, list of bank name ("BCA", "BNI", "Mandiri", "OCBC", "Maybank", "Panin", "BRI", "JAGO", "BTPN",)`)
		return
	}

	if payload.BankAccountName == "" {
		log.Printf("[Disburse][UserID %d] requested bank account name cannot be empty", userID)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "requested bank account name cannot be empty")
		return
	}

	if payload.BankAccountNumber == "" {
		log.Printf("[Disburse][UserID %d] requested bank account number cannot be empty", userID)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "requested bank account number cannot be empty")
		return
	}

	user, err := h.Storage.GetUserByUserID(userID)
	if err != nil {
		log.Printf("[Disburse][UserID %d] failed get user by user id with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	wallet, err := h.Storage.GetWalletByWalletID(user.WalletID)
	if err != nil {
		log.Printf("[Disburse][UserID %d] failed get wallet by wallet id with error %+v", userID, err)
		h.RenderResponse(w, r, "", http.StatusNotFound, err.Error())
		return
	}

	if wallet.Balance < payload.Amount {
		log.Printf("[Disburse][UserID %d] wallet balance (%.2f) is lower than requested amount (%.2f)", userID, wallet.Balance, payload.Amount)
		h.RenderResponse(w, r, "", http.StatusBadRequest, "insufficient balance")
		return
	}

	wallet.Balance -= payload.Amount
	timeNow := time.Now()
	disbursement := model.Disbursement{
		FromWalletID: wallet.WalletID,
		ToBankAccount: model.BankAccount{
			BankName:          payload.BankName,
			BankAccountName:   payload.BankAccountName,
			BankAccountNumber: payload.BankAccountNumber,
		},
		Amount:    payload.Amount,
		TimeStamp: timeNow,
	}

	_, err = h.Storage.InsertDisbursement(disbursement)
	if err != nil {
		log.Printf("[Disburse][UserID %d] failed insert disbursement data with error %+v", userID, err.Error())
		h.RenderResponse(w, r, "", http.StatusInternalServerError, "internal server error")
		return
	}

	err = h.Storage.UpdateWalletBalanceByWalletID(wallet)
	if err != nil {
		log.Printf("[Disburse][UserID %d] failed update wallet balance data with error %+v", userID, err.Error())
		h.RenderResponse(w, r, "", http.StatusInternalServerError, "internal server error")
		return
	}

	h.RenderResponse(w, r, "", http.StatusOK, fmt.Sprintf("disbursement complete, current balance %2.f", wallet.Balance))
}
