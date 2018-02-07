package server

import (
	"encoding/json"

	"github.com/stellar/go/services/bifrost/sse"
)

func (s *Server) onStellarAccountCreated(destination string) {
	association, err := s.Database.GetAssociationByStellarPublicKey(destination)
	if err != nil {
		s.log.WithField("err", err).Error("Error getting association")
		return
	}

	if association == nil {
		s.log.WithField("stellarPublicKey", destination).Error("Association not found")
		return
	}

	s.SSEServer.BroadcastEvent(association.Address, sse.AccountCreatedAddressEvent, nil)
}

func (s *Server) onStellarAccountCredited(destination, assetCode, amount string) {
	association, err := s.Database.GetAssociationByStellarPublicKey(destination)
	if err != nil {
		s.log.WithField("err", err).Error("Error getting association")
		return
	}

	if association == nil {
		s.log.WithField("stellarPublicKey", destination).Error("Association not found")
		return
	}

	data := map[string]string{
		"assetCode": assetCode,
		"amount":    amount,
	}

	j, err := json.Marshal(data)
	if err != nil {
		s.log.WithField("data", data).Error("Error marshalling json")
	}

	s.SSEServer.BroadcastEvent(association.Address, sse.AccountCreditedAddressEvent, j)
}
