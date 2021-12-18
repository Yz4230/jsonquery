package jsonquery

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJsonQuery_Key(t *testing.T) {
	type fields struct {
		doc interface{}
		err error
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *JsonQuery
	}{
		{
			"key exists in map 1",
			fields{Map{"a": 1}, nil},
			args{"a"},
			&JsonQuery{1, nil},
		},
		{
			"key exists in map 2",
			fields{Map{"a": Map{"b": 1}}, nil},
			args{"a"},
			&JsonQuery{Map{"b": 1}, nil},
		},
		{
			"key does not exists in map",
			fields{Map{"a": 1}, nil},
			args{"b"},
			&JsonQuery{nil, fmt.Errorf("key %s does not exists", "b")},
		},
		{
			"key in array",
			fields{Arr{Map{"a": 1}, Map{"a": 2}, Map{"b": 3}}, nil},
			args{"a"},
			&JsonQuery{Arr{1, 2}, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JsonQuery{
				Doc: tt.fields.doc,
				Err: tt.fields.err,
			}
			if got := j.Key(tt.args.key); !cmp.Equal(got.Doc, tt.want.Doc) && !errors.Is(got.Err, tt.want.Err) {
				t.Errorf("JsonQuery.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
