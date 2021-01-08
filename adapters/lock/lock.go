package lock

import "example.com/test_task/core/ports"

type lock struct {
	lock map[string]bool
}

func New()ports.Lock{
	return &lock{map[string]bool{}}
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
