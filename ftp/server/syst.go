package server

import (
	"fmt"
	"net"
	"runtime"
)

// Syst method. Return system type
func (f FTP) syst(conn net.Conn) {
	os := runtime.GOOS
	var platform string
	switch os {
	case "windows":
		platform = fmt.Sprintf(status215, os)
	case "linux":
		platform = fmt.Sprintf(status215, os)
	case "darwin":
		platform = fmt.Sprintf(status215, os)
	}
	f.respond(conn, platform)

}
