package simple_db

type Transaction struct {
	StorMap  map[string]*string
	InverMap map[string][]string
	Child    *Transaction
	Parent   *Transaction
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.InverMap = make(map[string][]string)
	t.StorMap = make(map[string]*string)
	return t
}

func (t *Transaction) New() (*Transaction) {
	c := NewTransaction()
	c.Parent = t
	t.Child = c
	return c
}

func (t *Transaction) Set(name string, val string) {
	t.StorMap[name] = &val
	slice, ok := t.InverMap[val]
	if ok {
		slice = append(slice, name)
	}else {
		slice = make([]string, 0)
		slice = append(slice, name)
		t.InverMap[val] = slice
	}
	t.InverMap[val] = slice
}

func (t *Transaction) Rollback() (parent *Transaction, res bool) {
	if t.Parent != nil {
		parent = t.Parent
		res = true
		parent.Child = nil
		t.Parent = nil
	}
	return
}

/**
	Returns true if it committed, or false if no transactions are active
 */
func (t *Transaction) Commit() (newTrans *Transaction, res bool) {
	if t.Parent == nil {
		res = false
	}else {
		t.iterateUp(func(parent *Transaction, tran *Transaction) {
			if parent == nil {
				return
			}
			for k, v := range tran.StorMap {
				parent.Set(k, *v)
			}
			parent.Child = nil
			newTrans = parent
		})
		res = true
	}
	return
}

func (t *Transaction) Unset(name string) {
	t.StorMap[name] = nil
}

func (t *Transaction) iterateUp(myfun func(*Transaction, *Transaction)) {
	parent := t.Parent
	tran := t
	for tran != nil {
		myfun(parent, tran)
		tran = parent
		if tran != nil {
			parent = tran.Parent
		}
	}
}

func (t *Transaction) NumEqualTo(name string) (count int) {
	acc := make(map[string]bool)
	t.iterateUp(func(_, tran *Transaction) {
		sli, ok := tran.InverMap[name]
		if ok {
			for _, e := range sli {
				acc[e] = true
			}
		}
	})

	count = len(acc)
	return
}

func (t *Transaction) Get(name string) (*string) {
	v, ok := t.StorMap[name]
	if ok || t.Parent == nil {
		return v
	}else {
		return t.Parent.Get(name)
	}
}
