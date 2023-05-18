package debug

import "os"

var (
	DEBUG bool
)

func init() {
	DEBUG = os.Getenv("DEBUG") != ""
}
