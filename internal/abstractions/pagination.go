package abstractions

type Pagination struct {
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
	Sort     string `query:"sort"`
}

type PaginationInfo struct {
	*Pagination
	Count int64 `json:"count"`
	Total int64 `json:"total"`
}
