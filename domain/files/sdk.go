package files

// Repository represents a repository
type Repository interface {
	Exists(path []string) bool
}

// Service represents a file service
type Service interface {
	Init(path []string) error
	Lock(path []string) error
	Unlock(path []string) error
	Save(path []string, bytes []byte) error
}
