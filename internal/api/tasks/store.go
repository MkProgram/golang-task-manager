package tasks

type SQLStore struct {
}

func NewSQLStore() Store {
	return &SQLStore{}
}

var _ Store = (*SQLStore)(nil)

func (s *SQLStore) List() []TaskDTO {
	tasks := []TaskDTO{
		{Id: 1, Description: "Lol digga, is nicht echt", Done: false},
		{Id: 2, Description: "Auch nicht echt", Done: false},
	}
	return tasks
}

func (s *SQLStore) Create(description string) TaskDTO {
	return TaskDTO{Id: 1, Description: description, Done: false}
}
