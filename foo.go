package simple_db

type Transaction struct {
	StorMap  map[string]*string
	InverMap map[string][]string
	Children []*Transaction
	Parent   *Transaction
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.InverMap = make(map[string][]string)
	t.StorMap = make(map[string]*string)
	t.Children = make([]*Transaction, 0)
	return t
}

func (t *Transaction) New() (*Transaction) {
	c := NewTransaction()
	c.Parent = t
	t.Children = append(t.Children, c)
	return c
}

func (t *Transaction) Set(name string, val string) {
	t.StorMap[name] = &val
	slice, ok := t.InverMap[val]
	if ok {
		slice = append(slice, name)
	}else {
		slice = make([]string, 1)
		slice = append(slice, name)
		t.InverMap[val] = slice
	}

}

func (t *Transaction) Unset(name string) {
	t.StorMap[name] = nil
}

func (t *Transaction) NumEqualTo(name string) (count int) {
	slice, ok := t.InverMap[name]
	if ok {
		count = len(slice)
	}else {
		count = 0
	}
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
