package ultra_table

type item[T IRow] struct {
	isDelete bool
	buf      []byte
	t        T
}

func newItem[T IRow](t T) (*item[T], error) {
	buf, err := t.Marshal()
	if err != nil {
		return nil, err
	}
	return &item[T]{
		isDelete: false,
		buf:      buf,
		t:        t,
	}, nil
}

func (item *item[T]) Delete() {
	item.isDelete = true
	item.buf = nil
	var t T
	item.t = t
}

func (item *item[T]) IsDeleted() bool {
	return item.isDelete
}

func (item *item[T]) GetItemValue() T {
	item.t.Unmarshal(item.buf)
	return item.t
}
