package models

type Sqlcfg struct {
	base  string
	where string
}

func (s Sqlcfg) Make() string {
	if s.where == "" {
		return s.base
	}
	return s.base + " " + "WHERE" + " " + s.where
}
