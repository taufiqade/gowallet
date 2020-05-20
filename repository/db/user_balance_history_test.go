package db

import (
	"reflect"
	"testing"
	"regexp"

	"github.com/taufiqade/gowallet/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func Test_dbUserBalanceHistory_GetBalanceID(t *testing.T) {
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
		want          models.UserBalanceHistory
		wantErr       bool
		configureMock func(fields fields, args args, want models.UserBalanceHistory)
	}{
		{
			name: "successfully retrieve history by id",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				id: 2,
			},
			want: models.UserBalanceHistory{
				ID: 1,
				UserBalanceID: 1001,
				BalanceBefore: 1000,
				BalanceAfter: 1100,
				Activity: "transfer",
				Type: "debit",
				Location: "local",
				IP: "127.0.0.1",
				UserAgent: "CHROME",
				Author: "Taufiq",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.UserBalanceHistory) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `user_balance_history` WHERE (user_balance_id=?) ORDER BY `user_balance_history`.`id` ASC LIMIT 1")
				selectRows := sqlmock.
					NewRows([]string{"id", "user_balance_id","balance_before","balance_after","activity","type","location","ip","user_agent","author"}).
					AddRow(want.ID,want.UserBalanceID,want.BalanceBefore,want.BalanceAfter,want.Activity,want.Type,want.Location,want.IP,want.UserAgent,want.Author)
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
		{
			name: "not found",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				id: 2,
			},
			want: models.UserBalanceHistory{},
			wantErr: true,
			configureMock: func(fields fields, args args, want models.UserBalanceHistory) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `user_balance_history` WHERE (user_balance_id=?) ORDER BY `user_balance_history`.`id` ASC LIMIT 1")
				selectRows := sqlmock.NewRows([]string{"id", "user_balance_id","balance_before","balance_after","activity","type","location","ip","user_agent","author"})
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserBalanceHistoryRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := repository.GetBalanceID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalanceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalanceID() got = %v, want %v", got, tt.want)
			}
		});
	};
}

func Test_dbUserBalanceHistory_Create(t *testing.T) {
	type fields struct {
		db   *gorm.DB
		mock sqlmock.Sqlmock
	}

	type args struct {
		data models.UserBalanceHistory
	}

	sqlMockConn, mock, _ := sqlmock.New()
	db, _ := gorm.Open("mysql", sqlMockConn)

	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.UserBalanceHistory
		wantErr       bool
		configureMock func(fields fields, args args, want models.UserBalanceHistory)
	}{
		{
			name: "Create transaction history",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				data: models.UserBalanceHistory{
					UserBalanceID: 1001,
					BalanceBefore: 1000,
					BalanceAfter: 1100,
					Activity: "transfer",
					Type: "debit",
					Location: "local",
					IP: "127.0.0.1",
					UserAgent: "CHROME",
					Author: "Taufiq",
				},
			},
			want: models.UserBalanceHistory{
				ID: 1,
				UserBalanceID: 1001,
				BalanceBefore: 1000,
				BalanceAfter: 1100,
				Activity: "transfer",
				Type: "debit",
				Location: "local",
				IP: "127.0.0.1",
				UserAgent: "CHROME",
				Author: "Taufiq",
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.UserBalanceHistory) {
				fields.mock.ExpectBegin()

				insertQuery := regexp.QuoteMeta("INSERT INTO `user_balance_history` (`user_balance_id`,`balance_before`,`balance_after`,`activity`,`type`,`location`,`ip`,`user_agent`,`author`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
				fields.mock.
					ExpectExec(insertQuery).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(int64(want.ID), 1))

				fields.mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserBalanceHistoryRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			err := repository.Create(&tt.args.data, repository.DB)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}