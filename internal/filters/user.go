package filters

type UserBy struct {
	ID       *string
	Username *string
}

type UserFilter struct {
	Username *StringFilter
}
