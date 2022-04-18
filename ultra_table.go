package ultra_table

import (
	"fmt"
	"sync"
)

type UltraTable[T IRow] struct {
	sync.RWMutex

	emptyMap     *BitMap
	table        []*item[T]
	fieldIndexer *fieldIndexer
}

//New ultraTable with generics
func New[T IRow]() *UltraTable[T] {
	ultraTable := &UltraTable[T]{
		table:        make([]*item[T], 0),
		fieldIndexer: newFieldIndexer(),
		emptyMap:     NewBitMap(),
	}
	return ultraTable
}

func NewWithInitializeData[T IRow](ts []T) (*UltraTable[T], error) {
	ultraTable := New[T]()
	for i := 0; i < len(ts); i++ {
		err := ultraTable.Add(ts[i])
		if err != nil {
			return nil, err
		}
	}
	return ultraTable, nil
}

func (u *UltraTable[T]) Clear() {
	u.Lock()
	defer u.Unlock()
	u.emptyMap.Clear()
	u.table = make([]*item[T], 0)
	u.fieldIndexer = newFieldIndexer()
}

func (u *UltraTable[T]) Cap() uint32 {
	u.RLock()
	defer u.RUnlock()
	return uint32(len(u.table))
}

func (u *UltraTable[T]) Len() uint32 {
	u.RLock()
	defer u.RUnlock()
	return u.len()
}

func (u *UltraTable[T]) Has(f iterator[T]) bool {
	u.RLock()
	defer u.RUnlock()
	return u.has(f)
}

func (u *UltraTable[T]) HasWithIdx(idxKey string, vKey interface{}) bool {
	u.RLock()
	defer u.RUnlock()
	return u.hasWithIdx(idxKey, vKey)
}

func (u *UltraTable[T]) Add(t T) error {
	u.Lock()
	defer u.Unlock()

	return u.add(t)
}

//Get benchmark performance near O(n), it is recommended to use GetWithIdx
func (u *UltraTable[T]) Remove(f iterator[T]) int {
	u.Lock()
	defer u.Unlock()
	return u.remove(f)
}

//RemoveWithIdx benchmark performance near O(1)
func (u *UltraTable[T]) RemoveWithIdx(idxKey string, vKey interface{}) int {
	u.Lock()
	defer u.Unlock()
	return u.removeWithIdx(idxKey, vKey)
}

func (u *UltraTable[T]) GetAll() []T {
	u.RLock()
	defer u.RUnlock()
	return u.getAll()
}

func (u *UltraTable[T]) Get(f iterator[T]) []T {
	u.RLock()
	defer u.RUnlock()
	return u.get(f)
}

func (u *UltraTable[T]) GetWithIdxCount(idxKey string, vKey interface{}) uint32 {
	u.RLock()
	defer u.RUnlock()
	return u.getWithIdxCount(idxKey, vKey)
}

//GetWithIdxIntersection like where a=? and b=?
func (u *UltraTable[T]) GetWithIdxIntersection(conditions map[string]interface{}) ([]T, error) {
	u.RLock()
	defer u.RUnlock()
	return u.getWithIdxIntersection(conditions)
}

//GetWithIdxAggregate like where a=? or b=?
func (u *UltraTable[T]) GetWithIdxAggregate(conditions map[string]interface{}) ([]T, error) {
	u.RLock()
	defer u.RUnlock()
	return u.getWithIdxAggregate(conditions)
}

//Get benchmark performance near O(1)
func (u *UltraTable[T]) GetWithIdx(idxKey string, vKey interface{}) ([]T, error) {
	u.RLock()
	defer u.RUnlock()
	return u.getWithIdx(idxKey, vKey)
}

//UpdateWithIdx
//update record with unique idx
func (u *UltraTable[T]) UpdateWithUniqueIdx(idxKey string, vKey interface{}, t T) error {
	u.Lock()
	defer u.Unlock()

	return u.updateWithUniqueIdx(idxKey, vKey, t)
}

//UpdateWithNormalIdx
//update record with normal idx
func (u *UltraTable[T]) UpdateWithNormalIdx(idxKey string, vKey interface{}, t T) (uint32, error) {
	u.Lock()
	defer u.Unlock()

	return u.updateWithNormalIdx(idxKey, vKey, t)
}

