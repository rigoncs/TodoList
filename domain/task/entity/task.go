package entity

import (
	"errors"
	_const "github.com/rigoncs/TodoList/const"
	"time"
)

type Task struct {
	Id        uint      `json:"id"`
	Uid       uint      `json:"uid"`
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Status    int       `json:"status"`
	Content   string    `json:"content"`
	StartTime int64     `json:"start_time"`
	EndTime   int64     `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(uid uint, userName, title, content string) (*Task, error) {
	if uid == 0 {
		return nil, errors.New("owner ID cannot be empty")
	}
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	now := time.Now()
	return &Task{
		Uid:       uid,
		UserName:  userName,
		Title:     title,
		Status:    _const.TaskStatusEnumInit,
		Content:   content,
		StartTime: now.Unix(),
		CreatedAt: now,
	}, nil
}

func (t *Task) Complete() error {
	now := time.Now()
	t.Status = _const.TaskStatusEnumFinished
	t.UpdatedAt = now
	t.EndTime = now.Unix()
	return nil
}

func (t *Task) AddUserInfo(uid uint, userName string) {
	t.Uid = uid
	t.UserName = userName
}

func (t *Task) BelongsToUser(userId uint) bool {
	return t.Uid == userId
}

func (t *Task) UpdateContent(title, content string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	t.Title = title
	t.Content = content
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) IsExist() bool {
	return t.Id > 0
}
