package abstraction

type Pagination struct {
	Page     *int    `query:"page" json:"page"`
	PageSize *int    `query:"page_size" json:"page_size"`
	Count    int64   `json:"count"`
	SortBy   *string `query:"sort_by" json:"sort_by"`
	Sort     *string `query:"sort" json:"sort"`
}

type PaginationSortBy struct {
	SortBy *string `query:"sort_by" json:"sort_by"`
	Sort   *string `query:"sort" json:"sort"`
}

type PaginationArr struct {
	Page     *int     `query:"page" json:"page"`
	PageSize *int     `query:"page_size" json:"page_size"`
	Count    int64    `json:"count"`
	SortBy   []string `query:"sort_by" url:"sort_by" json:"sort_by"`
	Sort     []string `query:"sort" url:"sort" json:"sort"`
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	Pages       int  `json:"pages"`
	MoreRecords bool `json:"more_records"`
}

type PaginationInfoArr struct {
	*PaginationArr
	Count       int  `json:"count"`
	Pages       int  `json:"pages"`
	MoreRecords bool `json:"more_records"`
}
