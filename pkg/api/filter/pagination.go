package filter

type Pagination struct {
	Limit  int `form:"limit" binding:"omitempty,min=1"`
	Offset int `form:"offset" binding:"omitempty,min=1"`
}
