package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_Merge(t *testing.T) {
	from := &Foo{
		//Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: Bar{
			Name: "qw",
			Flag: false,
		},
		Strs: []string{"a", "b"},
	}
	to := &Foo{
		Id: 11,
		Bar: Bar{
			Name: "aqw",
			Flag: true,
		},
		Strs: []string{"c", "b"},
	}
	err := Merge(to, from)
	assert.Equal(t, nil, err)
	fmt.Print(to)

}

func Test_Merge1(t *testing.T) {
	from := &Baz{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: &Bar{
			Name: "qw",
			Flag: true,
		},
	}
	to := &Foo{}
	err := Merge(to, from)
	assert.Equal(t, nil, err)
	fmt.Println(to)

	bazs := &[]Baz{{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: &Bar{
			Name: "qw",
			Flag: true,
		},
	}}
	foos := &[]Foo{{
		Id: 11,
	}}
	err = Merge(foos, bazs)
	assert.Equal(t, nil, err)
	fmt.Println(foos)

	foos = &[]Foo{}
	err = Merge(foos, bazs)
	assert.Equal(t, nil, err)
	fmt.Println(foos)

	bazs = &[]Baz{{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: &Bar{
			Name: "qw",
			Flag: true,
		},
		MBar: map[string]string{
			"a": "b",
			"c": "b",
		},
	}}
	foos = &[]Foo{{MBar: map[string]string{}}}
	err = Merge(foos, bazs)
	assert.Equal(t, nil, err)
	//assert.Equal(t, foos, bazs)
	fmt.Println(foos)

	foos = &[]Foo{}
	err = Merge(foos, bazs)
	assert.Equal(t, nil, err)
	//assert.Equal(t, foos, bazs)
	fmt.Println(foos)

	foos = &[]Foo{{MBar: map[string]string{"c": "d", "d": "d"}}}
	err = Merge(foos, bazs)
	assert.Equal(t, nil, err)
	//assert.Equal(t, foos, bazs)
	fmt.Println(foos)

	fs := &[]Foos{{MBar: nil}}
	err = Merge(fs, bazs)
	assert.Equal(t, nil, err)
	//assert.Equal(t, foos, bazs)
	fmt.Println(fs)

	bazs = &[]Baz{{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: &Bar{
			Name: "qw",
			Flag: true,
		},
	}}
	fs = &[]Foos{{MBar: nil}}
	err = Merge(fs, bazs)
	assert.Equal(t, nil, err)
	//assert.Equal(t, foos, bazs)
	fmt.Println(fs)
}

func Test_Merge2(t *testing.T) {
	from := &Baz{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: &Bar{
			Name: "qw",
			Flag: true,
		},
		CreateTime: time.Second * 10,
	}
	to := &Foo{}
	err := Merge(to, from)
	assert.Equal(t, nil, err)
	fmt.Println(to)

}

func Test_IsZero(t *testing.T) {
	bar := &Bar{
		Name: "asd",
		Flag: false,
	}

	flag := isZero(reflect.ValueOf(bar))
	assert.Equal(t, false, flag)
	m := make(map[string]string)
	flag = isZero(reflect.ValueOf(m))
	assert.Equal(t, true, flag)

	var mp = map[string]string{}
	flag = isZero(reflect.ValueOf(mp))
	assert.Equal(t, true, flag)

	v := ""
	flag = isZero(reflect.ValueOf(v))
	assert.Equal(t, true, flag)

	var str []string
	flag = isZero(reflect.ValueOf(str))
	assert.Equal(t, true, flag)

	str = []string{"1"}
	flag = isZero(reflect.ValueOf(str))
	assert.Equal(t, false, flag)
}


func Test_IsNil(t *testing.T) {
	bar := &Bar{
		Name: "asd",
		Flag: false,
	}

	flag := isNil(reflect.ValueOf(bar))
	assert.Equal(t, false, flag)
	m := make(map[string]string)
	flag = isNil(reflect.ValueOf(m))
	assert.Equal(t, false, flag)

	var mp = map[string]string{}
	flag = isNil(reflect.ValueOf(mp))
	assert.Equal(t, false, flag)

	v := ""
	flag = isNil(reflect.ValueOf(v))
	assert.Equal(t, true, flag)

	var str []string
	flag = isNil(reflect.ValueOf(str))
	assert.Equal(t, true, flag)

	str = []string{"1"}
	flag = isNil(reflect.ValueOf(str))
	assert.Equal(t, false, flag)
}