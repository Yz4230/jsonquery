package jsonquery

import (
	"fmt"
	"reflect"
)

type Map = map[string]interface{}
type Arr = []interface{}

type JsonQuery struct {
	doc interface{}
	err error
}

func New(doc interface{}) *JsonQuery {
	// Create instance with
	if reflect.TypeOf(doc).Kind() == reflect.Ptr {
		return &JsonQuery{*doc.(*interface{}), nil}
	}
	return &JsonQuery{doc, nil}
}

func (j *JsonQuery) Key(key string) *JsonQuery {
	if j.err != nil {
		return j
	}
	switch doc := j.doc.(type) {
	case Map:
		if val, ok := doc[key]; ok {
			return &JsonQuery{val, nil}
		}
		return &JsonQuery{nil, fmt.Errorf("key %s does not exists", key)}
	case Arr:
		var docarr []interface{}
		for _, d := range doc {
			if dm, ok := d.(Map); ok {
				if v, exists := dm[key]; exists {
					docarr = append(docarr, v)
				}
			} else {
				return &JsonQuery{nil, fmt.Errorf("%+v is not a map", d)}
			}
		}
		return &JsonQuery{docarr, nil}
	}

	return &JsonQuery{nil, fmt.Errorf("%+v is not a map", j.doc)}
}

func (j *JsonQuery) At(index int64) *JsonQuery {
	if j.err != nil {
		return j
	}
	doc := reflect.Indirect(reflect.ValueOf(j.doc)).Interface() // dereference
	if doc, ok := doc.(Arr); ok {
		if index < int64(len(doc)) {
			return &JsonQuery{doc[index], nil}
		}
		return &JsonQuery{nil, fmt.Errorf("index %d is out of range", index)}
	}
	return &JsonQuery{nil, fmt.Errorf("%+v is not an array", j.doc)}
}

func (j *JsonQuery) Expand() *JsonQuery {
	if j.err != nil {
		return j
	}
	doc := reflect.Indirect(reflect.ValueOf(j.doc)).Interface() // dereference
	if doc, ok := doc.(Arr); ok {
		var docarr []interface{}
		for _, d := range doc {
			if dm, ok := d.(Arr); ok {
				docarr = append(docarr, dm...)
			} else {
				return &JsonQuery{nil, fmt.Errorf("%+v is not an array", d)}
			}
		}
		return &JsonQuery{docarr, nil}
	}
	return &JsonQuery{nil, fmt.Errorf("%+v is not an array", j.doc)}
}

func (j *JsonQuery) End() (doc interface{}, err error) {
	return j.doc, j.err
}
