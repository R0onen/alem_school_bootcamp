package bootcamp

type Lock struct {
	lock bool
}

func (l *Lock) Lock() bool {
	if l.lock {
		return false // Lock is already acquired
	}
	l.lock = true
	return true // Lock acquired successfully
}

func (l *Lock) Unlock() bool {
	if !l.lock {
		return false // Lock is already released
	}
	l.lock = false
	return true // Lock released successfully
}

func (l Lock) IsLocked() bool {
	return l.lock
}

func NewLock() Lock {
	return Lock{
		lock: false,
	}
}
