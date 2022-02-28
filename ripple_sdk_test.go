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

package ripple_sdk

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPC(t *testing.T) {
	sdk := NewRippleSdk()
	sdk.NewRpcClient().SetAddress("https://s.altnet.rippletest.net:51234/")
	r, err := sdk.GetRpcClient().GetAccountInfo("rLi6oSF38EdP7mzhdccyxhfd8vp8FWbsWF")
	assert.Nil(t, err)
	temp, _ := json.Marshal(r)
	fmt.Println(string(temp))
}