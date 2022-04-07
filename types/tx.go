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
}

func GeneratePaymentTxJson(from, to, amount string) (string, error) {
	payment := &Payment{
		TransactionType: "payment",
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
