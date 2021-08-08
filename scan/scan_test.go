package scan

import (
	"log"
	"testing"
	"time"
)

func TestScan_TCP(t *testing.T) {
	sc := NewScan("127.0.0.1", 22, time.Millisecond*200)
	if sc.TCP() {
		log.Println("YES")
	} else {
		log.Println("NO")
	}
}
