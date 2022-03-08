package kvs

import (
	"testing"
)

func Test_transactionalKVS_Commit(t *testing.T) {
	type fields struct {
		transactions *transactionStack
		storage      storage
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should commit child to parent transaction",
			fields: fields{
				transactions: &transactionStack{storage{}, storage{}},
				storage:      newStorage(),
			},
			wantErr: false,
		},
		{
			name: "should commit parent transaction",
			fields: fields{
				transactions: &transactionStack{storage{}},
				storage:      newStorage(),
			},
			wantErr: false,
		},
		{
			name: "should return error, no transaction found",
			fields: fields{
				transactions: &transactionStack{},
				storage:      newStorage(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kvs{
				transactions: tt.fields.transactions,
				storage:      tt.fields.storage,
			}
			if err := k.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transactionalKVS_Set(t *testing.T) {
	type fields struct {
		transactions *transactionStack
		storage      storage
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "should set key to main kvs",
			fields: fields{
				transactions: &transactionStack{},
				storage:      newStorage(),
			},
			args: args{
				key:   "key",
				value: "value",
			},
		},
		{
			name: "should set key to current transaction kvs",
			fields: fields{
				transactions: &transactionStack{newStorage()},
				storage:      newStorage(),
			},
			args: args{
				key:   "key",
				value: "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kvs{
				transactions: tt.fields.transactions,
				storage:      tt.fields.storage,
			}
			k.Set(tt.args.key, tt.args.value)
		})
	}
}
