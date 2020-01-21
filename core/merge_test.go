package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Foo struct {
	Id       int    `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Bar      Bar    `json:"bar"`
}

type Baz struct {
	Id       int    `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Bar      *Bar   `json:"bar"`
}

type Bar struct {
	Name string `json:"name"`
	Flag bool   `json:"flag"`
}

func TestSetDefault(t *testing.T) {
	bar := Bar{
		Name: "asd",
		Flag: false,
	}
	foo := Foo{
		Id:       13,
		User:     "clare",
		Password: "xxxxx",
		Bar:      bar,
	}
	Info(&foo)
	fmt.Println(foo)
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
	assert.Equal(t, false, flag)

	var mp = map[string]string{}
	flag = isZero(reflect.ValueOf(mp))
	assert.Equal(t, false, flag)

	flag = isZero(reflect.ValueOf(mp))
	assert.Equal(t, false, flag)

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

func Test_Replace(t *testing.T) {
	from := &Foo{
		Id:       12,
		User:     "wq",
		Password: "wwww",
		Bar: Bar{
			Name: "qw",
			Flag: true,
		},
	}
	to := &Foo{}
	err := replace(to, from)
	assert.Equal(t, nil, err)
	fmt.Print(to)

}

func Test_Replace1(t *testing.T) {
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
	err := replace(to, from)
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
	err = replace(foos, bazs)
	assert.Equal(t, nil, err)
	fmt.Println(foos)

	foos = &[]Foo{}
	err = replace(foos, bazs)
	assert.Equal(t, nil, err)
	fmt.Println(foos)

}
