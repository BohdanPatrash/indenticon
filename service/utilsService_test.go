package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTesting(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Check if is testing",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Testing(); got != tt.want {
				t.Errorf("Testing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHash(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "first test",
			args: args{data: "test@mail.com"},
			want: []byte{240, 226, 162, 27, 207, 73, 156, 188, 113, 60, 71, 216, 240, 52, 214, 110, 144, 169, 159, 159, 252, 254, 150, 70, 108, 153, 113, 223, 220, 92, 152, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHash(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
