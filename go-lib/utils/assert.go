package utils

func Assert(err error, msg string) {
	if err != nil {
		logger.Fatal(msg, err)
	}
}
