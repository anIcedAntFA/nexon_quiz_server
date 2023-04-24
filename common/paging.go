package common

type QueryParams struct {
	CurrentPage int    `json:"current_page" form:"current_page"`
	PageSize    int    `json:"page_size" form:"page_size"`
	TotalItems  int64  `json:"total_items" form:"-"`
	SortBy      string `json:"sort_by" form:"sort_by"`
	OrderBy     string `json:"order_by" form:"order_by"`
	Search      string `json:"search" form:"search"`
}

func (qp *QueryParams) Fulfill() {
	if qp.CurrentPage <= 0 {
		qp.CurrentPage = 1
	}

	if qp.PageSize <= 5 {
		qp.PageSize = 5
	}

	if qp.PageSize >= 50 {
		qp.PageSize = 50
	}

	if qp.SortBy == "" {
		qp.SortBy = "content"
	}

	if qp.OrderBy == "" {
		qp.OrderBy = "asc"
	}
}

type EntityPagingResult struct {
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PageSize     int `json:"page_size"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}
