package model

var models []interface {
}

func Register(m interface{}) {
	models = append(models, m)
}

func Models() []interface{} {
	return models
}
