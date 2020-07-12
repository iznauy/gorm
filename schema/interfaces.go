package schema

import "github.com/iznauy/gorm/clause"

type GormDataTypeInterface interface {
	GormDataType() string
}

type CreateClausesInterface interface {
	CreateClauses() []clause.Interface
}

type QueryClausesInterface interface {
	QueryClauses() []clause.Interface
}

type UpdateClausesInterface interface {
	UpdateClauses() []clause.Interface
}

type DeleteClausesInterface interface {
	DeleteClauses() []clause.Interface
}
