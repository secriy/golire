package scan

import (
	"errors"
	"strconv"
	"strings"

	"github.com/secriy/golire/utils"
)

// // PortPool specific the range of ports.
// type PortPool struct {
// 	Start uint16
// 	Stop  uint16
// }

// Scan include ip address, mask and port pool of target host.
type Scan struct {
	IPPool   []string
	PortPool []uint16
	Process  int
}

// GetAllIP return all target IP address.
func (s *Scan) GetAllIP(ip string) (err error) {
	s.IPPool, err = utils.GenerateIPPool(ip)
	return
}

// GetAllPorts return all target ports.
func (s *Scan) GetAllPorts(port string) (err error) {
	ports := strings.Split(strings.Trim(port, ", "), ",")
	for _, v := range ports {
		multiPorts := strings.Split(strings.Trim(v, "-"), "-")
		start, err := s.PortTransfer(multiPorts[0])
		if err != nil {
			continue
		}
		if len(multiPorts) > 1 {
			end, err := s.PortTransfer(multiPorts[1])
			if err != nil {
				continue
			}
			for ; start <= end; start++ {
				s.PortPool = append(s.PortPool, start)
			}
		} else {
			s.PortPool = append(s.PortPool, start)
		}
	}
	s.PortPool = utils.Deduplicate(s.PortPool)
	return
}

// PortTransfer return the int value of port string.
func (s *Scan) PortTransfer(port string) (value uint16, err error) {
	v, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return
	}
	value = uint16(v)
	if value < 1 {
		err = errors.New("port number out of range")
	}
	return
}
