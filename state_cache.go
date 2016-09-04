package simple_db

/**
Represents a "running talley" state of what the db
should look like in the most recent transaction.
Can be rolled back using links to transactions.
 */

type StateCache struct {
	keyToVal   *BidirectionalMap
	keyToTrans *BidirectionalMap
}

func NewStateCache() *StateCache {
	st := new(StateCache)
	st.keyToVal = NewBidirectionalMap()
	st.keyToTrans = NewBidirectionalMap()
	return st
}

func (sc *StateCache) NumEqualTo(targetVal string) (count int) {
	keys, ok := sc.keyToVal.GetKeysFromValue(targetVal)
	if ok {
		return len(keys)
	} else {
		return 0
	}
}

func (sc *StateCache) Get(key string) string {
	v, ok := sc.keyToVal.Get(key)
	if ok {
		return v.(string)
	} else {
		return nullStr
	}
}

func (sc *StateCache) Set(key string, val string, trans *Transaction) {
	sc.keyToVal.Set(key, val)
	sc.keyToTrans.Set(key, trans)
}

/**
Resets values that originate from transaction to previous state.
Expensive operation: O(N*T)
 */
func (sc *StateCache) Rollback(trans *Transaction) {
	parent := trans.parent
	keys, ok := sc.keyToTrans.GetKeysFromValue(trans)
	if ok {
		for key, _ := range keys {
			keyStr := key.(string)
			nextVal, nextTran := parent.get(keyStr)
			sc.Set(keyStr, nextVal, nextTran)
		}
	}
}
