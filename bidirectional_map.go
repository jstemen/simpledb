package simple_db

type BidirectionalMap struct {
	keyToValue map[interface{}]interface{}
	//Need fast way to retrieve keys to delete
	valueToKeysMap map[interface{}]map[interface{}]bool
}

func NewBidirectionalMap() (bm *BidirectionalMap) {
	bm = new(BidirectionalMap)
	bm.keyToValue = make(map[interface{}]interface{})
	bm.valueToKeysMap = make(map[interface{}]map[interface{}]bool)
	return
}

func (bdm *BidirectionalMap) Set(key interface{}, value interface{}) {
	//remove old value
	oldVal := bdm.keyToValue[key]
	keyMap, ok := bdm.valueToKeysMap[oldVal]
	if ok {
		delete(keyMap, key)
	}
	bdm.keyToValue[key] = value

	//insert new value
	_, ok = bdm.valueToKeysMap[value]
	if !ok {
		bdm.valueToKeysMap[value] = make(map[interface{}]bool)
	}
	bdm.valueToKeysMap[value][key] = true
}

func (bdm *BidirectionalMap) GetKeysFromValue(value interface{}) (keys map[interface{}]bool, ok bool) {
	keys, ok = bdm.valueToKeysMap[value]
	return
}

func (bdm *BidirectionalMap) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = bdm.keyToValue[key]
	if ok {
		return
	} else {
		return nil, false
	}
}

func (bdm *BidirectionalMap) Delete(key interface{}) {
	value, ok := bdm.keyToValue[key]
	if ok {
		delete(bdm.valueToKeysMap, value)
	}
	delete(bdm.keyToValue, key)
}

func (bdm *BidirectionalMap) DeleteFromValue(val interface{}) {
	keys, ok := bdm.valueToKeysMap[val]
	if ok {
		for ikey, _ := range keys {
			delete(bdm.keyToValue, ikey)
		}
	}
	delete(bdm.valueToKeysMap, val)
}
