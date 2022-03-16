/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/rubblelabs/ripple/data"
)

type Payment struct {
	TransactionType string
	Account         string
	Destination     string
	Amount          string
	Hash            string
}

type MultisignPayment struct {
	TransactionType string
	Account         string
	Destination     string
	Amount          string
	Fee             string
	Sequence        uint32
	SigningPubKey   string
	Signers         []*Signer `json:"Signers,omitempty"`
	Hash            string
}

type Signer struct {
	Signer struct {
		Account       string
		SigningPubKey string
		TxnSignature  string
	} `json:"Signer"`
}

func GeneratePaymentTxJson(from, to, amount string) (string, error) {
	payment := &Payment{
		TransactionType: "Payment",
		Account:         from,
		Destination:     to,
		Amount:          amount,
	}
	r, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func GenerateMultisignPaymentTxJson(from, to, amount, fee string, sequence uint32) (string, error) {
	payment := &MultisignPayment{
		TransactionType: "Payment",
		Account:         from,
		Destination:     to,
		Amount:          amount,
		Fee:             fee,
		Sequence:        sequence,
		SigningPubKey:   "",
	}
	r, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func deserializeRawTx(rawTx string) (data.Transaction, error) {
	txData, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, fmt.Errorf("deserializeRawTx: cannot decode raw tx, err: %s", err)
	}
	tx, err := data.ReadTransaction(bytes.NewReader(txData))
	if err != nil {
		return nil, fmt.Errorf("deserializeRawTx: parse raw tx failed, err: %s", err)
	}
	return tx, nil
}
