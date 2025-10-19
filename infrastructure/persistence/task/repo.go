package task

import (
	"context"
	_const "github.com/rigoncs/TodoList/const"
	"github.com/rigoncs/TodoList/domain/task/entity"
	"github.com/rigoncs/TodoList/domain/task/repository"
	"github.com/rigoncs/TodoList/infrastructure/persistence/user"
	"github.com/rigoncs/TodoList/interfaces/types"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Task {
	return &RepositoryImpl{
		db: db,
	}
}

func Paginate(p types.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page <= 0 {
			p.Page = _const.DefaultPage
		}
		switch {
		case p.PageSize > _const.DefaultPageSizeMax:
			p.PageSize = _const.DefaultPageSizeMax
		case p.PageSize <= 0:
			p.PageSize = _const.DefaultPageSize
		}
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

func (r *RepositoryImpl) CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error) {
	task := Entity2PO(in)
	db := r.db.WithContext(ctx)
	err := db.Model(&Task{}).Create(task).Error
	if err != nil {
		return nil, err
	}
	var u *user.User
	err = db.Model(&user.User{}).Where("id = ?", task.Uid).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(task, u), nil
}

func (r *RepositoryImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	update := make(map[string]any)
	if task.Title != "" {
		update["title"] = task.Title
	}
	if task.Content != "" {
		update["content"] = task.Content
	}
	if task.Status != 0 {
		update["status"] = task.Status
	}
	err := r.db.WithContext(ctx).Model(&Task{}).Where("id = ? AND uid = ?", task.Id, task.Uid).Updates(&update).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) ListTaskByUid(ctx context.Context, uid uint, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	err := r.db.WithContext(ctx).Model(&Task{}).
		Joins("AS task LEFT JOIN user AS u ON task.uid = u.id").
		Where("u.id = ?", uid).Count(&count).
		Scopes(Paginate(p)).
		Select("u.id AS uid, u.user_name, task.*").Find(&tasks).Error
	if err != nil {
		return nil, count, err
	}
	return tasks, count, nil
}

func (r *RepositoryImpl) FindTaskByTid(ctx context.Context, tid, uid uint) (*entity.Task, error) {
	task := &entity.Task{}
	err := r.db.WithContext(ctx).Model(&Task{}).
		Joins("AS task LEFT JOIN user AS u ON task.uid = u.id").
		Where("task.id = ? AND u.id = ?", tid, uid).
		Select("u.id AS uid, u.user_name, task.*").Find(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *RepositoryImpl) SearchTask(ctx context.Context, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	err := r.db.WithContext(ctx).Model(&Task{}).
		Where("uid = ?", uid).
		Where("task.title LIKE ? OR task.content LIKE ?)", "%"+keyword+"%", "%"+keyword+"%").Count(&count).
		Scopes(Paginate(p)).
		Select("id, uid. title, status, content, start_time, end_time").Find(&tasks).Error
	if err != nil {
		return nil, count, err
	}
	return tasks, count, nil
}

func (r *RepositoryImpl) DeleteTask(ctx context.Context, uid, tid uint) error {
	err := r.db.WithContext(ctx).Model(&Task{}).Where("id = ? AND uid = ?", tid, uid).Delete(&Task{}).Error
	if err != nil {
		return err
	}
	return nil
}
