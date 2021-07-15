package ultra_table

import (
	"errors"
	"unsafe"

	"reflect"
	"sync"
)

var (
	RecordNotFound    = errors.New("record not found")
	OnlySupportStruct = errors.New("only support struct")
	ValueNotNil       = errors.New("value can not be nil")
)

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}
type IndexGroup struct {
	indexItems map[string]map[interface{}]*BitMap
	indexTags  map[string]*tag
}

func (indexGroup *IndexGroup) IndexTagLen() int {
	return len(indexGroup.indexTags)
}

type ItemIterator func(interface{}) bool

type UltraTable struct {
	mu         sync.RWMutex
	table      []interface{}
	indexGroup IndexGroup
	emptyMap   *BitMap
}

func NewUltraTable() *UltraTable {
	return &UltraTable{
		table:      make([]interface{}, 0),
		indexGroup: IndexGroup{indexItems: map[string]map[interface{}]*BitMap{}, indexTags: map[string]*tag{}},
		emptyMap:   NewBitMap(),
	}
}

func (u *UltraTable) Lock() {
	u.mu.Lock()
}
func (u *UltraTable) Unlock() {
	u.mu.Unlock()
}
func (u *UltraTable) Add(dest interface{}) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.add(dest)
}

//Get benchmark performance near O(n), it is recommended to use GetWithIdx
func (u *UltraTable) Remove(iterator ItemIterator) uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.remove(iterator)
}

//RemoveWithIdx benchmark performance near O(1)
func (u *UltraTable) RemoveWithIdx(idxKey string, vKey interface{}) uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.removeWithIdx(idxKey, vKey)
}

func (u *UltraTable) UpdateWithIdx(idxKey string, vKey interface{}, newDest interface{}) uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.updateWithIdx(idxKey, vKey, newDest)
}

//SaveWithIdx update or insert
func (u *UltraTable) SaveWithIdx(idxKey string, vKey interface{}, newDest interface{}) uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.hasWithIdx(idxKey, vKey) {
		return u.updateWithIdx(idxKey, vKey, newDest)
	}
	err := u.add(newDest)
	if err != nil {
		return 0
	} else {
		return 1
	}
}

func (u *UltraTable) SaveWithIdxAggregate(conditions map[string]interface{}, newDest interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.saveWithIdxAggregate(conditions, newDest)
}

func (u *UltraTable) SaveWithIdxIntersection(conditions map[string]interface{}, newDest interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.saveWithIdxIntersection(conditions, newDest)
}

//Get benchmark performance near O(1)
func (u *UltraTable) GetWithIdx(idxKey string, vKey interface{}) ([]interface{}, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.getWithIdx(idxKey, vKey)
}

//GetWithIdxAggregate like where a=? or b=?
func (u *UltraTable) GetWithIdxAggregate(conditions map[string]interface{}) ([]interface{}, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.getWithIdxAggregate(conditions)
}

//GetWithIdxIntersection like where a=? and b=?
func (u *UltraTable) GetWithIdxIntersection(conditions map[string]interface{}) ([]interface{}, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.getWithIdxIntersection(conditions)
}

//Get benchmark performance near O(n), it is recommended to use GetWithIdx
func (u *UltraTable) Get(iterator ItemIterator) []interface{} {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.get(iterator)
}

func (u *UltraTable) GetAll() []interface{} {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.getAll()
}

func (u *UltraTable) Clear() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.clear()
}

func (u *UltraTable) Len() uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.len()
}

func (u *UltraTable) Cap() uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return uint64(len(u.table))
}

func (u *UltraTable) Has(iterator ItemIterator) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.has(iterator)
}

func (u *UltraTable) HasWithIdx(idxKey string, vKey interface{}) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.hasWithIdx(idxKey, vKey)
}

func (u *UltraTable) GetWithIdxCount(idxKey string, vKey interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.getWithIdxCount(idxKey, vKey)
}
func (u *UltraTable) GetWithIdxAggregateCount(conditions map[string]interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.getWithIdxAggregateCount(conditions)
}

func (u *UltraTable) GetWithIdxIntersectionCount(conditions map[string]interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.getWithIdxIntersectionCount(conditions)
}

func (u *UltraTable) getWithIdxCount(idxKey string, vKey interface{}) uint64 {

	index, ok := u.indexGroup.indexItems[idxKey]
	if !ok {
		return 0
	}
	bitMap, ok := index[vKey]
	if !ok {
		return 0
	}
	return bitMap.Length()
}

