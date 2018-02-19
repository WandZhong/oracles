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
	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/robert-zaremba/errstack"
)

func distributeSWC(summaries []directbuy.Summary) errstack.E {
	_, cf := flags.MustEthFactory()
	swcC, addr, err := cf.GetSWC()
	utils.Assert(err, "Can't instantiate SWT contract")
	logger.Debug("Contract address", "swc", addr.Hex())

	checkOK(directbuy.CheckRequiredBalance(summaries, "SWC", swcC, cf))
	// TODO - whitelists are failing
	// checkOK(checkWhitelist(summaries, cf)

	if *flags.dryRun {
		logger.Debug("Dry run. Stopping execution.")
		return nil
	}
	return directbuy.Distribute(summaries, swcC, cf, nil)
}

/* Not finished. Tests are still failing
func checkWhitelist(summaries []directbuy.Summary, cf ethereum.ContractFactory) errstack.E {
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

	for _, r := range summaries {
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
