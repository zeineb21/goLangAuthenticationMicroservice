package repository

import (
	"fmt"
	"securityMS/pkg/models"
	"securityMS/projectLib/cqlWrapper"

	"github.com/gocql/gocql"
)

type LoginInfo interface {
	VerfiyCredentials(email string, password string) (bool, int)
}

type loginCrd struct {
	authorizedEmail    string
	authorizedPassword string
	Session            *gocql.Session
}

func NewloginCrd(session *gocql.Session) LoginInfo {
	return &loginCrd{
		authorizedEmail:    "",
		authorizedPassword: "",
		Session:            session,
	}
}

func (l *loginCrd) verfiyCredentials(email string, password string, ch chan<- *gocql.Iter) *loginCrd {
	defer close(ch)
	query := fmt.Sprintf(`SELECT email, password, tenantid  FROM user where email= '%v' and password = '%v'`, email, password)
	ch <- cqlWrapper.ExecQuery(query, l.Session)
	if ch != nil {
		return &loginCrd{
			authorizedEmail:    email,
			authorizedPassword: password,
			Session:            cqlWrapper.CreateCassandraSession("sec"),
		}
	} else {
		return nil
	}
}

func (l *loginCrd) VerfiyCredentials(email string, password string) (bool, int) {
	myChannel := make(chan *gocql.Iter, 1)
	go l.verfiyCredentials(email, password, myChannel)
	res := <-myChannel
	m := map[string]interface{}{}
	var crd models.User
	var tenant int
	for res.MapScan(m) {
		crd = models.User{Email: m["email"].(string), Password: m["password"].(string)}
		tenant = int(m["tenantid"].(int))
	}

	return email == crd.Email && password == crd.Password, tenant
}