func (u *UltraTable) getWithIdxIntersectionCount(conditions map[string]interface{}) uint64 {
	intersectionList := []*BitMap{}
	minLen := 0
	minLenIndex := 0
	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			return 0
		}
		sliceList, ok := index[vKey]
		if !ok {
			return 0
		}
		intersectionList = append(intersectionList, sliceList)
		if len(intersectionList) > 1 {
			if int(sliceList.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(sliceList.Length())
			}
		} else {
			minLen = int(sliceList.Length())
			minLenIndex = 0
		}
	}
	if len(intersectionList) == 0 {
		return 0
	}

	tempMap := map[uint32]uint64{}

	intersectionList[minLenIndex].Iterator(func(k uint32) {
		tempMap[k] = 1
		for i := 0; i < len(intersectionList); i++ {
			if i == minLenIndex {
				continue
			}
			if ok := intersectionList[i].IsExist(k); ok {
				tempMap[k] += 1
			}
		}
	})
	count := 0
	for _, v := range tempMap {
		if v == uint64(len(intersectionList)) {
			count++
		}
	}
	return uint64(count)
}

func (u *UltraTable) getWithIdxAggregateCount(conditions map[string]interface{}) uint64 {

	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			continue
		}
		sliceList, ok := index[vKey]
		if !ok {
			continue
		}
		aggregateList = append(aggregateList, sliceList)
	}
	if len(aggregateList) == 0 {
		return 0
	}

	tempMap := map[uint32]uint8{}

	for _, aggregateSlice := range aggregateList {
		aggregateSlice.Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}

	return uint64(len(tempMap))
}

func (u *UltraTable) removeIndex(idx uint32, dest interface{}) {
	for name, tag := range u.indexGroup.indexTags {

		m, ok := u.indexGroup.indexItems[name]
		if ok {
			ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&dest)).word)
			val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
			m[val].Remove(idx)
		}
	}
}

func (u *UltraTable) addIndex(dest interface{}, idx uint32) error {
	if dest == nil {
		return ValueNotNil
	}

	if len(u.indexGroup.indexTags) == 0 {
		_value := reflect.ValueOf(dest)
		if _value.Kind() != reflect.Struct {
			return OnlySupportStruct
		}
		for i := 0; i < _value.NumField(); i++ {
			tagStr := _value.Type().Field(i).Tag.Get(`idx`)
			if tagStr == "" || tagStr == "-" {
				continue
			}
			indexType, err := getIndex(tagStr)
			if err != nil {
				return err
			}
			name := _value.Type().Field(i).Name
			offset := _value.Type().Field(i).Offset

			t, err := GetTag(_value.Field(i).Interface(), offset, indexType)
			if err != nil {
				if u.indexGroup.IndexTagLen() > 0 {
					u.indexGroup.indexTags = make(map[string]*tag, 0)
				}
				return err
			}
			u.indexGroup.indexTags[name] = t
		}
	}
	if u.indexGroup.IndexTagLen() == 0 {
		return nil
	}
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&dest)).word)

	for name, tag := range u.indexGroup.indexTags {
		val := tag.GetPointerVal(unsafe.Pointer(ptr0 + tag.offset))
		m, ok := u.indexGroup.indexItems[name]
		if !ok {
			u.indexGroup.indexItems[name] = make(map[interface{}]*BitMap)
			u.indexGroup.indexItems[name][val] = NewBitMap()
			u.indexGroup.indexItems[name][val].Add(idx)
		} else {
			if m[val] == nil {
				m[val] = NewBitMap()
			}
			m[val].Add(idx)
		}
	}
	return nil
}

func (u *UltraTable) hasWithIdx(idxKey string, vKey interface{}) bool {
	index, ok := u.indexGroup.indexItems[idxKey]
	if !ok {
		return false
	}
	_, ok = index[vKey]
	return ok
}

func (u *UltraTable) len() uint64 {
	return uint64(len(u.table) - int(u.emptyMap.Length()))
}

func (u *UltraTable) has(iterator ItemIterator) bool {
	for i := 0; i < len(u.table); i++ {
		if u.table[i] == nil {
			continue
		}
		if iterator(u.table[i]) {
			return true
		}
	}
	return false
}

