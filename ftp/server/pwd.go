package server

import (
	"strings"
	"fmt"
	"net"
)

// Pwd method. Return currently path
func (f FTP) pwd(conn net.Conn) {
	remoteDir := strings.Split(f.workdir,"/")
	f.respond(conn,fmt.Sprintf(status257,remoteDir[len(remoteDir) - 1]))
}