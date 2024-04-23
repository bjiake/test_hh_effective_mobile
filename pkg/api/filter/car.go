package filter

// Car represents a filter for cars
type Car struct {
	// ID is the ID of the car to filter by
	ID *int64 `form:"id" binding:"omitempty" swagger:"description:ID of the car to filter by"`

	// RegNum is the registration number of the car to filter by
	RegNum *string `form:"regNum" binding:"omitempty" swagger:"description:Registration number of the car to filter by"`

	// Mark is the brand of the car to filter by
	Mark *string `form:"mark" binding:"omitempty" swagger:"description:Brand of the car to filter by"`

	// Model is the model of the car to filter by
	Model *string `form:"model" binding:"omitempty" swagger:"description:Model of the car to filter by"`

	// Year is the year of manufacture of the car to filter by
	Year *int32 `form:"year" binding:"omitempty" swagger:"description:Year of manufacture of the car to filter by"`

	// Owner is the ID of the car's owner to filter by
	Owner *int64 `form:"owner" binding:"omitempty" swagger:"description:ID of the car's owner to filter by"`
}
