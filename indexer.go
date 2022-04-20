package ultra_table

import (
	"reflect"
	"unsafe"
)

type fieldIndexer struct {
	indexTags        map[string]*tag
	normalIndexItems map[string]map[interface{}]*BitMap //normal index
	uniqueIndexItems map[string]map[interface{}]uint32  //unique index
}

func newFieldIndexer() *fieldIndexer {
	return &fieldIndexer{
		indexTags:        map[string]*tag{},
		normalIndexItems: map[string]map[interface{}]*BitMap{},
		uniqueIndexItems: map[string]map[interface{}]uint32{},
	}
}

func (f *fieldIndexer) getIndexTag(idxKey string) (*tag, error) {
	tag, ok := f.indexTags[idxKey]

	if !ok {
		return nil, RecordNotFound
	}
	return tag, nil
}

func (f *fieldIndexer) IndexTagLen() int {
	return len(f.indexTags)
}

func (f *fieldIndexer) buildIndex(item interface{}, idx uint32) error {
	if len(f.indexTags) == 0 {
		_value := reflect.ValueOf(item)
		if _value.Kind() != reflect.Ptr {
			return OnlySupportPtr
		}
		for i := 0; i < _value.Elem().NumField(); i++ {
			tagStr := _value.Elem().Type().Field(i).Tag.Get(`idx`)
			if tagStr == "" || tagStr == "-" {
				continue
			}
			indexType, err := getIndex(tagStr)
			if err != nil {
				return err
			}
			name := _value.Elem().Type().Field(i).Name
			offset := _value.Elem().Type().Field(i).Offset

			t, err := GetTag(_value.Elem().Field(i).Interface(), offset, indexType)
			if err != nil {
				if f.IndexTagLen() > 0 {
					f.indexTags = make(map[string]*tag, 0)
				}
				return err
			}
			f.indexTags[name] = t
		}
	}
	//if has not index, return
	if f.IndexTagLen() == 0 {
		return nil
	}
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&item)).word)
	for name, tag := range f.indexTags {
		//pre check unique index
		if tag.CheckIsUnique() {
			if m, ok := f.uniqueIndexItems[name]; ok {
				val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
				if _, ok := m[val]; ok {
					return UniqueIndex
				}
			}
		}
	}

	for name, tag := range f.indexTags {
		val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
		if tag.CheckIsNormal() {
			m, ok := f.normalIndexItems[name]
			if !ok {
				f.normalIndexItems[name] = make(map[interface{}]*BitMap)
				f.normalIndexItems[name][val] = NewBitMap()
				f.normalIndexItems[name][val].Add(idx)
			} else {
				if m[val] == nil {
					m[val] = NewBitMap()
				}
				m[val].Add(idx)
			}
		} else if tag.CheckIsUnique() {
			m, ok := f.uniqueIndexItems[name]
			if !ok {
				f.uniqueIndexItems[name] = map[interface{}]uint32{}
				f.uniqueIndexItems[name][val] = idx
			} else {
				m[val] = idx
			}
		}
	}
	return nil
}

func (f *fieldIndexer) isExistUnique() bool {
	for _, v := range f.indexTags {
		if v.CheckIsUnique() {
			return true
		}
	}
	return false
}

func (f *fieldIndexer) removeIndex(idx uint32, item interface{}) {
	for name, tag := range f.indexTags {

		if tag.CheckIsUnique() {
			m, ok := f.uniqueIndexItems[name]
			if ok {
				ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&item)).word)
				val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
				delete(m, val)
			}
		}

		if tag.CheckIsNormal() {
			m, ok := f.normalIndexItems[name]
			if ok {
				ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&item)).word)
				val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
				m[val].Remove(idx)
				if m[val].IsEmpty() {
					delete(m, val)
				}
			}
		}

	}
}

func (f *fieldIndexer) GetIdxBitMap(idxKey string, vKey interface{}) (*BitMap, error) {
	tag, err := f.getIndexTag(idxKey)
	if err != nil {
		return nil, err
	}
	if tag.CheckIsNormal() {
		index, ok := f.normalIndexItems[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		bitmap, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		return bitmap, nil
	}

	if tag.CheckIsUnique() {
		index, ok := f.uniqueIndexItems[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		idx, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		bitmap := NewBitMap()
		bitmap.Add(idx)
		return bitmap, nil
	}

	return nil, RecordNotFound
}

func (f *fieldIndexer) IsExist(idxKey string, vKey interface{}) bool {
	tag, err := f.getIndexTag(idxKey)
	if err != nil {
		return false
	}
	if tag.CheckIsNormal() {
		index, ok := f.normalIndexItems[idxKey]
		if !ok {
			return false
		}
		_, ok = index[vKey]
		return ok
	}
	if tag.CheckIsUnique() {
		index, ok := f.uniqueIndexItems[idxKey]
		if !ok {
			return false
		}
		_, ok = index[vKey]
		return ok
	}

	return false
}
