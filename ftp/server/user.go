package server

import (
	"fmt"
	"net"
)

// User method. Return a loged user
func (f *FTP) user(conn net.Conn,name []string) {
	f.respond(conn,fmt.Sprintf(status230,name[0]))
}