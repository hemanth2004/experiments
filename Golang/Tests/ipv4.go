package main

import (
	"fmt"
	"strconv"
	"strings"
)

type octet = uint8
type ipv4Address [4]octet

func FromString(s string) (addr ipv4Address) {
	s = strings.TrimSpace(s)
	splits := strings.Split(s, ".")
	if len(splits) == 4 {
		for i := 0; i < 4; i++ {
			num, err := strconv.Atoi(splits[i])
			if err == nil {
				addr[i] = uint8(num)
			}
		}
	}
	return addr
}
func (addr ipv4Address) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}
func (addr1 ipv4Address) And(addr2 ipv4Address) (res ipv4Address) {
	for i := 0; i < 4; i++ {
		res[i] = addr1[i] & addr2[i]
	}
	return res
}
func (addr ipv4Address) Not() (res ipv4Address) {
	for i := 0; i < 4; i++ {
		res[i] = ^res[i]
	}
	return res
}

type port uint16

func main() {
	ipv4 := FromString("192.168.200.24")
	subnet := FromString("255.255.0.0")
	networkAddress := ipv4.And(subnet)
	fmt.Println(networkAddress.ToString())
	fmt.Println(subnet.Not().ToString())
}
