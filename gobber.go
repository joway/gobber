package gobber

import (
	"reflect"
	"unsafe"
)

type Robber struct {
	typ   reflect.Type
	state map[string]*reflect.StructField //map[filedName]*structField
}

func New(template interface{}) *Robber {
	typ := reflect.TypeOf(template)
	return &Robber{
		typ:   typ,
		state: map[string]*reflect.StructField{},
	}
}

func (r *Robber) getStructField(fieldName string) (*reflect.StructField, bool) {
	sf, ok := r.state[fieldName]
	if ok {
		return sf, true
	}

	typ := r.typ
	var field reflect.StructField
	if typ.Kind() == reflect.Ptr {
		field, ok = typ.Elem().FieldByName(fieldName)
	} else {
		field, ok = typ.FieldByName(fieldName)
	}
	if !ok {
		return nil, false
	}

	r.state[fieldName] = &field
	return &field, true
}

func (r *Robber) getFieldPtr(target interface{}, offset uintptr) unsafe.Pointer {
	structPtr := ifaceToPtr(target)
	return unsafe.Pointer(uintptr(structPtr) + offset)
}

func (r *Robber) Get(target interface{}, fieldName string) unsafe.Pointer {
	structField, ok := r.getStructField(fieldName)
	if !ok {
		return nil
	}
	return r.getFieldPtr(target, structField.Offset)
}

func (r *Robber) Set(target interface{}, fieldName string, fieldValue interface{}) bool {
	structField, ok := r.getStructField(fieldName)
	if !ok {
		return false
	}
	toValue := reflect.ValueOf(fieldValue)
	if structField.Type != toValue.Type() {
		return false
	}

	fieldPtr := r.getFieldPtr(target, structField.Offset)
	reflect.NewAt(structField.Type, fieldPtr).Elem().Set(toValue)
	return true
}
