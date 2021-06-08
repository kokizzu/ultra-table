package ultra_table

import (
	"errors"

	"reflect"
	"sync"
)

var (
	RecordNotFound = errors.New("record not found")
)

type uIndex struct {
	uIndexList map[string]map[interface{}]map[uint64]uint8
}

type ItemIterator func(interface{}) bool

type UltraTable struct {
	mu            sync.RWMutex
	internalSlice []interface{}
	uIndex        uIndex
	emptyMap      map[uint64]uint8
}

func NewUltraTable() *UltraTable {
	return &UltraTable{
		internalSlice: make([]interface{}, 0),
		uIndex:        uIndex{uIndexList: map[string]map[interface{}]map[uint64]uint8{}},
		emptyMap:      map[uint64]uint8{},
	}
}

func (u *UltraTable) Add(dest interface{}) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if len(u.emptyMap) == 0 {
		u.internalSlice = append(u.internalSlice, dest)
		u.addIndex(dest, uint64(len(u.internalSlice)-1))
	} else {
		for i := range u.emptyMap {
			u.addIndex(dest, i)
			u.internalSlice[i] = dest
			delete(u.emptyMap, i)
			return
		}
	}
}

//Get benchmark performance near O(n), it is recommended to use GetWithIdx
func (u *UltraTable) Remove(iterator ItemIterator) uint64 {
	u.mu.Lock()
	defer u.mu.Unlock()
	var count uint64
	for i := 0; i < len(u.internalSlice); i++ {
		if u.internalSlice[i] == nil {
			continue
		}
		if iterator(u.internalSlice[i]) {
			u.removeIndex(uint64(i), u.internalSlice[i])
			u.internalSlice[i] = nil
			u.emptyMap[uint64(i)] = 0
			count++
		}
	}
	return count
}

//RemoveWithIdx benchmark performance near O(1)
func (u *UltraTable) RemoveWithIdx(idxKey string, vKey interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	index, ok := u.uIndex.uIndexList[idxKey]
	if !ok {
		return 0
	}
	_, ok = index[vKey]
	if !ok {
		return 0
	}
	count := uint64(0)
	for k := range index[vKey] {
		u.removeIndex(k, u.internalSlice[k])
		u.internalSlice[k] = nil
		u.emptyMap[k] = 0
		count++
	}
	return count
}

func (u *UltraTable) UpdateWithIdx(idxKey string, vKey interface{}, newDest interface{}) uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	index, ok := u.uIndex.uIndexList[idxKey]
	if !ok {
		return 0
	}
	_, ok = index[vKey]
	if !ok {
		return 0
	}
	count := 0
	for k := range index[vKey] {
		u.removeIndex(k, u.internalSlice[k])
		u.internalSlice[k] = nil
		u.emptyMap[k] = 0
		count++
	}
	for i := 0; i <= count-1; i++ {
		for j := range u.emptyMap {
			u.internalSlice[j] = newDest
			u.addIndex(newDest, j)
			delete(u.emptyMap, j)
			break
		}
	}
	return uint64(count)
}

//Get benchmark performance near O(1)
func (u *UltraTable) GetWithIdx(idxKey string, vKey interface{}) ([]interface{}, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	index, ok := u.uIndex.uIndexList[idxKey]
	if !ok {
		return nil, RecordNotFound
	}
	sliceList, ok := index[vKey]
	if !ok {
		return nil, RecordNotFound
	}
	var result []interface{}
	for k := range sliceList {
		result = append(result, u.internalSlice[k])
	}
	return result, nil
}

//GetWithIdxIntersection like where a=? and b=?
func (u *UltraTable) GetWithIdxIntersection(conditions map[string]interface{}) ([]interface{}, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	intersectionList := []map[uint64]uint8{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		index, ok := u.uIndex.uIndexList[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		sliceList, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		intersectionList = append(intersectionList, sliceList)
		if len(intersectionList) > 1 {
			if len(sliceList) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = len(sliceList)
			}
		} else {
			minLen = len(sliceList)
			minLenIndex = 0
		}
	}

	if len(intersectionList) == 0 {
		return nil, RecordNotFound
	}

	var result []interface{}
	tempMap := map[uint64]uint64{}

	for k := range intersectionList[minLenIndex] {
		tempMap[k] = 1
		for i := 0; i < len(intersectionList); i++ {
			if i == minLenIndex {
				continue
			}
			if _, ok := intersectionList[i][k]; ok {
				tempMap[k] = tempMap[k] + 1
			}
		}
	}

	for k, v := range tempMap {
		if v == uint64(len(intersectionList)) {
			result = append(result, u.internalSlice[k])
		}
	}
	return result, nil
}

//Get benchmark performance near O(n), it is recommended to use GetWithIdx
func (u *UltraTable) Get(iterator ItemIterator) []interface{} {
	u.mu.RLock()
	defer u.mu.RUnlock()
	var result []interface{}
	for i := 0; i < len(u.internalSlice); i++ {
		if u.internalSlice[i] == nil {
			continue
		}
		if iterator(u.internalSlice[i]) {
			result = append(result, u.internalSlice[i])
		}
	}
	return result
}

func (u *UltraTable) GetAll() []interface{} {
	u.mu.RLock()
	defer u.mu.RUnlock()

	result := make([]interface{}, u.Len())
	for i := 0; i < len(u.internalSlice); i++ {
		if u.internalSlice[i] == nil {
			continue
		}
		result[i] = u.internalSlice[i]
	}
	return result
}

func (u *UltraTable) Clear() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.internalSlice = make([]interface{}, 0)
	u.uIndex = uIndex{uIndexList: map[string]map[interface{}]map[uint64]uint8{}}
	u.emptyMap = make(map[uint64]uint8, 0)
}

func (u *UltraTable) Len() uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return uint64(len(u.internalSlice) - len(u.emptyMap))
}

func (u *UltraTable) Cap() uint64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return uint64(len(u.internalSlice))
}

func (u *UltraTable) Has(iterator ItemIterator) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	for i := 0; i < len(u.internalSlice); i++ {
		if u.internalSlice[i] == nil {
			continue
		}
		if iterator(u.internalSlice[i]) {
			return true
		}
	}
	return false
}

func (u *UltraTable) HasWithIdx(idxKey string, vKey interface{}) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	index, ok := u.uIndex.uIndexList[idxKey]
	if !ok {
		return false
	}
	_, ok = index[vKey]
	return ok
}

func (u *UltraTable) removeIndex(idx uint64, dest interface{}) {
	value := reflect.ValueOf(dest)
	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get("index")
		if tag == "" || tag == "-" {
			continue
		}
		m, ok := u.uIndex.uIndexList[tag]
		if ok {
			delete(m[value.Field(i).Interface()], idx)
		}
	}
}

func (u *UltraTable) addIndex(dest interface{}, idx uint64) {
	value := reflect.ValueOf(dest)
	for i := 0; i < value.NumField(); i++ {
		tag, ok := value.Type().Field(i).Tag.Lookup("index")
		if !ok {
			continue
		}
		if tag == "" || tag == "-" {
			continue
		}
		m, ok := u.uIndex.uIndexList[tag]
		if !ok {
			u.uIndex.uIndexList[tag] = make(map[interface{}]map[uint64]uint8)
			u.uIndex.uIndexList[tag][value.Field(i).Interface()] = map[uint64]uint8{idx: 0}
		} else {
			if m[value.Field(i).Interface()] == nil {
				m[value.Field(i).Interface()] = map[uint64]uint8{idx: 0}
			} else {
				m[value.Field(i).Interface()][idx] = 0
			}
		}
	}
}
