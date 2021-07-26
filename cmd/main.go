package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/secriy/golire/scan"
)

var (
	ips     = flag.String("i", "127.0.0.1/24", "CIDR, like '192.168.1.0/24'")
	ports   = flag.String("p", "22,3389", "All legal ports, like '1-28', '22,53,3389' and '22,49-80'.")
	timeout = flag.Int("t", 2000, "Scan timeout for one task, set 2 seconds by default.")
)

func main() {
	flag.Parse()
	tasks := scan.NewTasks(*ips, *ports)
	wg := &sync.WaitGroup{}
	results := make([]string, len(tasks.IPPool))
	for k, v := range tasks.IPPool {
		wg.Add(1)
		go func(idx int, ip string) {
			defer wg.Done()
			s := scan.NewScan(ip, &tasks.PortPool, *timeout)
			portList := s.GetAllTcpOpenPorts()
			if len(*portList) != 0 {
				results[idx] = fmt.Sprintf("IP:%s Port:%v", ip, *portList)
				// fmt.Printf("IP:%s Port:%v\n", ip, *portList)
			}
		}(k, v)
	}
	wg.Wait()
	for _, v := range results {
		if v != "" {
			fmt.Println(v)
		}
	}
}
