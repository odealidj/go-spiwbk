package abstraction

type Pagination struct {
	Page     int    `query:"page" json:"page" validate:"required"`
	PageSize int    `query:"page_size" json:"page_size" validate:"required"`
	SortBy   string `query:"sort_by" json:"sort_by" validate:"required"`
	Sort     string `query:"sort" json:"sort" validate:"required"`
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
}
