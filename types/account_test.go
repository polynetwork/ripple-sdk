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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/rubblelabs/ripple/data"
	"testing"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/stretchr/testify/assert"
)

func TestImportAccount(t *testing.T) {
	account, err := ImportAccount("shtew2z1TRsEvpnYUGtiyvqPnYywt")
	assert.Nil(t, err)
	var zeroSequence uint32
	accountId, err := crypto.AccountId(account.Key, &zeroSequence)
	assert.Nil(t, err)
	assert.Equal(t, "rLi6oSF38EdP7mzhdccyxhfd8vp8FWbsWF", accountId.String())
}

func TestNewAccount(t *testing.T) {
	account_m, wallet, err := NewAccount()
	assert.Nil(t, err)

	account_n, err := ImportAccount(wallet.Seed)
	assert.Nil(t, err)
	assert.Equal(t, account_m.Key, account_n.Key)
	assert.Equal(t, account_m.Account, account_n.Account)
}

func TestAddressToAccount(t *testing.T) {
	account, err := data.NewAccountFromAddress("rsHYGX2AoQ4tXqFywzEeeTDgXFTUfL1Fw9")
	assert.Nil(t, err)
	fmt.Println(hex.EncodeToString(account.Bytes()))
}

func TestMultiSign(t *testing.T) {
	signer, err := ImportAccount("shtew2z1TRsEvpnYUGtiyvqPnYywt")
	to, _ := data.NewAccountFromAddress("rT4vRkeJsgaq7t6TVJJPsbrQp5oKMGRfN")
	from, _ := data.NewAccountFromAddress("rsHYGX2AoQ4tXqFywzEeeTDgXFTUfL1Fw9")
	amount, _ := data.NewAmount("13000000")
	fee, _ := data.NewValue("50", false)
	feeCost, _ := data.NewAmount("50")
	amount, _ = amount.Subtract(feeCost)
	memoType, _ := hex.DecodeString("706f6c7968617368")
	memoData, _ := hex.DecodeString("706f6c7968617368")
	memos := data.Memos{}
	memo := data.Memo{}
	memo.Memo.MemoType = memoType
	memo.Memo.MemoData = memoData
	memos = append(memos, memo)

	payment := GeneratePayment(*from, *to, *amount, *fee, 25336389, memos)
	_, raw, err := data.Raw(payment)
	assert.Nil(t, err)
	p, err := signer.MultiSignTx(hex.EncodeToString(raw))
	assert.Nil(t, err)
	r, _ := json.Marshal(p)
	fmt.Println(string(r))

	// test check multi sign
	err = CheckMultiSign(hex.EncodeToString(raw), p.Signers[0].Signer.Account,
		p.Signers[0].Signer.SigningPubKey.Bytes(), *p.Signers[0].Signer.TxnSignature)
	assert.Nil(t, err)
}
