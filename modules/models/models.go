package models

type (
	PaginateRep struct {
		Start string `query:"start" validate:"max=64"`
		Limit int    `query:"limit" validate:"required,min=2max=10"`
	}
	PaginateRes struct {
		Data  any           `json:"data"`
		Limit int           `json:"limit"`
		Total int64         `json:"total"`
		First FirstPaginate `json:"first"`
		Next  NextPaginate  `jsoon:"next"`
	}

	FirstPaginate struct {
		Href string `json:"href"`
	}
	NextPaginate struct {
		Start string `json:"start"`
		Href  string `json:"href"`
	}
	KafkaOffset struct {
		Offset int64 `json:"offset" bson:"offset"`
	}
)
