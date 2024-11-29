package version

import "fmt"

var (
	app     = "qrcp"
	version = "0.11.1"
	author  = "PhateValleyman"
	mail    = "Jonas.Ned@outlook.com"
)

// String returns a string representation of the build.
func String() string {
	return fmt.Sprintf("%s v%s\n%s\n%s", app, version, author, mail)
}
