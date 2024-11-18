package main

import (
	"fmt"
	"net"
	"os"
)

func getBroadcastAddress(interfaceName string) (string, error) {
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", fmt.Errorf("could not find interface %s: %v", interfaceName, err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", fmt.Errorf("could not get addresses for interface %s: %v", interfaceName, err)
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.To4() == nil {
			continue
		}

		ip := ipNet.IP.To4()
		mask := ipNet.Mask
		broadcast := net.IPv4(
			ip[0]|^mask[0],
			ip[1]|^mask[1],
			ip[2]|^mask[2],
			ip[3]|^mask[3],
		)

		return broadcast.String(), nil
	}

	return "", fmt.Errorf("no suitable IPv4 address found on interface %s", interfaceName)
}

func broadcastMessage(broadcastAddr string, message string) error {
	conn, err := net.Dial("udp", broadcastAddr+":12345")
	if err != nil {
		return fmt.Errorf("could not dial broadcast address: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("could not send message: %v", err)
	}

	return nil
}

func main() {
	interfaceName := "Ethernet 4"
	message := "Hello, broadcast!"

	broadcastAddr, err := getBroadcastAddress(interfaceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting broadcast address: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Broadcasting to %s on interface %s\n", broadcastAddr, interfaceName)

	if err := broadcastMessage(broadcastAddr, message); err != nil {
		fmt.Fprintf(os.Stderr, "Error broadcasting message: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Broadcast message sent!")
}
