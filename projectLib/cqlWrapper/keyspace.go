package cqlWrapper

import (
	"fmt"

	"github.com/gocql/gocql"
)

type KeyspaceInterface interface {
	Name() string
	Session() *gocql.Session
	Create(replication string) error
	Drop() error
}

type Keyspace struct {
	name    string
	session *gocql.Session
}

//Create a keyspace takes the name of the keyspace as a parameter
func NewKeyspace(name string, s *gocql.Session) Keyspace {
	return Keyspace{
		name:    name,
		session: s,
	}
}

//Create a keyspace takes the replication as a parameter
func (ks Keyspace) Create(replication string) error {

	if ks.session == nil {
		ks.session = CreateCassandraSession("")
	}
	return ks.session.Query(fmt.Sprintf(`CREATE KEYSPACE %q WITH REPLICATION = %s`, ks.Name(), replication)).Exec()
}

//Drop a keypsace
func (ks Keyspace) Drop() error {
	return ks.session.Query(fmt.Sprintf(`DROP KEYSPACE %q`, ks.Name())).Exec()
}

//Getters
func (ks Keyspace) Name() string {
	return ks.name
}

func (ks Keyspace) Session() *gocql.Session {
	if ks.session == nil {
		ks.session = CreateCassandraSession("")
	}
	return ks.session
}
