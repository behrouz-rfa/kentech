package filters

type FilmBy struct {
	ID        *string
	Title     *string
	CreatorID *string
}

type FilmFilter struct {
	Title     *StringFilter
	CreatorID *StringFilter
}
