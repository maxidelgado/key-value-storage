package kvs

import (
	"reflect"
	"testing"
)

func Test_stack_isEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack stack
		want  bool
	}{
		{
			name:  "should be empty",
			stack: stack{},
			want:  true,
		},
		{
			name:  "should not be empty",
			stack: stack{transaction{}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.isEmpty(); got != tt.want {
				t.Errorf("isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_peek(t *testing.T) {
	tests := []struct {
		name    string
		stack   stack
		want    transaction
		wantErr bool
	}{
		{
			name:    "should return latest item in stack",
			stack:   stack{transaction{}},
			want:    transaction{},
			wantErr: false,
		},
		{
			name:    "should return error, no transactions in stack",
			stack:   stack{},
			want:    transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.peek()
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

func Test_stack_pop(t *testing.T) {
	tests := []struct {
		name    string
		stack   stack
		wantLen int
		want    transaction
		wantErr bool
	}{
		{
			name:    "should return latest transaction and remove it from the stack",
			stack:   stack{transaction{}, transaction{}},
			wantLen: 1,
			want:    transaction{},
			wantErr: false,
		},
		{
			name:    "should return error, no transactions in stack",
			stack:   stack{},
			want:    transaction{},
			wantLen: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(len(tt.stack), tt.wantLen) {
				t.Errorf("pop() len = %v, wantLent %v", len(tt.stack), tt.wantLen)
			}
		})
	}
}

func Test_stack_push(t *testing.T) {
	type args struct {
		storage transaction
	}
	var tests = []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name:    "should push an item to the stack",
			args:    args{storage: transaction{}},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{}
			s.push(tt.args.storage)
			if !reflect.DeepEqual(len(s), tt.wantLen) {
				t.Errorf("push() len = %v, wantLent %v", len(s), tt.wantLen)
			}
		})
	}
}
