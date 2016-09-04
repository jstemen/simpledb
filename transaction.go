package simple_db

type Transaction struct {
	stateCache *StateCache
	keyToVal   *BidirectionalMap
	child      *Transaction
	parent     *Transaction
}
/**
Initializes a transaction
 */
func NewTransaction() *Transaction {
	t := new(Transaction)
	t.stateCache = NewStateCache()
	t.keyToVal = NewBidirectionalMap()
	return t
}

/**
Spawns a child transaction
 */
func (t *Transaction) New() (*Transaction) {
	c := NewTransaction()
	c.stateCache = t.stateCache
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
	t.set(name, val)
	t.stateCache.Set(name, val, t)
}

func (t *Transaction) set(name string, val string) {
	t.keyToVal.Set(name, val)
}

/**
Discards one transaction, and returns parent transaction
parent - parent transaction
ok - true if we are in a transaction
 */
func (t *Transaction) Rollback() (parent *Transaction, ok bool) {
	if t.parent != nil {
		t.stateCache.Rollback(t)
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
			for k, v := range tran.keyToVal.keyToValue {
				parent.Set(k.(string), v.(string))
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
	t.stateCache.Set(name, "NULL", t)
	t.unset(name)
}

func (t *Transaction) unset(name string) {
	t.keyToVal.Set(name, "NULL")
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
	return t.stateCache.NumEqualTo(targetVal)
}
/**
Receives key value
 */
func (t *Transaction) Get(name string) (string) {
	return t.stateCache.Get(name)
}

func (t *Transaction) get(name string) (string, *Transaction) {
	v, ok := t.keyToVal.Get(name)
	if ok || t.parent == nil {
		strPtr, ok := v.(string)
		if ok {
			return strPtr, t
		}else {
			return "REall NULL", nil
		}
	}else {
		return t.parent.get(name)
	}

}
