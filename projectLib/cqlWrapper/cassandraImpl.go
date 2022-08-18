package cqlWrapper

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

type DBconnection interface {
	CreateCassandraSession(keyspaceName string) *gocql.Session
}

//create a session with the DB and takes an empty string if the user doesn't have a keyspace
func CreateCassandraSession(keyspaceName string) *gocql.Session {
	host := os.Getenv("HOST_URL")
	if host == "" {
		host = "cassandra:9042"
	}
	cluster := gocql.NewCluster(host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}
	if keyspaceName != "" {
		cluster.Keyspace = keyspaceName
	}
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln("Unable to open a session with the Cassandra database (err=" + err.Error() + ")")
	}
	log.Println("Cassandra init done")

	return session
}
