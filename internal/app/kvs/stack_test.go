package kvs

import (
	"reflect"
	"testing"
)

func Test_transactionStack_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    transactionStack
		want bool
	}{
		{
			name: "should be empty",
			s:    transactionStack{},
			want: true,
		},
		{
			name: "should not be empty",
			s:    transactionStack{storage{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isEmpty(); got != tt.want {
				t.Errorf("isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionStack_peek(t *testing.T) {
	tests := []struct {
		name    string
		s       transactionStack
		want    storage
		wantErr bool
	}{
		{
			name:    "should return latest item in stack",
			s:       transactionStack{storage{}},
			want:    storage{},
			wantErr: false,
		},
		{
			name:    "should return error, no transactions in stack",
			s:       transactionStack{},
			want:    storage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionStack_pop(t *testing.T) {
	tests := []struct {
		name    string
		s       transactionStack
		wantLen int
		want    storage
		wantErr bool
	}{
		{
			name:    "should return latest transaction and remove it from the stack",
			s:       transactionStack{storage{}, storage{}},
			wantLen: 1,
			want:    storage{},
			wantErr: false,
		},
		{
			name:    "should return error, no transactions in stack",
			s:       transactionStack{},
			want:    storage{},
			wantLen: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(len(tt.s), tt.wantLen) {
				t.Errorf("pop() len = %v, wantLent %v", len(tt.s), tt.wantLen)
			}
		})
	}
}

func Test_transactionStack_push(t *testing.T) {
	type args struct {
		storage storage
	}
	var tests = []struct {
		name    string
		s       transactionStack
		args    args
		wantLen int
	}{
		{
			name:    "should push an item to the stack",
			s:       transactionStack{},
			args:    args{storage: storage{}},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.push(tt.args.storage)
			if !reflect.DeepEqual(len(tt.s), tt.wantLen) {
				t.Errorf("push() len = %v, wantLent %v", len(tt.s), tt.wantLen)
			}
		})
	}
}
