package scan

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Scan struct {
	IP      string
	Port    uint16
	Timeout time.Duration
}

// NewScan return the Scan object.
func NewScan(ip string, port uint16, timeout time.Duration) *Scan {
	return &Scan{ip, port, timeout}
}

// TCP return the state of the specific tcp port with input host.
func (s *Scan) TCP() bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port), s.Timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println("连接数超出系统限制！" + err.Error())
			os.Exit(1)
		}
		return false
	}
	defer conn.Close()
	return true
}