func (u *UltraTable) add(dest interface{}) error {

	if u.emptyMap.Length() == 0 {
		err := u.addIndex(dest, uint32(len(u.table)))
		if err != nil {
			return err
		}
		u.table = append(u.table, dest)
	} else {
		i := u.emptyMap.Min()
		err := u.addIndex(dest, i)
		if err != nil {
			return err
		}
		u.table[i] = dest
		u.emptyMap.Remove(i)
		return nil
	}
	return nil
}

func (u *UltraTable) remove(iterator ItemIterator) uint64 {
	var count uint64
	for i := 0; i < len(u.table); i++ {
		if u.table[i] == nil {
			continue
		}
		if iterator(u.table[i]) {
			u.removeIndex(uint32(i), u.table[i])
			u.table[i] = nil
			u.emptyMap.Add(uint32(i))
			count++
		}
	}
	return count
}

func (u *UltraTable) addWithTransition(dest interface{}) (uint64, error) {
	if u.emptyMap.Length() == 0 {
		err := u.addIndex(dest, uint32(len(u.table)))
		if err != nil {
			return 0, err
		}
		u.table = append(u.table, dest)
		return u.len() - 1, nil
	} else {
		i := u.emptyMap.Min()
		err := u.addIndex(dest, i)
		if err != nil {
			return 0, err
		}
		u.table[i] = dest
		u.emptyMap.Remove(i)
		return uint64(i), nil
	}
}

func (u *UltraTable) removeWithTransition(arryIndex uint64) {

	if u.table[arryIndex] == nil {
		return
	}
	u.removeIndex(uint32(arryIndex), u.table[arryIndex])
	u.table[arryIndex] = nil
	u.emptyMap.Add(uint32(arryIndex))
}

func (u *UltraTable) removeWithIdx(idxKey string, vKey interface{}) uint64 {
	index, ok := u.indexGroup.indexItems[idxKey]
	if !ok {
		return 0
	}
	_, ok = index[vKey]
	if !ok {
		return 0
	}
	count := uint64(0)
	index[vKey].CloneIterator(func(k uint32) {
		u.removeIndex(k, u.table[k])
		u.table[k] = nil
		u.emptyMap.Add(k)
		count++
	})
	return count
}

func (u *UltraTable) updateWithIdx(idxKey string, vKey interface{}, newDest interface{}) uint64 {

	index, ok := u.indexGroup.indexItems[idxKey]
	if !ok {
		return 0
	}
	_, ok = index[vKey]
	if !ok {
		return 0
	}
	count := 0
	index[vKey].CloneIterator(func(k uint32) {
		u.removeIndex(uint32(k), u.table[k])
		u.table[k] = nil
		u.emptyMap.Add(k)
		count++
	})

	for i := 0; i <= count-1; i++ {
		j := u.emptyMap.Min()
		u.table[j] = newDest
		u.addIndex(newDest, j)
		u.emptyMap.Remove(j)
	}
	return uint64(count)
}

func (u *UltraTable) getWithIdx(idxKey string, vKey interface{}) ([]interface{}, error) {
	index, ok := u.indexGroup.indexItems[idxKey]
	if !ok {
		return nil, RecordNotFound
	}
	sliceList, ok := index[vKey]
	if !ok {
		return nil, RecordNotFound
	}
	result := make([]interface{}, sliceList.Length())
	count := 0
	sliceList.Iterator(func(k uint32) {
		result[count] = u.table[k]
		count++
	})
	return result, nil
}

func (u *UltraTable) clear() {
	u.table = make([]interface{}, 0)
	u.indexGroup = IndexGroup{indexItems: map[string]map[interface{}]*BitMap{}}
	u.emptyMap.Clear()
}

func (u *UltraTable) getAll() []interface{} {
	emptyInc := 0
	result := make([]interface{}, u.Len())
	for i := 0; i < len(u.table); i++ {
		if u.table[i] == nil {
			emptyInc++
			continue
		}
		result[i-emptyInc] = u.table[i]
	}
	return result
}

func (u *UltraTable) get(iterator ItemIterator) []interface{} {
	var result []interface{}
	for i := 0; i < len(u.table); i++ {
		if u.table[i] == nil {
			continue
		}
		if iterator(u.table[i]) {
			result = append(result, u.table[i])
		}
	}
	return result
}

