package lock

type lock struct {
	lock map[string]bool
}

func (l *lock) Acquire(id string) error {
	l.lock[id] = true
	return nil
}

func (l *lock) Check(id string) (bool, error) {
	return l.lock[id], nil
}

func (l *lock) Release(id string) error {
	delete(l.lock, id)
	return nil
}