//SaveWithIdx update or insert
func (u *UltraTable[T]) SaveWithUniqueIdx(idxKey string, vKey interface{}, t T) error {
	u.Lock()
	defer u.Unlock()

	return u.saveWithUniqueIdx(idxKey, vKey, t)
}

func (u *UltraTable[T]) GetWithIdxAggregateCount(conditions map[string]interface{}) uint32 {
	u.RLock()
	defer u.RUnlock()

	return u.getWithIdxAggregateCount(conditions)
}

func (u *UltraTable[T]) GetWithIdxIntersectionCount(conditions map[string]interface{}) uint32 {
	u.RLock()
	defer u.RUnlock()

	return u.getWithIdxIntersectionCount(conditions)
}

func (u *UltraTable[T]) RemoveWithIdxIntersection(conditions map[string]interface{}) int {
	u.Lock()
	defer u.Unlock()
	return u.removeWithIdxIntersection(conditions)
}

func (u *UltraTable[T]) RemoveWithIdxAggregate(conditions map[string]interface{}) int {
	u.Lock()
	defer u.Unlock()
	return u.removeWithIdxAggregate(conditions)
}

func (u *UltraTable[T]) SaveWithNormalIdxAggregate(conditions map[string]interface{}, t T) (uint32, error) {
	u.Lock()
	defer u.Unlock()
	return u.saveWithNormalIdxAggregate(conditions, t)
}

func (u *UltraTable[T]) SaveWithNormalIdxIntersection(conditions map[string]interface{}, t T) (uint32, error) {
	u.Lock()
	defer u.Unlock()
	return u.saveWithNormalIdxIntersection(conditions, t)
}

func (u *UltraTable[T]) add(t T) error {
	item, err := newItem(t)
	if err != nil {
		return err
	}

	if u.emptyMap.Length() == 0 {
		if err := u.fieldIndexer.buildIndex(t, uint32(len(u.table))); err != nil {
			return err
		}
		u.table = append(u.table, item)
		return nil
	}

	i := u.emptyMap.Min()
	if err := u.fieldIndexer.buildIndex(t, i); err != nil {
		return err
	}
	u.table[i] = item
	u.emptyMap.Remove(i)
	return nil
}

func (u *UltraTable[T]) remove(f iterator[T]) int {
	var count int
	for i := 0; i < len(u.table); i++ {

		if u.table[i].IsDeleted() {
			continue
		}

		itemValue := u.table[i].GetItemValue()

		if f(itemValue) {
			u.fieldIndexer.removeIndex(uint32(i), itemValue)
			u.table[i].Delete()
			u.emptyMap.Add(uint32(i))
			count++
		}
	}
	return count
}

func (u *UltraTable[T]) getAll() []T {
	emptyInc := 0
	result := make([]T, u.Len())
	for i := 0; i < len(u.table); i++ {
		if u.table[i].IsDeleted() {
			emptyInc++
			continue
		}
		result[i-emptyInc] = u.table[i].GetItemValue()
	}
	return result
}

func (u *UltraTable[T]) get(f iterator[T]) []T {
	var result []T
	for i := 0; i < len(u.table); i++ {
		if u.table[i].IsDeleted() {
			continue
		}
		itemValue := u.table[i].GetItemValue()
		if f(itemValue) {
			result = append(result, itemValue)
		}
	}
	return result
}

