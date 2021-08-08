package scan

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"

	"github.com/secriy/golire/utils"
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

func MustPing(host string) bool {
	_, ok := Ping(host, 48, 3)
	return ok
}

// Ping implements ping command by ICMP
func Ping(domain string, PS, count int) (err error, live bool) {
	var (
		icmp          = ICMP{8, 0, 0, 0, 0}
		localAddr     = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		remoteAddr, _ = net.ResolveIPAddr("ip", domain)
		originBytes   = make([]byte, MaxPg)
	)
	conn, err := net.DialIP("ip4:icmp", &localAddr, remoteAddr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()

	var buffer bytes.Buffer
	// write binary data to buffer
	err = binary.Write(&buffer, binary.BigEndian, icmp)
	if err != nil {
		return
	}
	err = binary.Write(&buffer, binary.BigEndian, originBytes[0:PS])
	if err != nil {
		return
	}
	b := buffer.Bytes()
	binary.BigEndian.PutUint16(b[2:], utils.CheckSum(b))

	rev := make([]byte, 1024)

	// count the lost
	dropPack := 0
	for i := count; i > 0; i-- {
		if _, err := conn.Write(buffer.Bytes()); err != nil {
			dropPack++
			time.Sleep(time.Second)
			continue
		}
		err = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
		if err != nil {
			return
		}
		_, err = conn.Read(rev)
		if err != nil {
			dropPack++
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Millisecond)
	}
	live = dropPack != count
	return
}
