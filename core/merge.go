package merge

import (
	"errors"
	"fmt"
	"reflect"
)

func Info(o interface{}) {
	v := reflect.ValueOf(o)
	GetField(v)
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

func GetField(v reflect.Value) {
	vr := indirect(v)
	t := vr.Type()
	for i := 0; i < vr.NumField(); i++ {
		field := t.Field(i)
		kind := vr.Field(i).Kind()
		switch kind {
		case reflect.Int:
			name := field.Name
			vr.Field(i).SetInt(1)
			fmt.Println("name:", name)
		case reflect.String:
			name := field.Name
			vr.Field(i).SetString("ssssss")
			fmt.Println("name:", name)
		case reflect.Bool:
			name := field.Name
			fmt.Println("name:", name)
		case reflect.Struct:
			GetField(reflect.Indirect(v).Field(i))
		}
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

func replace(to interface{}, from interface{}) error {
	toVal := reflect.ValueOf(to)
	fromVal := reflect.ValueOf(from)
	if !isPtr(toVal) || !isPtr(fromVal) {
		return errors.New("copy to value is unaddressable")
	}
	Copy(toVal, fromVal)
	return nil
}

func Copy(to reflect.Value, from reflect.Value) {
	tVa := indirect(to)
	tt := tVa.Type()
	fVa := indirect(from)
	ft := fVa.Type()
	if tVa.Kind() == reflect.Slice {
		if fVa.Kind() == reflect.Slice {
			CopySlice(tVa, fVa)
			return
		}
	}
	for j := 0; j < tVa.NumField(); j++ {
		tField := tt.Field(j)
		for i := 0; i < fVa.NumField(); i++ {
			fField := ft.Field(i)
			fKind := fVa.Field(i).Kind()
			switch fKind {
			case reflect.Int:
				if tField.Name == fField.Name {
					tVa.Field(i).SetInt(fVa.FieldByName(fField.Name).Int())
					break
				}
			case reflect.String:
				if tField.Name == fField.Name {
					tVa.Field(i).SetString(fVa.FieldByName(fField.Name).String())
					break
				}
			case reflect.Bool:
				if tField.Name == fField.Name {
					tVa.Field(i).SetBool(fVa.FieldByName(fField.Name).Bool())
					break
				}
			case reflect.Struct, reflect.Ptr:
				// to and from is Struct
				if (fVa.Field(j).Kind() == reflect.Struct || fVa.Field(j).Kind() == reflect.Ptr) && tField.Name == fField.Name {
					Copy(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j))
					break
				}
			}
		}
	}

}

func isPtr(val reflect.Value) bool {
	if reflect.Ptr == val.Kind() {
		return true
	}
	return false
}

func CopySlice(to reflect.Value, from reflect.Value) {
	toLen := to.Len()
	for i := 0; i < from.Len(); i++ {
		j := 0
		if j < toLen {
			Copy(to.Index(j), from.Index(i))
			j++
			continue
		} else {
			var elem reflect.Value
			typ := to.Type().Elem()
			if typ.Kind() == reflect.Ptr {
				elem = reflect.New(typ.Elem())
			}
			if typ.Kind() == reflect.Struct {
				elem = reflect.New(typ).Elem()
			}
			Copy(elem, from.Index(i))
			to.Set(reflect.Append(to, elem))
		}
	}

}

func deepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	if reflectType = indirectType(reflectType); reflectType.Kind() == reflect.Struct {
		for i := 0; i < reflectType.NumField(); i++ {
			v := reflectType.Field(i)
			if v.Anonymous {
				fields = append(fields, deepFields(v.Type)...)
			} else {
				fields = append(fields, v)
			}
		}
	}

	return fields
}

func change(a interface{}) {
	rv := reflect.ValueOf(a)
	changerv(rv)
}

func changerv(rv reflect.Value) {
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Struct {
		changeStruct(rv)
	}
	if rv.Kind() == reflect.Slice {
		changeSlice(rv)
	}
}

// assumes rv is a slice
func changeSlice(rv reflect.Value) {
	ln := rv.Len()
	if ln == 0 && rv.CanAddr() {
		var elem reflect.Value

		typ := rv.Type().Elem()
		if typ.Kind() == reflect.Ptr {
			elem = reflect.New(typ.Elem())
		}
		if typ.Kind() == reflect.Struct {
			elem = reflect.New(typ).Elem()
		}

		rv.Set(reflect.Append(rv, elem))
	}

	ln = rv.Len()
	for i := 0; i < ln; i++ {
		changerv(rv.Index(i))
	}
}

// assumes rv is a struct
func changeStruct(rv reflect.Value) {
	if !rv.CanAddr() {
		return
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		switch field.Kind() {
		case reflect.String:
			field.SetString("fred")
		case reflect.Int:
			field.SetInt(54)
		default:
			fmt.Println("unknown field")
		}
	}
}
