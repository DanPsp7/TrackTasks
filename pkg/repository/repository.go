package repository

type People interface {
}

type Tasks interface {
}

type TaskTime interface {
}

type Repository struct {
	People
	Tasks
	TaskTime
}

func NewRepository() *Repository {
	return &Repository{}
}
