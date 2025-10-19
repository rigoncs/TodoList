package types

type List[T any] struct {
	Count int64 `json:"count"`
	Items []T   `json:"items"`
}

type DetailReq struct {
	Id uint `json:"id" form:"id"`
}

type DeleteTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type UpdateTaskReq struct {
	ID      uint   `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type CreateTaskReq struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type Pagination struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type SearchTaskReq struct {
	Info string `json:"info" form:"info"`
	Pagination
}

type ListTasksReq struct {
	Pagination
}

type TaskResp struct {
	ID        uint   `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	View      uint64 `json:"view" form:"view"`
	Status    int    `json:"status" form:"status"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}
