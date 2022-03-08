package kvs

import (
	"testing"
)

func Test_KVS_Commit(t *testing.T) {
	type fields struct {
		stack  *stack
		mainTx transaction
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should commit child to parent transaction",
			fields: fields{
				stack:  &stack{transaction{}, transaction{}},
				mainTx: newTransaction(),
			},
			wantErr: false,
		},
		{
			name: "should commit parent transaction",
			fields: fields{
				stack:  &stack{transaction{}},
				mainTx: newTransaction(),
			},
			wantErr: false,
		},
		{
			name: "should return error, no transaction found",
			fields: fields{
				stack:  &stack{},
				mainTx: newTransaction(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kvs{
				stack:  tt.fields.stack,
				mainTx: tt.fields.mainTx,
			}
			if err := k.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_KVS_Set(t *testing.T) {
	type fields struct {
		stack  *stack
		mainTx transaction
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
			name: "should set key/value pair to kvs",
			fields: fields{
				stack:  &stack{},
				mainTx: newTransaction(),
			},
			args: args{
				key:   "key",
				value: "value",
			},
		},
		{
			name: "should set key to a transaction in the stack",
			fields: fields{
				stack:  &stack{newTransaction()},
				mainTx: newTransaction(),
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
				stack:  tt.fields.stack,
				mainTx: tt.fields.mainTx,
			}
			k.Set(tt.args.key, tt.args.value)
		})
	}
}
