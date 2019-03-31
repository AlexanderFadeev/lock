package lock

type RLockedValue interface {
	Get() interface{}

	Unlock()
}

type rLockedValue struct {
	*value
}

func (rlv *rLockedValue) Unlock() {
	rlv.mutex.RUnlock()
}
