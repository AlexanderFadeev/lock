package lock

type LockedValue interface {
	Get() interface{}
	Set(value interface{})

	Unlock()
}

type lockedValue struct {
	*value
}

func (lv *lockedValue) Unlock() {
	lv.mutex.Unlock()
}
