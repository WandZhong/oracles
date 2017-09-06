package utils

// Assert panics when err is not nil
func Assert(err error, msg string) {
	if err != nil {
		logger.Fatal(msg, err)
	}
}
