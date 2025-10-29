package core

import (
	"context"
	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
}

type UniversalRepo[T ITable] interface {
	Model(ctx context.Context) *gorm.DB
	FindWithPage(ctx context.Context, param PageParam, arg ...func(*gorm.DB)) ([]*T, *Pagination, error)
	FindById(ctx context.Context, id any) (*T, error)
	FindByIds(ctx context.Context, ids []any) ([]*T, error)
	FindAll(ctx context.Context, arg ...func(*gorm.DB)) ([]*T, error)
	FindByConditions(ctx context.Context, fn func(tx *gorm.DB) *gorm.DB) (*T, error)
	FindByWhere(ctx context.Context, where string, args ...any) (*T, error)
	FindCount(ctx context.Context, where string, args ...any) (int64, error)
	IsExist(ctx context.Context, where string, args ...any) (bool, error)
	UpdateById(ctx context.Context, id any, data any) (int64, error)
	UpdateByWhere(ctx context.Context, data any, where string, args ...any) (int64, error)
	Create(ctx context.Context, data *T) error
	Insert(ctx context.Context, data []*T) error
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	Delete(ctx context.Context, id any) error
}

type Repository[T ITable] struct {
	DB *gorm.DB
}

func NewRepository[T ITable](db *gorm.DB) Repository[T] {
	return Repository[T]{DB: db}
}

func (r *Repository[T]) Model(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx).Debug().Model(new(T))
}

func (r *Repository[T]) FindById(ctx context.Context, id any) (*T, error) {
	var item T
	err := r.Model(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *Repository[T]) FindByIds(ctx context.Context, ids []any) ([]*T, error) {

	var items []*T
	err := r.DB.WithContext(ctx).Find(&items, ids).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository[T]) FindAll(ctx context.Context, arg ...func(*gorm.DB)) ([]*T, error) {

	bd := r.Model(ctx)

	for _, fn := range arg {
		fn(bd)
	}

	var items []*T
	if err := bd.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository[T]) FindByConditions(ctx context.Context, fn func(tx *gorm.DB) *gorm.DB) (*T, error) {
	var item *T
	if err := fn(r.DB.WithContext(ctx)).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repository[T]) FindByWhere(ctx context.Context, where string, args ...any) (*T, error) {

	var item *T
	err := r.Model(ctx).Where(where, args...).First(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repository[T]) FindCount(ctx context.Context, where string, args ...any) (int64, error) {

	var count int64
	err := r.Model(ctx).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository[T]) IsExist(ctx context.Context, where string, args ...any) (bool, error) {

	var count int64
	err := r.Model(ctx).Select("1").Where(where, args...).Limit(1).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *Repository[T]) UpdateById(ctx context.Context, id any, data any) (int64, error) {
	res := r.Model(ctx).Where("id = ?", id).Updates(data)
	return res.RowsAffected, res.Error
}

func (r *Repository[T]) UpdateByWhere(ctx context.Context, data any, where string, args ...any) (int64, error) {
	res := r.Model(ctx).Where(where, args...).Updates(data)
	return res.RowsAffected, res.Error
}

func (r *Repository[T]) Create(ctx context.Context, data *T) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *Repository[T]) Insert(ctx context.Context, data []*T) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *Repository[T]) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.DB.WithContext(ctx).Transaction(fn)
}

func (r *Repository[T]) Delete(ctx context.Context, id any) error {
	return r.DB.WithContext(ctx).Delete(new(T), id).Error
}

func (r *Repository[T]) FindWithPage(ctx context.Context, param PageParam, arg ...func(*gorm.DB)) ([]*T, *Pagination, error) {
	bd := r.Model(ctx)
	for _, fn := range arg {
		fn(bd)
	}
	var items []*T

	if !param.Pagination {
		// sort
		if param.OrderFields != nil && len(param.OrderFields) > 0 {
			bd.Order(param.OrderFields.Parse())
		}
		if err := bd.Scan(&items).Error; err != nil {
			return nil, nil, err
		}
		return items, &Pagination{Total: int64(len(items))}, nil
	}
	//select count
	var count int64
	bd.Count(&count)

	// sort
	if param.OrderFields != nil && len(param.OrderFields) > 0 {
		bd.Order(param.OrderFields.Parse())
	}
	// start page
	current, pageSize := param.GetCurrent(), param.GetPageSize()
	bd.Offset((current - 1) * pageSize).Limit(pageSize)
	if err := bd.Scan(&items).Error; err != nil {
		return nil, nil, err
	}
	return items, &Pagination{
		Total:    count,
		Current:  current,
		PageSize: pageSize,
	}, nil

}
