package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	}
	to := &Foo{
		Id: 11,
		Bar: Bar{
			Name: "aqw",
			Flag: true,
		},
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
