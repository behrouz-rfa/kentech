package filters

import (
	"time"
)

type TimeRange struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

type StringFilter struct {
	Eq        *string `json:"eq"`
	Contains  *string `json:"contains"`
	Neq       *string `json:"neq"`
	StartWith *string `json:"startWith"`
	EndWith   *string `json:"endWith"`
}

type IntFilter struct {
	Eq  *int
	NEq *int
	Gt  *int
	Gte *int
	Lt  *int
	Lte *int
}

type FloatFilter struct {
	Eq  *float64
	NEq *float64
	Gt  *float64
	Gte *float64
	Lt  *float64
	Lte *float64
}
