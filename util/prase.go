package util

import (
	"errors"
	"net"
	"sort"
	"strconv"
	"strings"
)

// ParseHost parses s as a CIDR notation IP address and prefix length,
// return a string slice of IP address.
func ParseHost(s string) []string {
	res := make([]string, 0)
	ip, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		Fatal(s + " is not in correct CIDR format")
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incr(ip) {
		res = append(res, ip.String())
	}
	return res
}

// ParsePort parses s to a string slice of ports.
func ParsePort(s string) ([]uint16, error) {
	res := make([]uint16, 0)
	ports := strings.Split(strings.Trim(s, ", "), ",")
	for _, v := range ports {
		var start uint16
		multiPorts := strings.Split(strings.Trim(v, "-"), "-")
		start, err := portToUint16(multiPorts[0])
		if err != nil {
			Fatal(s + " is an incorrect port format string")
		}
		if len(multiPorts) > 1 {
			end, err := portToUint16(multiPorts[1])
			if err != nil {
				Fatal(s + " is an incorrect port format string")
			}
			for ; start != 0 && start <= end; start++ {
				res = append(res, start)
			}
		} else {
			res = append(res, start)
		}
	}
	res = deduplicate(res)
	sort.Slice(res, func(i, j int) bool {
		return i < j
	})
	return res, nil
}

// portToUint16 return the int value of port string.
func portToUint16(port string) (value uint16, err error) {
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

func incr(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i]++; ip[i] > 0 {
			break
		}
	}
}

func deduplicate(list []uint16) []uint16 {
	set := make(map[uint16]struct{}, len(list))
	j := 0
	for _, v := range list {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		list[j] = v
		j++
	}
	return list[:j]
}
