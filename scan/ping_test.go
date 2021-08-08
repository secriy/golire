package scan

import (
	"log"
	"testing"
)

func TestPingToScan(t *testing.T) {
	if MustPing("127.0.0.1") {
		log.Println("YES")
	}
}
