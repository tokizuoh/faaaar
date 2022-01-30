package types

type Sqlcfg struct {
	base  string
	where string
}

func (s Sqlcfg) Query() string {
	if s.where == "" {
		return s.base
	}
	return s.base + " " + "WHERE" + " " + s.where
}
