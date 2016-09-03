package simple_db

type Transaction struct {
	storMap  map[string]*string
	inverMap map[string]map[string]bool
	child    *Transaction
	parent   *Transaction
}
/**
Initializes a transaction
 */
func NewTransaction() *Transaction {
	t := new(Transaction)
	t.inverMap = make(map[string]map[string]bool)
	t.storMap = make(map[string]*string)
	return t
}

/**
Spawns a child transaction
 */
func (t *Transaction) New() (*Transaction) {
	c := NewTransaction()
	c.parent = t
	t.child = c
	return c
}

/**
Sets value in transaction
name - name key to set
val - value of key to set
 */
func (t *Transaction) Set(name string, val string) {
	//remove old reference in InverMap
	t.Unset(name)

	//store new
	t.storMap[name] = &val
	keyMap, ok := t.inverMap[val]
	if ok {
		keyMap[name] = true
	}else {
		keyMap = make(map[string]bool)
		keyMap[name] = true
		t.inverMap[val] = keyMap
	}
}

/**
Discards one transaction, and returns parent transaction
parent - parent transaction
ok - true if we are in a transaction
 */
func (t *Transaction) Rollback() (parent *Transaction, ok bool) {
	if t.parent != nil {
		parent = t.parent
		ok = true
		parent.child = nil
		t.parent = nil
	}
	return
}

/**
Commits changes to parent transactions
commitedTrans - The resulting transaction that hold the committed state
ok - True if we are in a nested transaction / we actually had to do stuff
 */
func (t *Transaction) Commit() (commitedTrans *Transaction, ok bool) {
	if t.parent == nil {
		ok = false
	}else {
		t.iterateUp(func(parent *Transaction, tran *Transaction) {
			if parent == nil {
				return
			}
			for k, v := range tran.storMap {
				parent.Set(k, *v)
			}
			parent.child = nil
			commitedTrans = parent
		})
		ok = true
	}
	return
}

/**
Unsets value in transaction
name - key that should be unset
 */
func (t *Transaction) Unset(name string) {
	oldVal, ok := t.storMap[name]
	if ok {
		keyMap, ok := t.inverMap[*oldVal]
		if ok {
			delete(keyMap, name)
		}
	}
	t.storMap[name] = nil
}

/**
Helper method that walks up the iteration tree
 */
func (t *Transaction) iterateUp(myfun func(*Transaction, *Transaction)) {
	parent := t.parent
	tran := t
	for tran != nil {
		myfun(parent, tran)
		tran = parent
		if tran != nil {
			parent = tran.parent
		}
	}
}

/**
Counts number of keys that are set to specified value in both immediate and accessor transactions
 */
func (t *Transaction) NumEqualTo(targetVal string) (count int) {
	acc := make(map[string]bool)
	t.iterateUp(func(_, tran *Transaction) {
		keyMap, ok := tran.inverMap[targetVal]
		if ok {
			for key, _ := range keyMap {
				//valid that the key is still the same in the current transaction
				realVal := t.Get(key)
				if realVal !=nil && *t.Get(key) == targetVal {
					acc[key] = true
				}
			}
		}
	})

	count = len(acc)
	return
}
/**
Receives key value
 */
func (t *Transaction) Get(name string) (*string) {
	v, ok := t.storMap[name]
	if ok || t.parent == nil {
		return v
	}else {
		return t.parent.Get(name)
	}
}
