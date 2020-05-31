package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock "github.com/taufiqade/gowallet/mock"
	"github.com/taufiqade/gowallet/models"
)

func Test_GetUserByID(t *testing.T) {
	type fields struct {
		userRepo *mock.DBUserRepositoryMock
	}

	type args struct {
		id int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	defaultDate, _ := time.Parse("2006-01-02 15:04:05", "2020-03-29 00:00:00")
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.Users
		wantErr       bool
		configureMock func(fields fields, args args, want models.Users)
	}{
		{
			name: "Successfully returned user",
			fields: fields{
				userRepo: mock.NewDBUserRepositoryMock(ctrl),
			},
			args: args{
				id: 1,
			},
			want: models.Users{
				ID:        1,
				Email:     "test@testing.com",
				Password:  "",
				Name:      "test",
				Type:      "admin",
				UpdatedAt: defaultDate.Add(time.Hour * 24),
				CreatedAt: defaultDate.Add(time.Hour * 24),
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.Users) {
				fields.userRepo.
					EXPECT().
					GetUserByID(args.id).
					Return(want, nil)
			},
		}, {
			name: "User not found",
			fields: fields{
				userRepo: mock.NewDBUserRepositoryMock(ctrl),
			},
			args: args{
				id: 0,
			},
			want:    models.Users{},
			wantErr: true,
			configureMock: func(fields fields, args args, want models.Users) {
				fields.userRepo.
					EXPECT().
					GetUserByID(args.id).
					Return(want, errors.New("user not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserService{
				userRepo: tt.fields.userRepo,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := service.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}

}
