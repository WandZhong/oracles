package setup

import (
	"errors"
	"regexp"
)

// Assert panics and breaks the program if the err is not nil
func Assert(err error) {
	if err != nil {
		logger.Fatal("Can't perform initialization", err)
	}
}

func checkAppName(name string) error {
	const nameRegexp = "^[[:alnum:]\\-_.]{2,200}"
	if ok, err := regexp.Match(nameRegexp, []byte(name)); err != nil {
		return errors.New("Can't parse regular expression for name check. " + err.Error())
	} else if !ok {
		return errors.New("Wrong app name. Should match the following regexp: " + nameRegexp)
	}
	return nil
}
