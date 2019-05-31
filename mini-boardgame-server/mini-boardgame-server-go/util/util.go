package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func StructContain(big interface{}, small interface{}) (bool, error) {
	ok, err := structContain("", big, small)
	if err == "" {
		return ok, nil
	} else {
		return ok, errors.New(err)
	}
}

func structContain(prefix string, big interface{}, small interface{}) (ok bool, err string) {
	defer func() {
		if err != "" && prefix != "" {
			err = prefix + "->" + err
		}
	}()
	if small == nil || big == nil {
		if small == nil && big == nil {
			return true, ""
		} else if small == nil {
			return false, "Expect null but actual not"
		} else {
			return false, "Expect non-null but actual null"
		}
	}
	bv := reflect.ValueOf(big)
	sv := reflect.ValueOf(small)
	if bv.Kind() != sv.Kind() {
		return false, fmt.Sprintf("Expect %T but %T", small, big)
	}
	switch sv.Kind() {
	case reflect.Struct:
		bt := bv.Type()
		st := sv.Type()
		for i := 0; i < st.NumField(); i++ {
			sf := st.Field(i)
			_, ok := bt.FieldByName(sf.Name)
			if !ok {
				return false, fmt.Sprintf("Expect field %s but not present", sf.Name)
			}
			bvf := bv.FieldByName(sf.Name)
			svf := sv.FieldByName(sf.Name)
			if ok, err := structContain(sf.Name, bvf.Interface(), svf.Interface()); !ok {
				return false, err
			}
		}
		return true, ""
	case reflect.Ptr:
		return structContain("", bv.Elem().Interface(), sv.Elem().Interface())
	case reflect.Map:
		for _, sk := range sv.MapKeys() {
			var bmv *reflect.Value
			for _, bk := range bv.MapKeys() {
				if sk.Interface() == bk.Interface() {
					tmp := bv.MapIndex(bk)
					bmv = &tmp
					break
				}
			}
			if bmv == nil {
				return false, fmt.Sprintf("Expect key %s but not present", sk.String())
			}
			smv := sv.MapIndex(sk)
			if ok, err := structContain(sk.String(), bmv.Interface(), smv.Interface()); !ok {
				return false, err
			}
		}
		return true, ""
	case reflect.Array, reflect.Slice:
		if sv.Len() != bv.Len() {
			return false, fmt.Sprintf("Expect len %d but %d", sv.Len(), bv.Len())
		}
		for i := 0; i < sv.Len(); i++ {
			if ok, err := structContain(strconv.Itoa(i), bv.Index(i).Interface(), sv.Index(i).Interface()); !ok {
				return false, err
			}
		}
		return true, ""
	default:
		if big == small {
			return true, ""
		} else {
			return false, fmt.Sprintf("Expect %v but %v", small, big)
		}
	}
}
