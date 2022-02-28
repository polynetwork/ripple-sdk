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
	"testing"

	"github.com/rubblelabs/ripple/crypto"
	"github.com/stretchr/testify/assert"
)

func TestImportAccount(t *testing.T) {
	account, err := ImportAccount("shtew2z1TRsEvpnYUGtiyvqPnYywt")
	assert.Nil(t, err)
	var zeroSequence uint32
	accountId, err := crypto.AccountId(account.key, &zeroSequence)
	assert.Nil(t, err)
	assert.Equal(t, "rLi6oSF38EdP7mzhdccyxhfd8vp8FWbsWF", accountId.String())
}

func TestNewAccount(t *testing.T) {
	account_m, wallet, err := NewAccount()
	assert.Nil(t, err)

	account_n, err := ImportAccount(wallet.Seed)
	assert.Nil(t, err)
	assert.Equal(t, account_m.key, account_n.key)
	assert.Equal(t, account_m.account, account_n.account)
}
