package scan

import (
	"testing"
)

func TestGetAllTcpOpenPorts(t *testing.T) {
	ports := make([]uint16, 0)
	for i := 1; i < 3000; i++ {
		ports = append(ports, uint16(i))
	}
	scan := NewScan("172.22.22.239", &ports, 3000)
	res := scan.GetAllTcpOpenPorts()
	t.Log(res)
}
