package axapi

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func SerializeToJson(a interface{}) ([]byte, error) {
	t1 := reflect.TypeOf(a)
        var t reflect.Type
        var v reflect.Value
        if (t1.Kind() == reflect.Ptr) {
                t = t1.Elem()
                v = reflect.ValueOf(a).Elem()
        } else {
                t = t1
                v = reflect.ValueOf(a)
        }
	if v.IsZero() {
		return nil, nil
	}

	payloads := make([]string, 0, 1024)
	for i := 0; i < v.NumField(); i++ {
		vField := v.Field(i)
		tField := t.Field(i)
		jstag := tField.Tag.Get("json")
		if len(jstag) == 0 { //not a json field. omit it
			continue
		}
		dValue := tField.Tag.Get("dval")
		omitFlag := false
		vKind := vField.Kind()

		if vField.IsZero() && len(dValue) == 0 { //omit zore value, if it has no default value (default value may be non zore value)
			continue
		}
		payload := "\"" + jstag + "\":"

		switch vKind {
		case reflect.Struct:
			x, err := SerializeToJson(vField.Interface())
			if err != nil {
				return nil, err
			}
			if x == nil {
				if tField.Name == "Inst" { //special case
					payload += "{}"
				} else {
					omitFlag = true
				}
			} else {
				payload += string(x)
			}
			break
		case reflect.Int:
			tmpValue := vField.Int()
			dValueInt, _ := strconv.ParseInt(dValue, 10, 64)
			if tmpValue == dValueInt {
				omitFlag = true
			} else {
				payload += strconv.FormatInt(tmpValue, 10)
			}
			break
		case reflect.String:
			dValue := tField.Tag.Get("dval")
			tmpValue := vField.String()
			if tmpValue == dValue {
				omitFlag = true
			} else {
				payload += "\"" + tmpValue + "\""
			}
			break
		case reflect.Slice:
			s_elems := make([]string, 0, 1024)
			for i := 0; i < vField.Len(); i++ {
				x, err := SerializeToJson(vField.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				s_elems = append(s_elems, string(x))
			}
			if len(s_elems) == 0 {
				omitFlag = true
			} else {
				payload += "[" + strings.Join(s_elems, ",") + "]"
			}
			break
		default:
			return nil, errors.New("Unsupport type: " + vKind.String())
		}
		if omitFlag == false {
			payloads = append(payloads, payload)
		}
	}

	if len(payloads) > 0 {
		ret := make([]byte, 0, 1024)
		ret = append(ret, 123)
		ret = append(ret, strings.Join(payloads, ",")...)
		ret = append(ret, 125)
		return ret, nil
	}
	return nil, nil
}
