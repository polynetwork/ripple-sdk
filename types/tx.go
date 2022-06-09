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
	"fmt"
	"github.com/rubblelabs/ripple/data"
)

type MultisignPayment struct {
	TransactionType string
	Account         string
	Destination     string
	Amount          string
	Fee             string
	Sequence        uint32
	SigningPubKey   string
	Memos           []Memo    `json:"Memos,omitempty"`
	Signers         []*Signer `json:"Signers,omitempty"`
	Hash            string    `json:"Hash,omitempty"`
}

type Memo struct {
	Memo struct {
		MemoType   string
		MemoData   string
		MemoFormat string
	}
}

type Signer struct {
	Signer struct {
		Account       string
		SigningPubKey string
		TxnSignature  string
	} `json:"Signer"`
}

func GeneratePayment(from, to data.Account, amount data.Amount, fee data.Value, sequence uint32) *data.Payment {
	payment := &data.Payment{
		Destination: to,
		Amount:      amount,
	}
	txBase := data.TxBase{
		TransactionType: data.PAYMENT,
		Account:         from,
		Sequence:        sequence,
		Fee:             fee,
	}
	payment.TxBase = txBase
	return payment
}

func DeserializeRawMultiSignTx(rawTx string) (*data.Payment, error) {
	txData, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, fmt.Errorf("deserializeRawTx: cannot decode raw tx, err: %s", err)
	}
	tx, err := data.ReadTransaction(bytes.NewReader(txData))
	if err != nil {
		return nil, fmt.Errorf("deserializeRawTx: parse raw tx failed, err: %s", err)
	}
	payment := tx.(*data.Payment)
	payment.InitialiseForMultiSigning()
	return payment, nil
}
