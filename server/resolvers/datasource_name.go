package resolvers

import "fmt"

type datasourceName struct {
	user     string
	password string
	protocol string
	dbname   string
}

func (dsn datasourceName) string() string {
	return fmt.Sprintf("%s:%s@%s/%s", dsn.user, dsn.password, dsn.protocol, dsn.dbname)
}
