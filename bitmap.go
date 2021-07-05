package ultra_table

import "github.com/RoaringBitmap/roaring"

type BitMap struct {
	*roaring.Bitmap
}

func NewBitMap() *BitMap {
	return &BitMap{
		Bitmap: roaring.New(),
	}
}

func (b *BitMap) CloneIterator(f func(uint32)) {

	i := b.Bitmap.Clone().Iterator()
	for i.HasNext() {
		f(i.Next())
	}
}

func (b *BitMap) Iterator(f func(uint32)) {

	i := b.Bitmap.Iterator()
	for i.HasNext() {
		f(i.Next())
	}
}

func (b *BitMap) Min() uint32 {
	return b.Bitmap.Minimum()
}

func (b *BitMap) Add(num uint32) {
	b.Bitmap.Add(num)
}

func (b *BitMap) Clear() {
	b.Bitmap.Clear()
}

func (b *BitMap) IsExist(num uint32) bool {
	return b.Bitmap.Contains(num)
}

func (b *BitMap) Length() uint64 {
	return b.Bitmap.GetCardinality()
}

func (b *BitMap) Remove(num uint32) {
	b.Bitmap.Remove(num)
}
