package factory

type Storage interface {
	Save(data string) error
}

type S3Storage struct{}

func (s *S3Storage) Save(data string) error { return nil }

type LocalStorage struct{}

func (l *LocalStorage) Save(data string) error { return nil }

func NewStorage(env string) Storage {
	if env == "prod" {
		return &S3Storage{}
	}
	return &LocalStorage{}
}
