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
	"fmt"
	"math/rand"
	"time"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/data"
)

type Account struct {
	Account data.Account
	Key     crypto.Key
}

type Wallet struct {
	Address string
	Seed    string
}

func ImportAccount(secret string) (*Account, error) {
	account := &Account{}
	accountSeedContent, err := crypto.NewRippleHash(secret)
	if err != nil {
		return nil, fmt.Errorf("ImportAccount: cannot parse account secret, err: %s", err)
	}
	accountSeed := accountSeedContent.Payload()
	accountKey, err := crypto.NewECDSAKey(accountSeed)
	if err != nil {
		return nil, fmt.Errorf("ImportAccount: construct account secret, err: %s", err)
	}
	account.Key = accountKey

	var signSequence uint32
	accountAddr, err := crypto.AccountId(accountKey, &signSequence)
	if err != nil {
		return nil, fmt.Errorf("ImportAccount：gen account addr failed, err: %s", err)
	}
	var acc data.Account
	copy(acc[:], accountAddr.Payload())
	account.Account = acc
	return account, nil
}

func NewAccount() (*Account, *Wallet, error) {
	timeStamp := time.Now().Unix()
	randomSeed := rand.NewSource(timeStamp)
	randomG := rand.New(randomSeed)
	accountSeed, err := crypto.GenerateFamilySeed(fmt.Sprintf("%f", randomG.Float64()))
	if err != nil {
		return nil, nil, fmt.Errorf("new account secret failed, err: %s", err)
	}
	accountKey, err := crypto.NewECDSAKey(accountSeed.Payload())
	if err != nil {
		return nil, nil, fmt.Errorf("new account key failed, err: %s", err)
	}

	var signSequence uint32
	accountAddr, err := crypto.AccountId(accountKey, &signSequence)
	if err != nil {
		return nil, nil, fmt.Errorf("new account addr failed, err: %s", err)
	}

	var addr data.Account
	copy(addr[:], accountAddr.Payload())

	return &Account{Account: addr, Key: accountKey},
		&Wallet{accountAddr.String(), accountSeed.String()}, nil
}

func (this *Account) SignTx(rawTx string) (hash, signedTx string, err error) {
	tx, err := deserializeRawTx(rawTx)
	if err != nil {
		return "", "", fmt.Errorf("SignAccountTx: deserialized tx failed, err: %s", err)
	}
	if tx.GetBase().Account != this.Account {
		return "", "", fmt.Errorf("SignAccountTx: tx account not self account")
	}
	return signTx(tx, this.Key)
}

func signTx(tx data.Transaction, key crypto.Key) (hash, signedTx string, err error) {
	var signTxSequence uint32
	err = data.Sign(tx, key, &signTxSequence)
	if err != nil {
		return "", "", fmt.Errorf("signTx: sign tx failed, err: %s", err)
	}
	txHash, signedTxData, err := data.Raw(tx)
	if err != nil {
		return "", "", fmt.Errorf("signTx: serialize signed tx failed, err: %s", err)
	}
	hash = txHash.String()
	signedTx = fmt.Sprintf("%x", signedTxData)
	return hash, signedTx, nil
}
