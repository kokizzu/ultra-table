package ultra_table

type Transaction struct {
	transLogs []*transLog
}

type TransType int

const (
	INSERT TransType = iota
	DELETE
	UPDATE
)

type transLog struct {
	transType  TransType
	ultraTable *UltraTable
	arryIndex  uint64
	olds       []interface{}

	idxKey string
	vKey   interface{}
}

func (p *transLog) Rollback() {
	p.ultraTable.Lock()
	defer p.ultraTable.Unlock()
	switch p.transType {
	case INSERT:
		p.ultraTable.removeWithTransition(p.arryIndex)
	case DELETE:
		for _, v := range p.olds {
			p.ultraTable.add(v)
		}
	case UPDATE:
		p.ultraTable.removeWithIdx(p.idxKey, p.vKey)
		for _, v := range p.olds {
			p.ultraTable.add(v)
		}
	}
}

func Begin() *Transaction {
	return &Transaction{
		transLogs: make([]*transLog, 0),
	}
}

func (t *Transaction) Commit() {
	if len(t.transLogs) > 0 {
		t.transLogs = make([]*transLog, 0)
	}
}

func (t *Transaction) Rollback() {
	for i := len(t.transLogs) - 1; i >= 0; i-- {
		t.transLogs[i].Rollback()
	}
}

func (t *Transaction) Add(ultraTable *UltraTable, dest interface{}) error {
	ultraTable.Lock()
	defer ultraTable.Unlock()
	arryIndex, err := ultraTable.addWithTransition(dest)
	if err != nil {
		return err
	}
	t.transLogs = append(t.transLogs, &transLog{
		ultraTable: ultraTable,
		transType:  INSERT,
		arryIndex:  arryIndex,
	})
	return nil
}
func (t *Transaction) RemoveWithIdx(ultraTable *UltraTable, idxKey string, vKey interface{}) uint64 {
	ultraTable.Lock()
	defer ultraTable.Unlock()

	olds, err := ultraTable.getWithIdx(idxKey, vKey)
	if err != nil {
		return 0
	}

	count := ultraTable.removeWithIdx(idxKey, vKey)
	t.transLogs = append(t.transLogs, &transLog{
		ultraTable: ultraTable,
		transType:  DELETE,
		olds:       olds,
		idxKey:     idxKey,
		vKey:       vKey,
	})
	return count
}

func (t *Transaction) UpdateWithIdx(ultraTable *UltraTable, idxKey string, vKey interface{}, newDest interface{}) uint64 {
	ultraTable.Lock()
	defer ultraTable.Unlock()

	olds, err := ultraTable.getWithIdx(idxKey, vKey)
	if err != nil {
		return 0
	}

	count := ultraTable.updateWithIdx(idxKey, vKey, newDest)

	t.transLogs = append(t.transLogs, &transLog{
		ultraTable: ultraTable,
		transType:  UPDATE,
		olds:       olds,
		idxKey:     idxKey,
		vKey:       vKey,
	})
	return count
}
