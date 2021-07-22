package utils

import "net"

// GenerateIPPool return all ip address according to CIDR
func GenerateIPPool(cidr string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	pool := make([]string, 0)
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incr(ip) {
		pool = append(pool, ip.String())
	}
	return pool[1 : len(pool)-1], nil
}

// incr increase the ip address by one.
func incr(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i]++; ip[i] > 0 {
			break
		}
	}
}
