package golire

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/secriy/golire/scan"
	"github.com/secriy/golire/utils"
)

// Tasks include ip address, mask and port pool of target host.
type Tasks struct {
	IPPool   []string
	PortPool []uint16
}

type PortPool chan uint16

// Run execute all tasks with concurrency.
func (t *Tasks) Run(timeout time.Duration, gNum int) {
	wg := &sync.WaitGroup{}
	for _, ip := range t.IPPool {
		pool := make(PortPool, 100)

		go putPort(pool, &t.PortPool) // put port into channel

		for i := 0; i < gNum; i++ {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()

				for port := range pool {
					sc := scan.NewScan(ip, port, timeout)
					if sc.TCP() {
						fmt.Printf("%s:%d open\n", ip, port)
					}
				}
			}(ip)
		}
		wg.Wait()
	}
}

// NewTasks return a Tasks instance.
func NewTasks(ips string, ports string) (*Tasks, error) {
	tasks := new(Tasks)
	if err := tasks.allIP(ips); err != nil {
		return nil, err
	}
	if len(tasks.IPPool) < 1 {
		return nil, errors.New("no surviving host")
	}
	if err := tasks.allPorts(ports); err != nil {
		return nil, err
	}
	return tasks, nil
}

// allIP return all target IP address.
func (t *Tasks) allIP(cidr string) (err error) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incr(ip) {
		wg.Add(1)
		go func(host string) {
			if scan.MustPing(host) {
				defer mu.Unlock()
				mu.Lock()
				t.IPPool = append(t.IPPool, host)
			}
			wg.Done()
		}(ip.String())
	}
	wg.Wait()
	return nil
}

// allPorts return all target ports.
func (t *Tasks) allPorts(port string) (err error) {
	ports := strings.Split(strings.Trim(port, ", "), ",")
	for _, v := range ports {
		var start, end uint16
		multiPorts := strings.Split(strings.Trim(v, "-"), "-")
		start, err = portTransfer(multiPorts[0])
		if err != nil {
			return err
		}
		if len(multiPorts) > 1 {
			end, err = portTransfer(multiPorts[1])
			if err != nil {
				return
			}
			for ; start != 0 && start <= end; start++ {
				t.PortPool = append(t.PortPool, start)
			}
		} else {
			t.PortPool = append(t.PortPool, start)
		}
	}
	t.PortPool = utils.Deduplicate(t.PortPool)
	return
}

// putPort put the port into PortPool channel.
func putPort(pool PortPool, ports *[]uint16) {
	for _, port := range *ports {
		pool <- port
	}
	close(pool)
}

// portTransfer return the int value of port string.
func portTransfer(port string) (value uint16, err error) {
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

// incr increase the ip address by one.
func incr(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i]++; ip[i] > 0 {
			break
		}
	}
}
