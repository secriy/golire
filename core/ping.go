package core

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

var host string

// PingDefault return is the host living by default arguments.
func PingDefault(host string) bool {
	return Ping(host, 3) // ping 3 times
}

// Ping implements the ICMP message of ping.
func Ping(addr string, count int) bool {
	host = addr
	for i := 0; i < count; i++ {
		raddr := net.IPAddr{IP: net.ParseIP(addr)} // remote address
		if err := sendICMPMsg(newICMPMsg(uint16(i)), &raddr); err == nil {
			return true
		}
	}
	return false
}

// sendICMPMsg send ICMP request.
func sendICMPMsg(icmp ICMP, destAddr *net.IPAddr) error {
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		pingFailed(err)
		return err
	}
	defer conn.Close()

	var buffer bytes.Buffer

	// write message to buffer
	if err := binary.Write(&buffer, binary.BigEndian, icmp); err != nil {
		pingFailed(err)
		return err
	}

	// send message
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		pingFailed(err)
		return err
	}

	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 2))

	rev := make([]byte, 1024)
	if _, err := conn.Read(rev); err != nil {
		pingFailed(err)
		return err
	}
	return nil
}

// newICMPMsg return a new ICMP message.
func newICMPMsg(seq uint16) ICMP {
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}

	var buffer bytes.Buffer
	if err := binary.Write(&buffer, binary.BigEndian, icmp); err != nil {
		pingFailed(err)
		return icmp
	}

	icmp.CheckSum = util.CheckSum(buffer.Bytes()) // calc checksum

	buffer.Reset()
	return icmp
}

// debug message
func pingFailed(err error) {
	util.Debug("ping " + host + " failed: " + err.Error())
}
