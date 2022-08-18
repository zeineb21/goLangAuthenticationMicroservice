package cqlWrapper

import (
	"fmt"

	"github.com/gocql/gocql"
)

type TableInterface interface {
	Name() string
	Session() *gocql.Session
	Create(query string) error
	Drop() error
}

type Table struct {
	name     string
	keyspace Keyspace
	session  *gocql.Session
}

//Create an object of type Table
func NewTable(name string, ks Keyspace, s *gocql.Session) Table {
	return Table{
		name:     name,
		keyspace: ks,
		session:  s,
	}
}

//Create a table
func (t Table) Create(query string) error {

	if t.session == nil {
		t.session = CreateCassandraSession("")
	}
	return t.session.Query(fmt.Sprintf(`CREATE TABLE %q ( %q )`, t.Name(), query)).Exec()
}

//Drop a table
func (t Table) Drop() error {
	if t.session == nil {
		t.session = CreateCassandraSession("")
	}
	return t.session.Query(fmt.Sprintf(`DROP TABLE %q.%q`, t.Keyspace().Name(), t.Name())).Exec()
}

//Insert a line into a table: takes the whole query as a parameter
func (t Table) Insert(query string) {
	if err := t.session.Query(query).Exec(); err != nil {
		fmt.Println("Error while inserting")
	}
}

//Update a table : takes the whole query as a parameter
func (t Table) Update(query string) {
	if err := t.session.Query(query).Exec(); err != nil {
		fmt.Println("Error while updating")
	}
}

//Getters
func (t Table) Name() string {
	return t.name
}

func (t Table) Session() *gocql.Session {
	if t.session == nil {
		t.session = CreateCassandraSession("")
	}
	return t.session
}

func (t Table) Keyspace() Keyspace {
	return t.keyspace
}
