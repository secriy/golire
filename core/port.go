package core

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/secriy/golire/util"
)

// TODO 扫描准确率低

// DetectTCPWithPort return the state of the specific tcp port with input host,
// need IP address and port number specified.
func DetectTCPWithPort(addr string, port uint16, timeout time.Duration) bool {
	return DetectTCP(fmt.Sprintf("%s:%d", addr, port), timeout)
}

// DetectUDPWithPort return the state of the specific udp port with input host,
// need IP address and port number specified.
func DetectUDPWithPort(addr string, port uint16, timeout time.Duration) bool {
	return DetectUDP(fmt.Sprintf("%s:%d", addr, port), timeout)
}

// DetectTCP return the state of the specific tcp port with input host.
func DetectTCP(host string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			util.Fatal("The number of connections out of limit: " + err.Error())
		}
		util.Debug("Detect TCP port error: " + err.Error())
		return false
	}
	defer conn.Close()
	return true
}

// DetectUDP return the state of the specific udp port with input host.
func DetectUDP(host string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("udp", host, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			util.Fatal("The number of connections out of limit: " + err.Error())
		}
		util.Debug("Detect UDP port error: " + err.Error())
		return false
	}
	defer conn.Close()
	return true
}
