package server

import (
	"path/filepath"
	"fmt"
	"os"
	"net"
)

// Size method. Return file size
func (f FTP) size(conn net.Conn,data []string) {
	if len(data) != 1 {
		f.respond(conn,status501)
		return
	}
	info, err := os.Stat(filepath.Join(f.workdir,data[0]))
	if os.IsExist(err) {
		f.respond(conn,status550)
	}
	f.respond(conn,fmt.Sprintf(status213,info.Size()))

}