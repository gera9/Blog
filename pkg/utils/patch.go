package utils

import (
	"errors"
	"reflect"
)

var (
	ErrSrcNotStruct        = errors.New("source is not a struct")
	ErrDstNotPointer       = errors.New("destination is not a pointer")
	ErrDstNotPointerStruct = errors.New("destination is not a pointer to struct")
	ErrDiffTypes           = errors.New("source and destination types do not match")
)

func PatchStruct(dst, src any) error {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Struct {
		return ErrSrcNotStruct
	}

	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Pointer || dstVal.IsNil() {
		return ErrDstNotPointer
	}

	srcType := srcVal.Type()
	dstElem := dstVal.Elem()

	dstType := dstElem.Type()
	if dstType.Kind() != reflect.Struct {
		return ErrDstNotPointerStruct
	}

	if dstType != srcType {
		return ErrDiffTypes
	}

	for i := range srcType.NumField() {
		srcType := srcType.Field(i)
		if !srcType.IsExported() {
			continue
		}

		srcFieldVal := srcVal.Field(i)
		if srcFieldVal.IsZero() {
			continue
		}

		dstFieldVal := dstElem.Field(i)

		// Handle pointer nil fields
		if srcType.Type.Kind() == reflect.Pointer && dstFieldVal.IsNil() {
			// If dst is nil, set it to src
			newVal := reflect.New(srcType.Type.Elem())
			newVal.Elem().Set(srcFieldVal.Elem())
			dstFieldVal.Set(newVal)
			continue
		}

		// Handle nested structs recursively (except for time.Time)
		if srcType.Type.Kind() == reflect.Struct && srcType.Type.PkgPath() != "time" {
			PatchStruct(dstFieldVal.Addr().Interface(), srcFieldVal.Interface())
			continue
		}

		// Remaining types set directly if different
		if !reflect.DeepEqual(dstFieldVal.Interface(), srcFieldVal.Interface()) {
			dstFieldVal.Set(srcFieldVal)
		}
	}

	return nil
}
