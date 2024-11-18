package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("DOWN. \n%s is unreachable. \nError: %s", address, err)
	} else {
		status = fmt.Sprintf("UP. \n%s is reachable. \nFrom: %s\nTo: %s", address, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status

}
