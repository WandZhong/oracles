// Copyright (c) 2017 Sweetbridge Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/robert-zaremba/errstack"
)

func distributeSWC(records []Record) {
	_, cf := flags.MustEthFactory()
	swcC, addr, err := cf.GetSWC()
	utils.Assert(err, "Can't instantiate SWT contract")
	logger.Debug("Contract address", "swc", addr.Hex())

	checkOK(checkSWCbalance(records, swcC, cf))
	// TODO - whitelists are failing
	// checkOK(checkWhitelist(records, cf)
	checkOK(transferSWC(records, swcC, cf))
}

func transferSWC(records []Record, swcC *contracts.SweetToken, cf ethereum.ContractFactory) errstack.E {
	if *flags.dryRun {
		logger.Debug("Dry run. Stopping execution.")
		return nil
	}

	txo := cf.Txo()
	for _, r := range records {
		logger.Debug("Transfering", "SWC", wad.WeiToString(r.Amount),
			"dest", r.Address.Hex(), "nonce", txo.Nonce)
		tx, err := swcC.Transfer(txo, r.Address, r.Amount)
		if err != nil {
			logger.Error("Can't transfer TOKEN", err)
			break
		} else {
			ethereum.LogTx("Transferred", tx)
			ethereum.IncTxoNonce(txo, tx)
			// logger.Debug(">>>> nonce after", "txo", txo.Nonce, "tx", tx.Nonce())
		}
	}
	return nil
}

/* Not finished. Tests are still failing
func checkWhitelist(records []Record, cf ethereum.ContractFactory) errstack.E {
	swcLogic, _, err := cf.GetSWClogic()
	utils.Assert(err, "Can't instantiate SWTlogic contract")

	var listName = [32]byte{}
	copy(listName[:], []byte("whitelist"))
	logger.Debug("Constructing whitelist name", "bytes", listName)
	txo := cf.Txo()
	if ok, err := swcLogic.ListExists(nil, listName); err != nil {
		return errstack.WrapAsInf(err, "Can't check if whitelist exists")
	} else if !ok {
		tx, err := swcLogic.AddWhiteList(txo, listName)
		if err != nil {
			return errstack.WrapAsInf(err, "Can't create SWC whitelist")
		}
		ethereum.LogTx("SWC whitelist created", tx)
	}

	for _, r := range records {
		if ok, err := swcLogic.WhiteLists(nil, r.Address, listName); err != nil {
			return errstack.WrapAsInfF(err, "Can't check if user %s is in the whitelist",
				r.Address)
		} else if ok {
			continue
		}
		tx, err := swcLogic.AddToWhiteList(txo, listName, r.Address)
		if err != nil {
			return errstack.WrapAsInfF(err, "Can't add user %s to the whitelist", r.Address)
		}
		ethereum.LogTx(fmt.Sprintf("User %s added to the whitelist", r.Address), tx)
	}
	return nil

}
*/

func checkSWCbalance(records []Record, token *contracts.SweetToken, cf ethereum.ContractFactory) errstack.E {
	var total = new(big.Int)
	for _, r := range records {
		total.Add(total, r.Amount)
	}
	addr := cf.Addr()
	logger.Debug("SWC distribution account holder", "address", addr.Hex())
	balance, err := token.BalanceOf(nil, addr)
	if err != nil {
		return errstack.WrapAsInf(err, "Can't check SWC balance")
	}
	if balance.Cmp(total) < 0 {
		return errstack.NewReqF("Not enough funds in the source account = %v, min_expected=%v",
			wad.WeiToString(balance), wad.WeiToInt(total))
	}
	logger.Debug("Distribution account balance", "swc.wei", balance.String(),
		"required", total.String())
	return nil
}
