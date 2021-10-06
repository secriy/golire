package module

import (
	"testing"

	"github.com/secriy/golire/util"
)

func TestPing(t *testing.T) {
	type args struct {
		domain string
		count  int
	}
	tests := []struct {
		name     string
		args     args
		wantLive bool
	}{
		{"acgfate", args{"47.100.15.172", 3}, true},
		{"baidu", args{"220.181.38.251", 3}, true},
		{"local", args{"127.0.0.1", 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			util.SetLevel(util.LevelDebug)
			if gotLive := Ping(tt.args.domain, tt.args.count); gotLive != tt.wantLive {
				t.Errorf("Ping() = %v, want %v", gotLive, tt.wantLive)
			}
		})
	}
}
