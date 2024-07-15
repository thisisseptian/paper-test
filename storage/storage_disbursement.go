package storage

import (
	"sort"

	"paper-test/model"
)

var (
	disbursementIDCounter int64
)

func incrementDisbursementID() int64 {
	disbursementIDCounter++
	return disbursementIDCounter
}

func (s *Storage) InsertDisbursement(disbursement model.Disbursement) (int64, error) {
	disbursementID := incrementDisbursementID()
	disbursement.DisbursementID = disbursementID
	s.disbursements[disbursementID] = &disbursement

	return disbursementID, nil
}

func (s *Storage) GetDisbursements(walletID int64) ([]model.Disbursement, error) {
	disbursementList := []model.Disbursement{}

	for _, disbursement := range s.disbursements {
		if disbursement.FromWalletID == walletID {
			disbursementList = append(disbursementList, *disbursement)
		}
	}

	sort.Slice(disbursementList, func(i, j int) bool {
		return disbursementList[i].DisbursementID < disbursementList[j].DisbursementID
	})

	return disbursementList, nil
}
