package service

import (
	"fmt"
	"securityMS/pkg/repository"
	"securityMS/projectLib/cqlWrapper"
	"testing"
)

func TestInvalidCrd(t *testing.T) {
	session := cqlWrapper.CreateCassandraSession("test")
	var tests = []struct {
		email, password string
		want            bool
	}{
		{"zeineb", "zeineb", true},
		{"zeineb", "123", false},
		{"123", "zeineb", false},
		{"123", "123", false},
	}
	loginR := repository.NewloginCrd(session)
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.email, tt.password)
		t.Run(testname, func(t *testing.T) {
			res, _ := loginR.VerfiyCredentials(tt.email, tt.password)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}

}
