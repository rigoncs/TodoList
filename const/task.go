package _const

const (
	DefaultPageSizeMax = 100
	DefaultPageSize    = 10
	DefaultPage        = 1
)

const (
	TaskStatusEnumInit = iota
	TaskStatusEnumFinished
)

const (
	TaskStatusInit     = "未完成"
	TaskStatusFinished = "已完成"
)

var (
	TaskStatusMap = map[int]string{
		TaskStatusEnumInit:     TaskStatusInit,
		TaskStatusEnumFinished: TaskStatusFinished,
	}
)
