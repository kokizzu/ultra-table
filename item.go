package ultra_table

type item[T IRow] struct {
	isDelete bool
	buf      []byte
}



func newItem[T IRow](t T) (*item[T], error) {
	buf, err := t.Marshal()
	if err != nil {
		return nil, err
	}
	return &item[T]{
		isDelete: false,
		buf:      buf,
	}, nil
}

func (item *item[T]) Delete() {
	item.isDelete = true
	item.buf = nil
}

func (item *item[T]) IsDeleted() bool {
	return item.isDelete
}

func (item *item[T]) GetItemValue(deepCp IDeepCp[T]) T {
	t := deepCp.DeepCp()
	t.Unmarshal(item.buf)
	return t
}
