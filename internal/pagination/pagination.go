package pagination

type Pagination struct {
	Page  uint64
	Limit uint64
}

func New() *Pagination {
	return &Pagination{}
}

const LIMIT uint64 = 0

func (p *Pagination) DefaultValues() {
	if p.Limit == 0 {
		p.Limit = 10
	}
}

// ToValue returns the memory address of the string.
func ToValue[T any](s T) *T {
	return &s
}
