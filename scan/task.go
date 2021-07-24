package scan

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/secriy/golire/utils"
)

// Tasks include ip address, mask and port pool of target host.
type Tasks struct {
	IPPool   []string
	PortPool []uint16
}

// NewTasks return Tasks object.
func NewTasks(ips string, ports string) *Tasks {
	tasks := &Tasks{}
	if err := tasks.GetAllIP(ips); err != nil {
		log.Fatalln(err)
	}
	if err := tasks.GetAllPorts(ports); err != nil {
		log.Fatalln(err)
	}
	return tasks
}

// GetAllIP return all target IP address.
func (t *Tasks) GetAllIP(ip string) (err error) {
	t.IPPool, err = utils.GenerateIPPool(ip)
	return
}

// GetAllPorts return all target ports.
func (t *Tasks) GetAllPorts(port string) (err error) {
	ports := strings.Split(strings.Trim(port, ", "), ",")
	for _, v := range ports {
		multiPorts := strings.Split(strings.Trim(v, "-"), "-")
		start, err := t.PortTransfer(multiPorts[0])
		if err != nil {
			continue
		}
		if len(multiPorts) > 1 {
			end, err := t.PortTransfer(multiPorts[1])
			if err != nil {
				continue
			}
			for ; start <= end; start++ {
				t.PortPool = append(t.PortPool, start)
			}
		} else {
			t.PortPool = append(t.PortPool, start)
		}
	}
	t.PortPool = utils.Deduplicate(t.PortPool)
	return
}

// PortTransfer return the int value of port string.
func (t *Tasks) PortTransfer(port string) (value uint16, err error) {
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
