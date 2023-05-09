package app

// type SortDirection string

// const (
// 	ASC  SortDirection = "ASC"
// 	DESC SortDirection = "DESC"
// )

type Query struct {
	Take          int    `form:"take,omitempty" binding:"min=0"`
	Page          int    `form:"page,omitempty" binding:"min=1"`
	SortField     string `form:"sortField,omitempty"`
	SortDirection string `form:"sortDirection,omitempty" enums:"ASC,DESC"`
	Search        string `form:"search,omitempty"`
}
