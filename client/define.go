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

package client

import "github.com/polynetwork/ripple-sdk/types"

const (
	RPC_TX                 = "tx"
	RPC_FEE                = "fee"
	RPC_ACCOUNT_INFO       = "account_info"
	RPC_SIGN_FOR           = "sign_for"
	RPC_SIGN               = "sign"
	RPC_SUBMIT             = "submit"
	RPC_SUBMIT_MULTISIGNED = "submit_multisigned"
	RPC_LEDGER_CLOSED      = "ledger_closed"
	RPC_LEDGER             = "ledger"
)

type JsonRpcRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type txReqParam struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}

type accountInfoReqParam struct {
	Account string `json:"account"`
	Strict  bool   `json:"strict"`
	Queue   bool   `json:"queue"`
}

type sigForReqParam struct {
	Account string                  `json:"account"`
	Secret  string                  `json:"secret"`
	TxJson  *types.MultisignPayment `json:"tx_json"`
}

type SignRes struct {
	Result struct {
		Status       string                  `json:"status"`
		TxBlob       string                  `json:"tx_blob"`
		TxJson       *types.MultisignPayment `json:"tx_json"`
		ErrorMessage string                  `json:"error_message"`
	} `json:"result"`
}

type SubmitRes struct {
	Result struct {
		Status              string         `json:"status"`
		TxBlob              string         `json:"tx_blob"`
		TxJson              *types.Payment `json:"tx_json"`
		ErrorMessage        string         `json:"error_message"`
		EngineResult        string         `json:"engine_result"`
		EngineResultMessage string         `json:"engine_result_message"`
	} `json:"result"`
}

type SubmitMultisignRes struct {
	Result struct {
		Status              string                  `json:"status"`
		TxBlob              string                  `json:"tx_blob"`
		TxJson              *types.MultisignPayment `json:"tx_json"`
		ErrorMessage        string                  `json:"error_message"`
		EngineResult        string                  `json:"engine_result"`
		EngineResultMessage string                  `json:"engine_result_message"`
	} `json:"result"`
}

type submitTxReq struct {
	TxBlob string `json:"tx_blob"`
}

type submitMultisignedTxReq struct {
	TxJson *types.MultisignPayment `json:"tx_json"`
}

type heightResp struct {
	Result struct {
		LedgerHash  string `json:"ledger_hash"`
		LedgerIndex uint32 `json:"ledger_index"`
		Status      string `json:"status"`
	} `json:"result"`
}

type ledgerReqParam struct {
	LedgerIndex  uint32 `json:"ledger_index"`
	Transactions bool   `json:"transactions"`
	Expand       bool   `json:"expand"`
}
