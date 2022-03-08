package kvs

import (
	"reflect"
	"testing"
)

func Test_mapKVS_ApplyChanges(t *testing.T) {
	type fields struct {
		storage map[string]string
		count   map[string]int
	}
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantStorage map[string]string
		wantCount   map[string]int
	}{
		{
			name: "should replace b key with a new value",
			fields: fields{
				storage: map[string]string{"a": "1", "b": "2"},
				count:   map[string]int{"1": 1, "2": 1},
			},
			args: args{
				m: map[string]string{"b": "3"},
			},
			wantStorage: map[string]string{"a": "1", "b": "3"},
			wantCount:   map[string]int{"1": 1, "2": 0, "3": 1},
		},
		{
			name: "should add c key and its value",
			fields: fields{
				storage: map[string]string{"a": "1", "b": "2"},
				count:   map[string]int{"1": 1, "2": 1},
			},
			args: args{
				m: map[string]string{"c": "3"},
			},
			wantStorage: map[string]string{"a": "1", "b": "2", "c": "3"},
			wantCount:   map[string]int{"1": 1, "2": 1, "3": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				memoryMap: tt.fields.storage,
				count:     tt.fields.count,
			}
			s.ApplyChanges(tt.args.m)
			if !reflect.DeepEqual(s.memoryMap, tt.wantStorage) {
				t.Errorf("s.memoryMap = %v, want %v", s.memoryMap, tt.wantStorage)
			}
			if !reflect.DeepEqual(s.count, tt.wantCount) {
				t.Errorf("s.count = %v, want %v", s.count, tt.wantCount)
			}
		})
	}
}

func Test_mapKVS_Count(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		count map[string]int
		args  args
		want  int
	}{
		{
			name:  "should return proper count",
			count: map[string]int{"1": 2},
			args: args{
				value: "1",
			},
			want: 2,
		},
		{
			name:  "should return zero count",
			count: map[string]int{"1": 2},
			args: args{
				value: "3",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				count: tt.count,
			}
			if got := s.Count(tt.args.value); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapKVS_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		storage map[string]string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return a value",
			args: args{
				key: "key",
			},
			storage: map[string]string{"key": "value"},
			want:    "value",
			wantErr: false,
		},
		{
			name: "should return error",
			args: args{
				key: "not-existing-key",
			},
			storage: map[string]string{"key": "value"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := storage{
				memoryMap: tt.storage,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
