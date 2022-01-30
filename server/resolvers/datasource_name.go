package resolvers

import "fmt"

type datasourceName struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
}

func (dsn datasourceName) string() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbname, dsn.sslmode)
}
