package utils

import "os"

// Assert panics when err is not nil
func Assert(err error, msg string) {
	if err != nil {
		logger.Fatal(msg, err)
	}
}

// AssertIsFile panic if `path` is not a file.
// `key` is the log argument to name the path entity
func AssertIsFile(path, key string) {
	if stat, err := os.Stat(path); err != nil || stat.IsDir() {
		logger.Fatal("Expected a valid file path.", key, path, err)
	}

}
