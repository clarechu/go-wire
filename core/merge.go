package core

import (
	"errors"
	"fmt"
	"reflect"
)

type Mold func(v reflect.Value) bool

func Merge(to interface{}, from interface{}) error {
	toVal := reflect.ValueOf(to)
	fromVal := reflect.ValueOf(from)
	if !isPtr(toVal) || !isPtr(fromVal) {
		return errors.New("merge to value is unaddressable")
	}
	mergeValue(toVal, fromVal, isNil)
	return nil
}

func Merges(to interface{}, from interface{}, mold Mold) error {
	toVal := reflect.ValueOf(to)
	fromVal := reflect.ValueOf(from)
	if !isPtr(toVal) || !isPtr(fromVal) {
		return errors.New("merge to value is unaddressable")
	}
	mergeValue(toVal, fromVal, mold)
	return nil
}

func mergeValue(to reflect.Value, from reflect.Value, mold Mold) {
	tVa := indirect(to)
	tt := tVa.Type()
	fVa := indirect(from)
	ft := fVa.Type()
	if tVa.Kind() == reflect.Slice {
		if fVa.Kind() == reflect.Slice {
			mergeSlice(tVa, fVa, mold)
			return
		}
	}
	for j := 0; j < tVa.NumField(); j++ {
		tField := tt.Field(j)
		tKind := tVa.Field(j).Kind()
		for i := 0; i < fVa.NumField(); i++ {
			fField := ft.Field(i)
			fKind := fVa.Field(i).Kind()
			switch fKind {
			case reflect.Map:
				if tField.Name == fField.Name && tKind == reflect.Map {
					if mold(fVa.FieldByName(fField.Name)) {
						break
					}
					MergeMap(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j))
					break
				}
			case reflect.Bool:
				if tKind == reflect.Bool && tField.Name == fField.Name {
					/*					if !isZero(fVa.FieldByName(fField.Name)) {
										break
									}*/
					tVa.Field(j).SetBool(fVa.FieldByName(fField.Name).Bool())
					break
				}
			case reflect.Ptr:
				// to and from is Struct
				if tKind == reflect.Ptr && tField.Name == fField.Name {
					if isZero(fVa.FieldByName(fField.Name)) {
						break
					}
					mergeValue(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j), mold)
					break
				}
			case reflect.Array:
				if tKind == fKind && tField.Name == fField.Name {
					mergeArray(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j))
					break
				}
				break
			case reflect.Slice:
				if tKind == fKind && tField.Name == fField.Name {
					mergeSlice(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j), mold)
					break
				}
				break
			case reflect.Struct:
				// to and from is Struct
				if tKind == reflect.Struct && tField.Name == fField.Name {
					fmt.Println(tKind, "...", fKind)
					fmt.Println(tField.Name, "...", fField.Name)
					if isZero(fVa.FieldByName(fField.Name)) {
						break
					}
					mergeValue(reflect.Indirect(to).Field(i), reflect.Indirect(from).Field(j), mold)
					break

				}
			default:
				if tKind == fKind && tField.Name == fField.Name {
					if isZero(fVa.FieldByName(fField.Name)) {
						break
					} else {
						tVa.Field(j).Set(fVa.FieldByName(fField.Name))
						break
					}
				}
			}
		}
	}

}

func MergeMap(to reflect.Value, from reflect.Value) {
	fmt.Println("len = ", from.Len())
	if to.IsNil() {
		mapValue := reflect.MakeMap(from.Type())
		to.Set(mapValue)
		for key, element := range from.MapKeys() {
			fmt.Println(key, element) // how to get the value?
			v := from.MapIndex(element)
			//mapType := reflect.MapOf(reflect.TypeOf(element), reflect.TypeOf(from.MapIndex(element)))
			to.SetMapIndex(element, v)
		}
	} else {
		for key, element := range from.MapKeys() {
			fmt.Println(key, element) // how to get the value?
			v := from.MapIndex(element)
			//mapType := reflect.MapOf(reflect.TypeOf(element), reflect.TypeOf(from.MapIndex(element)))
			to.SetMapIndex(element, v)
		}
	}
	//to.Set(reflect.MakeMap(from.Type()))
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func:
		return v.IsNil()
	case reflect.Map, reflect.Slice:
		if v.IsNil() || v.Len() == 0 {
			return true
		}
		return false
	case reflect.Bool:
		return v.Bool()
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
	default:
		z := reflect.Zero(v.Type())
		return v.Interface() == z.Interface()
	}
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Bool:
		fmt.Println(v.Bool())
		return v.Bool()
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
	default:
		z := reflect.Zero(v.Type())
		return v.Interface() == z.Interface()
	}
}

func mergeSlice(to reflect.Value, from reflect.Value, mold Mold) {
	if from.IsZero() {
		return
	}
	toLen := to.Len()
	for i := 0; i < from.Len(); i++ {
		j := 0
		if j < toLen {
			if to.Index(j).Kind() == reflect.Ptr || to.Index(j).Kind() == reflect.Struct {
				mergeValue(to.Index(j), from.Index(i), mold)
				continue
			}
			to.Set(reflect.Append(to, from.Index(i)))
			j++
			continue
		} else {
			appendSlice(to, from, i, mold)
		}
	}

}

func appendSlice(to reflect.Value, from reflect.Value, i int, mold Mold) {
	if to.IsNil() {
		mapValue := reflect.MakeSlice(from.Type(), 0, 0)
		to.Set(mapValue)
	}
	var elem reflect.Value
	typ := to.Type().Elem()
	if typ.Kind() == reflect.Ptr {
		elem = reflect.New(typ.Elem())
		mergeValue(elem, from.Index(i), mold)
		to.Set(reflect.Append(to, elem))
		return
	}
	if typ.Kind() == reflect.Struct {
		elem = reflect.New(typ).Elem()
		mergeValue(elem, from.Index(i), mold)
		to.Set(reflect.Append(to, elem))
		return
	}
	to.Set(reflect.Append(to, from.Index(i)))
}

func mergeArray(to reflect.Value, from reflect.Value) {}
