package module

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"

	"github.com/secriy/golire/util"
)

// ICMP message
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

const (
	MaxPg = 2000
)

// PingDefault return is the host living by default arguments.
func PingDefault(host string) bool {
	return Ping(host, 48, 3)
}

// Ping implements the ICMP message of ping.
func Ping(domain string, PS, count int) (live bool) {
	var (
		icmp        = ICMP{8, 0, 0, 0, 0}
		laddr       = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		raddr, _    = net.ResolveIPAddr("ip", domain)
		originBytes = make([]byte, MaxPg)
	)
	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)
	if err != nil {
		pingFailed(domain, err)
		return
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(time.Duration(1000))); err != nil {
		pingFailed(domain, err)
		return false
	}

	buf := new(bytes.Buffer)
	// write binary data to buffer
	err = binary.Write(buf, binary.BigEndian, icmp)
	if err != nil {
		pingFailed(domain, err)
		return
	}
	err = binary.Write(buf, binary.BigEndian, originBytes[0:PS])
	if err != nil {
		pingFailed(domain, err)
		return
	}
	b := buf.Bytes()
	binary.BigEndian.PutUint16(b[2:], util.CheckSum(b))

	rev := make([]byte, 1024)

	// ping count times and if one of them succeed return true,
	// otherwise return false.
	for ; count > 0; count-- {
		if _, err := conn.Write(buf.Bytes()); err != nil {
			pingFailed(domain, err)
			time.Sleep(time.Second)
			continue
		}
		err = conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100))
		if err != nil {
			pingFailed(domain, err)
			return
		}
		_, err = conn.Read(rev)
		if err != nil {
			pingFailed(domain, err)
			time.Sleep(time.Second)
			continue
		}
		return true
	}
	return
}

// debug message
func pingFailed(domain string, err error) {
	util.Debug("ping " + domain + " failed: " + err.Error())
}
