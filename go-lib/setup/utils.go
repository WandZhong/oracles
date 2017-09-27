package setup

import (
	"errors"

	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
)

func checkAppName(name string) error {
	if ok := utils.ReName.Match([]byte(name)); !ok {
		return errors.New("Wrong app name. Should match the following regexp: " +
			utils.ReName.String())
	}
	return nil
}

func assert(err error) {
	if err != nil {
		logger.Fatal("Assertion error", err)
	}
}
