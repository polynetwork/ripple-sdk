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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/polynetwork/ripple-sdk/types"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rubblelabs/ripple/websockets"
)

//RpcClient for ontology rpc api
type RpcClient struct {
	addr       string
	httpClient *http.Client
}

//NewRpcClient return RpcClient instance
func NewRpcClient() *RpcClient {
	return &RpcClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false, //enable keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
			},
			Timeout: time.Second * 300, //timeout for http response
		},
	}
}

//SetAddress set rpc server address. Simple http://localhost:20336
func (this *RpcClient) SetAddress(addr string) *RpcClient {
	this.addr = addr
	return this
}

//SetHttpClient set http client to RpcClient. In most cases SetHttpClient is not necessary
func (this *RpcClient) SetHttpClient(httpClient *http.Client) *RpcClient {
	this.httpClient = httpClient
	return this
}

func (this *RpcClient) GetCurrentHeight() (uint32, error) {
	respData, err := this.sendRpcRequest(RPC_LEDGER_CLOSED, []interface{}{})
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight: send req err: %s", err)
	}
	result := &heightResp{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	if result.Result.Status != "success" {
		return 0, fmt.Errorf("GetCurrentHeight, resp failed, status: %s", result.Result.Status)
	}
	return result.Result.LedgerIndex, nil
}

func (this *RpcClient) GetLedger(height uint32) (*websockets.LedgerResult, error) {
	ledgerReqParam := ledgerReqParam{
		LedgerIndex:  height,
		Transactions: true,
		Expand:       true,
	}
	respData, err := this.sendRpcRequest(RPC_LEDGER, []interface{}{ledgerReqParam})
	if err != nil {
		return nil, fmt.Errorf("GetLedger: send req err: %s", err)
	}
	result := &websockets.LedgerCommand{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return nil, fmt.Errorf("GetLedger: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	return result.Result, nil
}

//SignFor sign method for multi-sign account
func (this *RpcClient) SignFor(account, secret string, txJson *types.MultisignPayment) (*SignRes, error) {
	sigForReqParam := sigForReqParam{
		Account: account,
		Secret:  secret,
		TxJson:  txJson,
	}
	respData, err := this.sendRpcRequest(RPC_SIGN_FOR, []interface{}{sigForReqParam})
	if err != nil {
		return nil, fmt.Errorf("SignFor: send req err: %s", err)
	}
	result := &SignRes{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return nil, fmt.Errorf("SignFor: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	return result, nil
}

func (this *RpcClient) SubmitMultisigned(txJson *types.MultisignPayment) (*SubmitMultisignRes, error) {
	submitMultisignedTxReq := submitMultisignedTxReq{
		TxJson: txJson,
	}
	respData, err := this.sendRpcRequest(RPC_SUBMIT_MULTISIGNED, []interface{}{submitMultisignedTxReq})
	if err != nil {
		return nil, fmt.Errorf("SubmitMultisigned: send req err: %s", err)
	}
	submitRes := &SubmitMultisignRes{}
	err = json.Unmarshal(respData, submitRes)
	if err != nil {
		return nil, fmt.Errorf("SubmitMultisigned: unmarshal submit tx resp err: %s", err)
	}
	return submitRes, nil
}

func (this *RpcClient) GetAccountInfo(account string) (*websockets.AccountInfoResult, error) {
	accountReqParam := accountInfoReqParam{
		Account: account,
		Strict:  true,
		Queue:   false,
	}
	respData, err := this.sendRpcRequest(RPC_ACCOUNT_INFO, []interface{}{accountReqParam})
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo: send req err: %s", err)
	}
	result := &websockets.AccountInfoCommand{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	return result.Result, nil
}

func (this *RpcClient) GetFee() (*websockets.FeeResult, error) {
	respData, err := this.sendRpcRequest(RPC_FEE, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("GetFee: send req err: %s", err)
	}
	result := &websockets.FeeCommand{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return nil, fmt.Errorf("GetFee: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	return result.Result, nil
}

//Tx return the tx info of hash
func (this *RpcClient) GetTx(hash string) (*websockets.TxResult, error) {
	txReqParam := txReqParam{
		Transaction: hash,
		Binary:      false,
	}
	respData, err := this.sendRpcRequest(RPC_TX, []interface{}{txReqParam})
	if err != nil {
		return nil, fmt.Errorf("GetTx: send req err: %s", err)
	}
	result := &websockets.TxCommand{}
	err = json.Unmarshal(respData, result)
	if err != nil {
		return nil, fmt.Errorf("GetTx: unmarshal resp err: %s, origin resp is %s", err, string(respData))
	}
	return result.Result, nil
}

//sendRpcRequest send Rpc request to ripple
func (this *RpcClient) sendRpcRequest(method string, params []interface{}) ([]byte, error) {
	rpcReq := &JsonRpcRequest{
		Method: method,
		Params: params,
	}
	data, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("JsonRpcRequest json.Marsha error:%s", err)
	}
	resp, err := this.httpClient.Post(this.addr, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, PostErr{fmt.Errorf("http post request:%s error:%s", data, err)}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rpc response body error:%s", err)
	}
	return body, nil
}

type PostErr struct {
	Err error
}

func (err PostErr) Error() string {
	return err.Err.Error()
}
