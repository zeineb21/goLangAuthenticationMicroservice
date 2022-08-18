package cqlWrapper

import "github.com/gocql/gocql"

//Execute a query : takes the whole query as a parameter and returns an iter
func ExecQuery(query string, s *gocql.Session) *gocql.Iter {
	return s.Query(query).Consistency(gocql.One).Iter()
}
