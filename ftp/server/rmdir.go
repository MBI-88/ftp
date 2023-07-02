package server

import (
	"net"
)

// Rmdir method. Remove dir
func (f FTP) rmdir(conn net.Conn, dirs []string) {
	if len(dirs) != 1 {
		f.respond(conn, status501)
		return
	}
	f.helperRmdir(conn,dirs[0])

}