func (u *UltraTable) getWithIdxIntersection(conditions map[string]interface{}) ([]interface{}, error) {
	intersectionList := []*BitMap{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		sliceList, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		intersectionList = append(intersectionList, sliceList)
		if len(intersectionList) > 1 {
			if int(sliceList.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(sliceList.Length())
			}
		} else {
			minLen = int(sliceList.Length())
			minLenIndex = 0
		}
	}

	if len(intersectionList) == 0 {
		return nil, RecordNotFound
	}

	var result []interface{}
	tempMap := map[uint32]uint64{}

	intersectionList[minLenIndex].Iterator(func(k uint32) {
		tempMap[k] = 1
		for i := 0; i < len(intersectionList); i++ {
			if i == minLenIndex {
				continue
			}

			if ok := intersectionList[i].IsExist(k); ok {
				tempMap[k] += 1
			}
		}
	})

	for k, v := range tempMap {
		if v == uint64(len(intersectionList)) {
			result = append(result, u.table[k])
		}
	}
	return result, nil
}

func (u *UltraTable) getWithIdxAggregate(conditions map[string]interface{}) ([]interface{}, error) {
	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			continue
		}
		sliceList, ok := index[vKey]
		if !ok {
			continue
		}
		aggregateList = append(aggregateList, sliceList)
	}
	if len(aggregateList) == 0 {
		return nil, RecordNotFound
	}

	tempMap := map[uint32]uint8{}

	for _, aggregateSlice := range aggregateList {
		aggregateSlice.Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}
	result := make([]interface{}, len(tempMap))
	count := 0
	for k := range tempMap {
		result[count] = u.table[k]
		count++
	}
	return result, nil
}

func (u *UltraTable) saveWithIdxIntersection(conditions map[string]interface{}, newDest interface{}) uint64 {

	intersectionList := []*BitMap{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			intersectionList = []*BitMap{}
			break
		}
		sliceList, ok := index[vKey]
		if !ok {
			intersectionList = []*BitMap{}
			break
		}
		intersectionList = append(intersectionList, sliceList)
		if len(intersectionList) > 1 {
			if int(sliceList.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(sliceList.Length())
			}
		} else {
			minLen = int(sliceList.Length())
			minLenIndex = 0
		}
	}

	tempMap := map[uint32]uint64{}
	if len(intersectionList) >= minLenIndex+1 && intersectionList[minLenIndex] != nil {
		intersectionList[minLenIndex].Iterator(func(k uint32) {
			tempMap[k] = 1
			for i := 0; i < len(intersectionList); i++ {
				if i == minLenIndex {
					continue
				}
				if ok := intersectionList[i].IsExist(k); ok {
					tempMap[k] += 1
				}
			}
		})
	}
	isReplace := false
	count := 0
	for k, v := range tempMap {
		if v == uint64(len(intersectionList)) {
			u.removeIndex(uint32(k), u.table[k])
			u.table[k] = newDest
			u.addIndex(newDest, k)
			isReplace = true
			count++
		}
	}
	if !isReplace {
		if u.emptyMap.Length() == 0 {
			err := u.addIndex(newDest, uint32(len(u.table)))
			if err != nil {
				return 0
			}
			u.table = append(u.table, newDest)
		} else {
			i := u.emptyMap.Min()
			err := u.addIndex(newDest, i)
			if err != nil {
				return 0
			}
			u.table[i] = newDest
			u.emptyMap.Remove(i)
		}
		return 1
	} else {
		return uint64(count)
	}
}

func (u *UltraTable) saveWithIdxAggregate(conditions map[string]interface{}, newDest interface{}) uint64 {

	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {
		index, ok := u.indexGroup.indexItems[idxKey]
		if !ok {
			continue
		}
		sliceList, ok := index[vKey]
		if !ok {
			continue
		}
		aggregateList = append(aggregateList, sliceList)
	}
	if len(aggregateList) == 0 {
		if u.emptyMap.Length() == 0 {
			err := u.addIndex(newDest, uint32(len(u.table)))
			if err != nil {
				return 0
			}
			u.table = append(u.table, newDest)
		} else {
			i := u.emptyMap.Min()
			err := u.addIndex(newDest, i)
			if err != nil {
				return 0
			}
			u.table[i] = newDest
			u.emptyMap.Remove(i)
		}
		return 1
	}

	tempMap := map[uint32]uint64{}

	for _, aggregateSlice := range aggregateList {
		aggregateSlice.Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}
	for k := range tempMap {
		u.removeIndex(uint32(k), u.table[k])
		u.table[k] = newDest
		u.addIndex(newDest, k)
	}
	return uint64(len(tempMap))
}
