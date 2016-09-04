package simple_db

import "fmt"

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
	}else {
		return 0
	}
}

func (sc *StateCache) Get(key string) string{
	v, ok :=sc.keyToVal.Get(key)
	if ok{
		return v.(string)
	}else{
		return "NULL"
	}
}

func (sc *StateCache) Set(key string, val string, trans *Transaction) {
	sc.keyToVal.Set(key,val)
	sc.keyToTrans.Set(key,trans)
}

func (sc *StateCache) Rollback(trans *Transaction) {
	parent := trans.parent
	keys, ok := sc.keyToTrans.GetKeysFromValue(trans)
	fmt.Printf("found the fllowing keys %#v", keys)
	if ok {
		for key, _ := range keys {
			keyStr := key.(string)
			nextVal, nextTran := parent.get(keyStr)
			sc.Set(keyStr,nextVal,nextTran)
		}
	}
}
