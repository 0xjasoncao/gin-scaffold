package core

import (
	"fmt"
	"slices"
	"strings"
)

type PageParam struct {
	Pagination  bool        `form:"-,default=true"`                        // Pagination
	Current     int         `form:"current,default=1"`                     // Current page
	PageSize    int         `form:"pageSize,default=20" binding:"max=100"` // Page size
	OrderFields OrderFields `form:"order_fields" binding:"max=5"`
}

// Pagination 分页元数据
type Pagination struct {
	Total    int64 `json:"total"`    // Total count
	Current  int   `json:"current"`  // Current Page
	PageSize int   `json:"pageSize"` // Page Size
}

// GetCurrent less than 1 return 1
func (a PageParam) GetCurrent() int {
	current := a.Current
	if current < 1 {
		return 1
	}
	return current
}

// GetPageSize less than or equal to 0 return 20
func (a PageParam) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize <= 0 {
		pageSize = 20
	}
	return pageSize
}

type OrderDirection int

const (
	OrderByASC  OrderDirection = 1
	OrderByDESC OrderDirection = 2
)

type OrderField struct {
	Key       string         `json:"key"`
	Direction OrderDirection `json:"order"`
}

func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

type OrderFields []*OrderField

func (fs OrderFields) AddIdSortField() OrderFields {
	return append(fs, NewOrderField("id", OrderByDESC))
}

func (fs OrderFields) Parse() string {
	var orders []string
	keys := make([]string, len(fs))

	for i, item := range fs {
		key := item.Key
		direction := "ASC"
		if item.Direction == OrderByDESC {
			direction = "DESC"
		}

		if slices.Contains(keys, key) {
			continue
		}
		keys[i] = key
		orders = append(orders, fmt.Sprintf("%s %s", key, direction))
	}
	return strings.Join(orders, ",")
}
