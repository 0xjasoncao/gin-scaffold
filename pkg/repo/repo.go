package BasicRepo

import (
	"context"

	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
}

type BasicRepo[T ITable] struct {
	Db *gorm.DB
}

func NewBasicRepo[T ITable](db *gorm.DB) BasicRepo[T] {
	return BasicRepo[T]{Db: db}
}

func (r *BasicRepo[T]) Model(ctx context.Context) *gorm.DB {
	return r.Db.WithContext(ctx).Model(new(T))
}

// FindById 根据主键查询单条记录
func (r *BasicRepo[T]) FindById(ctx context.Context, id any) (*T, error) {

	var item T
	err := r.Model(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// FindByIds 根据主键查询多条记录
func (r *BasicRepo[T]) FindByIds(ctx context.Context, ids []any) ([]*T, error) {

	var items []*T
	err := r.Db.WithContext(ctx).Find(&items, ids).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

// FindAll 查询多条数据
func (r *BasicRepo[T]) FindAll(ctx context.Context, arg ...func(*gorm.DB)) ([]*T, error) {

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

// FindByConditions 根据条件查询一条数据
func (r *BasicRepo[T]) FindByConditions(ctx context.Context, fn func(tx *gorm.DB) *gorm.DB) (*T, error) {
	var item *T
	if err := fn(r.Db.WithContext(ctx)).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

// FindByWhere 根据条件查询一条数据
func (r *BasicRepo[T]) FindByWhere(ctx context.Context, where string, args ...any) (*T, error) {

	var item *T
	err := r.Model(ctx).Debug().Where(where, args...).First(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

// FindCount 根据条件统计数据总数
func (r *BasicRepo[T]) FindCount(ctx context.Context, where string, args ...any) (int64, error) {

	var count int64
	err := r.Model(ctx).Where(where, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// IsExist 根据条件查询数据是否存在
func (r *BasicRepo[T]) IsExist(ctx context.Context, where string, args ...any) (bool, error) {

	var count int64
	err := r.Model(ctx).Select("1").Where(where, args...).Limit(1).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// UpdateById 根据主键ID更新
func (r *BasicRepo[T]) UpdateById(ctx context.Context, id any, data map[string]any) (int64, error) {
	res := r.Model(ctx).Where("id = ?", id).Updates(data)
	return res.RowsAffected, res.Error
}

// UpdateByWhere 批量更新
func (r *BasicRepo[T]) UpdateByWhere(ctx context.Context, data any, where string, args ...any) (int64, error) {
	res := r.Model(ctx).Where(where, args...).Updates(data)
	return res.RowsAffected, res.Error
}

// Create 创建数据
func (r *BasicRepo[T]) Create(ctx context.Context, data *T) error {
	return r.Db.WithContext(ctx).Create(data).Error
}

// Insert 批量创建
func (r *BasicRepo[T]) Insert(ctx context.Context, data []*T) error {
	return r.Db.WithContext(ctx).Create(data).Error
}

// Txx 事物闭包函数
func (r *BasicRepo[T]) Txx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.Db.WithContext(ctx).Transaction(fn)
}

// Delete 删除数据
func (r *BasicRepo[T]) Delete(ctx context.Context, id any) error {
	return r.Db.WithContext(ctx).Delete(new(T), id).Error
}
