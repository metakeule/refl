package refl

import (
	"fmt"
	"reflect"
)

func Inspect(i interface{}) (s string) {
	if Kind(i) == "float64" {
		s = fmt.Sprintf("%f (%s)", i, Type(i))
	} else {
		s = fmt.Sprintf("%#v (%s)", i, Type(i))
	}
	return
}

func Apply(a interface{}, meth string, b interface{}) {
	x := reflect.ValueOf(a).Interface()
	y := reflect.ValueOf(b).Interface()
	//refl.SetField(ps, "Inner", t)
	//p(refl.Inspect(refl.GetField(ps, "Inner")))
	Call(&x, meth, &y)
}

// returns the type
func Type(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

// returns the kind
func Kind(i interface{}) string {
	return reflect.TypeOf(i).Kind().String()
}

// calls a function with vals, but doesn't return anything
func Call(ø interface{}, meth string, vals ...interface{}) {
	m := reflect.ValueOf(ø).MethodByName(meth)
	if !m.IsValid() {
		panic("can't find method " + meth + " of " + Inspect(ø))
	}
	params := []reflect.Value{}
	for i := range vals {
		if vals[i] != nil {
			params = append(params, reflect.ValueOf(vals[i]))
		}
	}
	m.Call(params)
}

// get an attribute of a struct
func GetField(ø interface{}, field string) interface{} {
	p := reflect.ValueOf(ø).Elem().FieldByName(field)
	return p.Interface()
}

// set an attribute of a struct
func SetField(ø interface{}, field string, val interface{}) {
	p := reflect.ValueOf(ø).Elem().FieldByName(field)
	if p.CanSet() {
		p.Set(reflect.ValueOf(val))
	} else {
		panic("can't set field " + field + " of " + Inspect(ø))
	}
}

// sets a value
func Set(ø interface{}, val interface{}) {
	p := reflect.ValueOf(ø)
	if p.Elem().CanSet() {
		p.Elem().Set(reflect.ValueOf(val))
	} else {
		panic("can't set " + Inspect(ø))
	}
}
