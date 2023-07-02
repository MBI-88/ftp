package server

import (
	"fmt"
	"net"
)


func (f FTP) chmod(conn net.Conn,data []string) {
	fmt.Println(data)
	if len(data) != 0 && data[0] != "CHMOD" {
		f.respond(conn,status501)
		return
	}
	f.helperChmod(conn,data[1],data[2:])
}