package request

type NewTaskRequest struct {
	Title    string `form:"title"`
	Contents string `form:"contents"`
}