func (u *UltraTable[T]) getWithIdx(idxKey string, vKey interface{}) ([]T, error) {
	tag, err := u.fieldIndexer.getIndexTag(idxKey)
	if err != nil {
		return nil, err
	}

	if tag.CheckIsUnique() {
		index, ok := u.fieldIndexer.uniqueIndexItems[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		idx, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		return []T{u.table[idx].GetItemValue()}, nil
	}

	if tag.CheckIsNormal() {
		index, ok := u.fieldIndexer.normalIndexItems[idxKey]
		if !ok {
			return nil, RecordNotFound
		}
		list, ok := index[vKey]
		if !ok {
			return nil, RecordNotFound
		}
		result := make([]T, list.Length())
		count := 0
		list.Iterator(func(k uint32) {
			result[count] = u.table[k].GetItemValue()
			count++
		})
		return result, nil
	}
	return nil, RecordNotFound
}

func (u *UltraTable[T]) getWithIdxCount(idxKey string, vKey interface{}) uint32 {

	bitmap, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
	if err != nil {
		return 0
	}
	return uint32(bitmap.Length())
}

func (u *UltraTable[T]) getWithIdxIntersection(conditions map[string]interface{}) ([]T, error) {
	intersectionList := []*BitMap{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
		if err != nil {
			return nil, err
		}
		intersectionList = append(intersectionList, list)
		if len(intersectionList) > 1 {
			if int(list.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(list.Length())
			}
		} else {
			minLen = int(list.Length())
			minLenIndex = 0
		}
	}
	if len(intersectionList) == 0 {
		return nil, RecordNotFound
	}

	var result []T
	tempMap := map[uint32]int{}

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
		if v == int(len(intersectionList)) {
			result = append(result, u.table[k].GetItemValue())
		}
	}
	return result, nil
}

func (u *UltraTable[T]) getWithIdxIntersectionCount(conditions map[string]interface{}) uint32 {
	intersectionList := []*BitMap{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
		if err != nil {
			return 0
		}
		intersectionList = append(intersectionList, list)
		if len(intersectionList) > 1 {
			if int(list.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(list.Length())
			}
		} else {
			minLen = int(list.Length())
			minLenIndex = 0
		}
	}
	if len(intersectionList) == 0 {
		return 0
	}

	tempMap := map[uint32]int{}

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
	count := uint32(0)
	for _, v := range tempMap {
		if v == int(len(intersectionList)) {
			count++
		}
	}
	return count
}

func (u *UltraTable[T]) len() uint32 {
	return uint32(len(u.table)) - uint32(u.emptyMap.Length())
}

func (u *UltraTable[T]) getWithIdxAggregate(conditions map[string]interface{}) ([]T, error) {
	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {

		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)

		if err != nil {
			if err == RecordNotFound {
				continue
			}
			return nil, err
		}
		aggregateList = append(aggregateList, list)
	}
	if len(aggregateList) == 0 {
		return nil, RecordNotFound
	}

	tempMap := map[uint32]uint8{}
	for i := 0; i < len(aggregateList); i++ {
		aggregateList[i].Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}
	result := make([]T, len(tempMap))
	count := 0
	for k := range tempMap {
		result[count] = u.table[k].GetItemValue()
		count++
	}
	return result, nil
}

func (u *UltraTable[T]) getWithIdxAggregateCount(conditions map[string]interface{}) uint32 {
	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {

		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)

		if err != nil {
			continue
		}
		aggregateList = append(aggregateList, list)
	}
	if len(aggregateList) == 0 {
		return 0
	}

	tempMap := map[uint32]uint8{}
	for i := 0; i < len(aggregateList); i++ {
		aggregateList[i].Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}

	return uint32(len(tempMap))
}

func (u *UltraTable[T]) removeWithIdx(idxKey string, vKey interface{}) int {
	bitmap, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
	if err != nil {
		return 0
	}
	count := int(0)
	bitmap.CloneIterator(func(k uint32) {
		u.fieldIndexer.removeIndex(k, u.table[k].GetItemValue())
		u.table[k].Delete()
		u.emptyMap.Add(k)
		count++
	})
	return count
}

func (u *UltraTable[T]) hasWithIdx(idxKey string, vKey interface{}) bool {
	return u.fieldIndexer.IsExist(idxKey, vKey)
}

func (u *UltraTable[T]) has(f iterator[T]) bool {
	for i := 0; i < len(u.table); i++ {
		if u.table[i].IsDeleted() {
			continue
		}
		if f(u.table[i].GetItemValue()) {
			return true
		}
	}
	return false
}

