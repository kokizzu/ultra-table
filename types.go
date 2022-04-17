package ultra_table

import (
	"errors"
	"fmt"
	"unsafe"
)

type IRow interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type tag struct {
	offset    uintptr
	fieldType FieldType
	indexType IndexType
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

type FieldType string

const (
	String     FieldType = `string`
	Int        FieldType = `int`
	Int8       FieldType = `int8`
	Int16      FieldType = `int16`
	Int32      FieldType = `int32`
	Int64      FieldType = `int64`
	Uint       FieldType = `uint`
	Uint8      FieldType = `uint8`
	Uint16     FieldType = `uint16`
	Uint32     FieldType = `uint32`
	Uint64     FieldType = `uint64`
	Float32    FieldType = `float32`
	Float64    FieldType = `float64`
	Complex64  FieldType = `complex64`
	Complex128 FieldType = `complex128`
	Byte       FieldType = `byte`
	Rune       FieldType = `rune`
)

type IndexType string

const (
	UNIQUE IndexType = `unique`
	NORMAL IndexType = `normal`
)

func getIndex(inputIndex string) (IndexType, error) {
	switch IndexType(inputIndex) {
	case UNIQUE:
		return UNIQUE, nil
	case NORMAL:
		return NORMAL, nil
	}
	return ``, fmt.Errorf(`index type not support`)
}

func (tag *tag) GetPointerVal(p unsafe.Pointer) interface{} {
	var val interface{}
	switch tag.fieldType {
	case String:
		val = *(*string)(p)
	case Int:
		val = *(*int)(p)
	case Int8:
		val = *(*int8)(p)
	case Int16:
		val = *(*int16)(p)
	case Int32:
		val = *(*int32)(p)
	case Int64:
		val = *(*int64)(p)
	case Uint:
		val = *(*uint)(p)
	case Uint8:
		val = *(*uint8)(p)
	case Uint16:
		val = *(*uint16)(p)
	case Uint32:
		val = *(*uint32)(p)
	case Uint64:
		val = *(*uint64)(p)
	case Float32:
		val = *(*float32)(p)
	case Float64:
		val = *(*float64)(p)
	case Complex64:
		val = *(*complex64)(p)
	case Complex128:
		val = *(*complex128)(p)
	case Byte:
		val = *(*byte)(p)
	case Rune:
		val = *(*rune)(p)
	}
	return val
}

func (tag *tag) CheckIsNormal() bool {
	return tag.indexType == NORMAL
}

func (tag *tag) CheckIsUnique() bool {
	return tag.indexType == UNIQUE
}

func GetTag(dest interface{}, offset uintptr, indexType IndexType) (*tag, error) {
	if dest == nil {
		return nil, errors.New(`unsupported index type`)
	}
	switch dest.(type) {
	case string:
		return getStringTag(offset, indexType), nil
	case int:
		return getIntTag(offset, indexType), nil
	case int8:
		return getInt8Tag(offset, indexType), nil
	case int16:
		return getInt16Tag(offset, indexType), nil
	case int32:
		return getInt32Tag(offset, indexType), nil
	case int64:
		return getInt64Tag(offset, indexType), nil
	case uint:
		return getUintTag(offset, indexType), nil
	case uint8:
		return getUint8Tag(offset, indexType), nil
	case uint16:
		return getUint16Tag(offset, indexType), nil
	case uint32:
		return getUint32Tag(offset, indexType), nil
	case uint64:
		return getUint64Tag(offset, indexType), nil
	case float32:
		return getFloat32Tag(offset, indexType), nil
	case float64:
		return getFloat64Tag(offset, indexType), nil
	case complex64:
		return getComplex64Tag(offset, indexType), nil
	case complex128:
		return getComplex128Tag(offset, indexType), nil
	}
	return nil, errors.New(`unsupported index type`)
}

func getStringTag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: String,
		indexType: indexType,
	}
}
func getIntTag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Int,
		indexType: indexType,
	}
}
func getInt8Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Int8,
		indexType: indexType,
	}
}
func getInt16Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Int16,
		indexType: indexType,
	}
}
func getInt32Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Int32,
		indexType: indexType,
	}
}
func getInt64Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Int64,
		indexType: indexType,
	}
}
func getUintTag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Uint,
		indexType: indexType,
	}
}
func getUint8Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Uint8,
		indexType: indexType,
	}
}
func getUint16Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Uint16,
		indexType: indexType,
	}
}
func getUint32Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Uint32,
		indexType: indexType,
	}
}
func getUint64Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Uint64,
		indexType: indexType,
	}
}
func getFloat32Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Float32,
		indexType: indexType,
	}
}
func getFloat64Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Float64,
		indexType: indexType,
	}
}
func getComplex64Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Complex64,
		indexType: indexType,
	}
}
func getComplex128Tag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Complex128,
		indexType: indexType,
	}
}
func getByteTag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Byte,
		indexType: indexType,
	}
}
func getRuneTag(offset uintptr, indexType IndexType) *tag {
	return &tag{
		offset:    offset,
		fieldType: Rune,
		indexType: indexType,
	}
}
