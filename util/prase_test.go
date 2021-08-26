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
			got, err := ParsePort(tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
