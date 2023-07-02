package server

import (
	"net"
)

// conType method. Define type of connection
func (f *FTP) conType(conn net.Conn, data []string) {
	if len(data) != 1 {
		f.respond(conn,status501)
		return 
	}
	switch data[0] {
	case "A":
		f.eofile = "\r\n"
	case "I":
		f.eofile = "\n"
	default:
		f.eofile = "\n"
	}
	f.respond(conn,status200)
}
