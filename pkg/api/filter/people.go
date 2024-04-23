package filter

// People represents a filter for people
type People struct {
	// ID is the ID of the people to filter by
	ID *int64 `form:"id" binding:"omitempty" swagger:"description:ID of the people to filter by"`

	// Name is the name of the people to filter by
	Name *string `form:"name" binding:"omitempty" swagger:"description:Name of the people to filter by"`

	// SurName is the surname of the people to filter by
	SurName *string `form:"surName" binding:"omitempty" swagger:"description:Surname of the people to filter by"`

	// Patronymic is the patronymic of the people to filter by
	Patronymic *string `form:"patronymic" binding:"omitempty" swagger:"description:Patronymic of the people to filter by"`
}
