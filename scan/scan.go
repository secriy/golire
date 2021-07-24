package scan

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Scan struct {
	IP      string
	Port    *[]uint16
	Timeout time.Duration
}

// NewScan return the Scan object.
func NewScan(ip string, port *[]uint16, timeout int) *Scan {
	return &Scan{ip, port, time.Duration(timeout)}
}

// GetAllTcpOpenPorts return all tcp ports which are opening.
func (s *Scan) GetAllTcpOpenPorts() *[]uint16 {
	wg := &sync.WaitGroup{}
	openingPorts := make([]uint16, 0)
	for _, v := range *s.Port {
		wg.Add(1)
		go func(port uint16) {
			if s.IsTcpPortOpen(port) {
				openingPorts = append(openingPorts, port)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	return &openingPorts
}

// IsTcpPortOpen scan the specified address and port.
func (s *Scan) IsTcpPortOpen(port uint16) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.IP, port), time.Millisecond*s.Timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println("连接数超出系统限制！" + err.Error())
			os.Exit(1)
		}
		return false
	}
	_ = conn.Close()
	return true
}