func (u *UltraTable[T]) updateWithUniqueIdx(idxKey string, vKey interface{}, t T) error {
	tag, err := u.fieldIndexer.getIndexTag(idxKey)
	if err != nil {
		return err
	}
	if !tag.CheckIsUnique() {
		return fmt.Errorf("idx key must unique")
	}

	index, ok := u.fieldIndexer.uniqueIndexItems[idxKey]
	if !ok {
		return RecordNotFound
	}
	idx, ok := index[vKey]
	if !ok {
		return RecordNotFound
	}
	oldItem := u.table[idx].GetItemValue()

	if u.removeWithIdx(idxKey, vKey) == 1 {
		if err := u.add(t); err != nil {
			u.add(oldItem)
			return err
		}
		return nil
	} else {
		return fmt.Errorf("not found item")
	}
}

func (u *UltraTable[T]) updateWithNormalIdx(idxKey string, vKey interface{}, t T) (uint32, error) {
	if u.getWithIdxCount(idxKey, vKey) <= 0 {
		return 0, nil
	}

	if u.fieldIndexer.isExistUnique() {
		return 0, fmt.Errorf("idx must not exist unique")
	}
	count := u.removeWithIdx(idxKey, vKey)
	for i := 0; i < count; i++ {
		u.add(t)
	}
	return uint32(count), nil
}

func (u *UltraTable[T]) saveWithUniqueIdx(idxKey string, vKey interface{}, t T) error {
	err := u.updateWithUniqueIdx(idxKey, vKey, t)
	if err != nil {
		if err == RecordNotFound {
			return u.add(t)
		}
		return err
	}
	return err
}

func (u *UltraTable[T]) removeWithIdxIntersection(conditions map[string]interface{}) int {
	intersectionList := []*BitMap{}

	minLen := 0
	minLenIndex := 0

	for idxKey, vKey := range conditions {
		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
		if err != nil {
			return 0
		}
		intersectionList = append(intersectionList, list)
		if len(intersectionList) > 1 {
			if int(list.Length()) < minLen {
				minLenIndex = len(intersectionList) - 1
				minLen = int(list.Length())
			}
		} else {
			minLen = int(list.Length())
			minLenIndex = 0
		}
	}
	if len(intersectionList) == 0 {
		return 0
	}

	tempMap := map[uint32]int{}

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
	for k, v := range tempMap {
		if v == int(len(intersectionList)) {
			u.fieldIndexer.removeIndex(k, u.table[k].GetItemValue())
			u.table[k].Delete()
			u.emptyMap.Add(k)
			count++
		}
	}
	return count
}

func (u *UltraTable[T]) removeWithIdxAggregate(conditions map[string]interface{}) int {
	aggregateList := []*BitMap{}

	for idxKey, vKey := range conditions {
		list, err := u.fieldIndexer.GetIdxBitMap(idxKey, vKey)
		if err != nil {
			continue
		}
		aggregateList = append(aggregateList, list)
	}

	tempMap := map[uint32]uint8{}
	for i := 0; i < len(aggregateList); i++ {
		aggregateList[i].Iterator(func(index uint32) {
			_, ok := tempMap[index]
			if ok {
				return
			}
			tempMap[index] = 0
		})
	}

	count := 0
	for k := range tempMap {
		u.fieldIndexer.removeIndex(k, u.table[k].GetItemValue())
		u.table[k].Delete()
		u.emptyMap.Add(k)
		count++
	}
	return count
}

func (u *UltraTable[T]) saveWithNormalIdxIntersection(conditions map[string]interface{}, t T) (uint32, error) {
	if u.fieldIndexer.isExistUnique() {
		return 0, fmt.Errorf("idx must not exist unique")
	}

	if u.getWithIdxIntersectionCount(conditions) <= 0 {
		if err := u.add(t); err != nil {
			return 0, err
		}
		return 1, nil
	}
	count := u.removeWithIdxIntersection(conditions)
	for i := 0; i < count; i++ {
		u.add(t)
	}
	return uint32(count), nil
}

func (u *UltraTable[T]) saveWithNormalIdxAggregate(conditions map[string]interface{}, t T) (uint32, error) {
	if u.fieldIndexer.isExistUnique() {
		return 0, fmt.Errorf("idx must not exist unique")
	}

	if u.getWithIdxAggregateCount(conditions) <= 0 {
		if err := u.add(t); err != nil {
			return 0, err
		}
		return 1, nil
	}
	count := u.removeWithIdxAggregate(conditions)
	for i := 0; i < count; i++ {
		u.add(t)
	}
	return uint32(count), nil
}
