package util

import (
	"reflect"
	"testing"
)

func TestParsePort(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		{"test1", args{port: "1-3"}, []uint16{1, 2, 3}, false},
		{"test2", args{port: "1-3,2"}, []uint16{1, 2, 3}, false},
		{"test3", args{port: "1-3,2,4"}, []uint16{1, 2, 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePort(tt.args.port)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHost(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test1", args{s: "172.22.22.0/30"}, []string{"172.22.22.1", "172.22.22.2"}, false},
		{"single", args{s: "172.22.22.2"}, []string{"172.22.22.2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseHost(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
