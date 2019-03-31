package lock

import (
	"sync"
)

type Value interface {
	Lock() LockedValue
	RLock() RLockedValue
}

type value struct {
	v     interface{}
	mutex sync.RWMutex
}

func NewValue(v interface{}) Value {
	return &value{
		v: v,
	}
}

func (v *value) Lock() LockedValue {
	v.Lock()
	return &lockedValue{v}
}

func (v *value) RLock() RLockedValue {
	v.mutex.RLock()
	return &rLockedValue{v}
}

func (v *value) Get() interface{} {
	return v.v
}

func (v *value) Set(value interface{}) {
	v.v = value
}
