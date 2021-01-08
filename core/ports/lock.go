package ports

type Lock interface {
	Acquire(id string) error
	Check(id string) (bool,error)
	Release(id string) error
}