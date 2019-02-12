package service

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestQuery(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/circle_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Query(tt.args.db)
		})
	}
}
