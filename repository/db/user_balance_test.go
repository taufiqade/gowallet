package db

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
)

func Test_dbUserBalance_GetByUserID(t *testing.T) {
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
		want          models.UserBalance
		wantErr       bool
		configureMock func(fields fields, args args, want models.UserBalance)
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
			want: models.UserBalance{
				ID:             1,
				UserID:         2,
				Balance:        1000,
				BalanceAchieve: 1100,
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.UserBalance) {
				selectQuery := regexp.QuoteMeta("SELECT * FROM `user_balance` WHERE (user_id=?) ORDER BY `user_balance`.`id` ASC LIMIT 1")
				selectRows := sqlmock.
					NewRows([]string{"id", "user_id", "balance", "balance_achieve"}).
					AddRow(want.ID, want.UserID, want.Balance, want.BalanceAchieve)
				fields.mock.
					ExpectQuery(selectQuery).
					WillReturnRows(selectRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserBalanceRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			got, err := repository.GetByUserID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbUserBalance_Update(t *testing.T) {
	type fields struct {
		db   *gorm.DB
		mock sqlmock.Sqlmock
	}

	type args struct {
		id   int
		data models.UserBalance
	}

	sqlMockConn, mock, _ := sqlmock.New()
	db, _ := gorm.Open("mysql", sqlMockConn)

	tests := []struct {
		name          string
		fields        fields
		args          args
		want          models.UserBalance
		wantErr       bool
		configureMock func(fields fields, args args, want models.UserBalance)
	}{
		{
			name: "Create transaction history",
			fields: fields{
				db:   db,
				mock: mock,
			},
			args: args{
				id: 1,
				data: models.UserBalance{
					UserID:         2,
					Balance:        1000,
					BalanceAchieve: 1100,
				},
			},
			want: models.UserBalance{
				ID:             1,
				UserID:         2,
				Balance:        1000,
				BalanceAchieve: 1100,
			},
			wantErr: false,
			configureMock: func(fields fields, args args, want models.UserBalance) {
				fields.mock.ExpectBegin()

				insertQuery := regexp.QuoteMeta("UPDATE `user_balance` SET `balance` = ?, `balance_achieve` = ?, `user_id` = ? WHERE (user_id=?)")
				fields.mock.
					ExpectExec(insertQuery).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(int64(want.ID), 1))

				fields.mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &UserBalanceRepository{
				DB: tt.fields.db,
			}
			tt.configureMock(tt.fields, tt.args, tt.want)
			err := repository.Update(tt.args.id, &tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
