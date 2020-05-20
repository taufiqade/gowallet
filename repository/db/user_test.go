package db

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

func Test_dbUser_GetUserByID(t *testing.T) {
	type fields struct {
		db   *gorm.DB
		mock sqlmock.Sqlmock
	}

	type args struct {
		id int
	}

	sqlMockConn, mock, _ := sqlmock.New()
	db, _ := gorm.Open("mysql", sqlMockConn)

	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.Users
		wantErr       bool
		configureMock func(fields fields, args args, want models.Users)
	}{
		// test for positif testing
		{
			name: "successfully retrieve user by id",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				id: 2,
			},
			want: models.Users{
				ID:       2,
				Email:    "user2@test.com",
				Password: "$2a$04$FUaLfg5Dd3Vyo1C2P.zEeubhROc7qoEdx1k7jC5HtiK9gUsCWcd6O",
				Name:     "user2",
				Type:     "user",
				// UpdatedAt: "2020-05-06T07:04:06+07:00",
				// CreatedAt: "2020-05-06T07:04:06+07:00",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (id=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.
					NewRows([]string{"id", "email", "password", "name", "type"}).
					AddRow(want.ID, want.Email, want.Password, want.Name, want.Type)
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
		// test for negative testing
		{
			name: "Not found",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				id: 2,
			},
			want:    models.Users{},
			wantErr: true,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (id=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.NewRows([]string{"id", "email", "password", "name", "type"})
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := repository.GetUserByID(tt.args.id)
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

func Test_dbUser_GetUserByEmail(t *testing.T) {
	type fields struct {
		db   *gorm.DB
		mock sqlmock.Sqlmock
	}

	type args struct {
		email string
	}

	sqlMockConn, mock, _ := sqlmock.New()
	db, _ := gorm.Open("mysql", sqlMockConn)

	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.Users
		wantErr       bool
		configureMock func(fields fields, args args, want models.Users)
	}{
		// test for positif testing
		{
			name: "successfully retrieve user by id",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				email: "user2@test.com",
			},
			want: models.Users{
				ID:       2,
				Email:    "user2@test.com",
				Password: "$2a$04$FUaLfg5Dd3Vyo1C2P.zEeubhROc7qoEdx1k7jC5HtiK9gUsCWcd6O",
				Name:     "user2",
				Type:     "user",
				// UpdatedAt: "2020-05-06T07:04:06+07:00",
				// CreatedAt: "2020-05-06T07:04:06+07:00",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (email=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.
					NewRows([]string{"id", "email", "password", "name", "type"}).
					AddRow(want.ID, want.Email, want.Password, want.Name, want.Type)
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
		// test for negative testing
		{
			name: "Not found",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				email: "user2@test.com",
			},
			want:    models.Users{},
			wantErr: true,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (email=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.NewRows([]string{"id", "email", "password", "name", "type"})
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := repository.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbUser_Create(t *testing.T) {
	type fields struct {
		db   *gorm.DB
		mock sqlmock.Sqlmock
	}

	type args struct {
		data httpRequest.UserRequest
	}

	sqlMockConn, mock, _ := sqlmock.New()
	db, _ := gorm.Open("mysql", sqlMockConn)

	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.Users
		wantErr       bool
		configureMock func(fields fields, args args, want models.Users)
	}{
		{
			name: "Create user",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				data: httpRequest.UserRequest{
					Email: "test@gmail.com",
					Name:  "Taufiq Adesurya",
				},
			},
			want: models.Users{
				ID:    1,
				Email: "test@gmail.com",
				Name:  "Taufiq Adesurya",
				Type:  "admin",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.Users) {
				fields.mock.ExpectBegin()

				insertQuery := regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`,`name`,`type`,`updated_at`,`created_at`) VALUES (?,?,?,?,?,?)")
				fields.mock.
					ExpectExec(insertQuery).
					WithArgs(args.data.Email, sqlmock.AnyArg(), args.data.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(int64(want.ID), 1))

				fields.mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := repository.Create(&tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID != tt.want.ID {
				t.Errorf("Create() got ID = %d, want %d", got.ID, tt.want.ID)
			}
			if got.Name != tt.want.Name {
				t.Errorf("Create() got name = %s, want %s", got.Name, tt.want.Name)
			}
		})
	}
}
