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
	"math/rand"
	"time"

	"github.com/polynetwork/ripple-sdk/client"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RippleSdk is the main struct for user
type RippleSdk struct {
	client.ClientMgr
}

// NewPolySdk return RippleSdk.
func NewRippleSdk() *RippleSdk {
	rippleSdk := &RippleSdk{}
	return rippleSdk
}
