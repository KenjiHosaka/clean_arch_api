package request

type NewTaskRequest struct {
	Title    string `form:"title" example:"New Title"`
	Contents string `form:"contents" example:"New Contents"`
}
