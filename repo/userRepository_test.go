package repo

import (
	"reflect"
	"testing"

	"github.com/BohdanPatrash/indenticon/dto"
)

func TestSaveUser(t *testing.T) {
	Init()
	DatabaseSetup()
	type args struct {
		user *dto.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "firstTest",
			args:    args{user: &dto.User{Email: "test@mail.com"}},
			wantErr: false,
		},
		{
			name:    "secondTest",
			args:    args{user: &dto.User{Email: "test@mail.com"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			GetAllUsers()
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	Init()
	tests := []struct {
		name    string
		want    []dto.User
		wantErr bool
	}{
		{
			name:    "first test",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "second test",
			want:    make([]dto.User, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllUsers()
			DatabaseSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
