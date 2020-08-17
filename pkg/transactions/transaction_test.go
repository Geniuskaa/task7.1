package transaction

import (
	"reflect"
	"testing"
)

func Test_mapRowToTransaction(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want Transaction
	}{ {"export.csv", args{[]string{"01","0001","0003","1000000","1597597057"}}, Transaction{
		Id:      "01",
		From:    "0001",
		To:      "0003",
		Amount:  1000000,
		Created: 1597597057,
	}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapRowToTransaction(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapRowToTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}