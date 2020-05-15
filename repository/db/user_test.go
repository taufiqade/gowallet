package db

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
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
		name string
		fields fields
		args args
		want models.Users
		wantErr bool
		configureMock func (fields fields, args args, want models.Users)  
	}{
		// test for positif testing
		{
			name: "successfully retrieve user by id",
			fields: fields{
				db: db,
				mock: mock,
			},
			args: args{
				id: 2,
			}, 
			want: models.Users{
				ID: 2,
				Email: "user2@test.com",
				Password: "$2a$04$FUaLfg5Dd3Vyo1C2P.zEeubhROc7qoEdx1k7jC5HtiK9gUsCWcd6O",
				Name: "user2",
				Type: "user",
				// UpdatedAt: "2020-05-06T07:04:06+07:00",
				// CreatedAt: "2020-05-06T07:04:06+07:00",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (id=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.
					NewRows([]string{"id","email","password","name","type"}).
					AddRow(want.ID,want.Email,want.Password,want.Name,want.Type)
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
		// test for negative testing
		{
			name: "Not found",
			fields: fields{
				db: db,
				mock: mock,
			},
			args: args{
				id: 2,
			},
			want: models.Users{},
			wantErr: true,
			configureMock: func(fields fields, args args, want models.Users) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (id=?) ORDER BY `users`.`id` ASC LIMIT 1")
				selectRows := sqlmock.NewRows([]string{"id","email","password","name","type"})
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
