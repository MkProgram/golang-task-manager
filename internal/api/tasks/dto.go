package tasks

type CreateTaskRequest struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateTaskRequest struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskDTO struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
