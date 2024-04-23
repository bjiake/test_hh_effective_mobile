package filter

type Filter struct {
	ID     *int64  `form:"id" binding:"omitempty"`
	RegNum *string `form:"regNum" binding:"omitempty"`
	Mark   *string `form:"mark" binding:"omitempty"`
	Model  *string `form:"model" binding:"omitempty"`
	Year   *int32  `form:"year" binding:"omitempty"`
	Owner  *int64  `form:"owner" binding:"omitempty"`
}
