package service

import (
	"fmt"
	"testing"
)

var (
	secrete = "secrete"
)

func TestCreateToken(t *testing.T) {
	var tests = []struct {
		Arg1  TokenClaim
		token string
	}{
		{TokenClaim{Username: "nezha", ID: "1", Timestamp: 12345}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjEyMzQ1LCJ1c2VybmFtZSI6Im5lemhhIn0.xHGOIRzIylTLWx-ceTa6UMsw4uO-kQk4asfZoT0XKms"},
		{TokenClaim{Username: "panda", ID: "2", Timestamp: 0}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIiLCJ0aW1lc3RhbXAiOjAsInVzZXJuYW1lIjoicGFuZGEifQ.ODaFlPsegh2nguRZtJrXWRXJCZutCKTuKBH6DFSltdg"},
	}

	for _, test := range tests {
		token, err := CreateToken(test.Arg1, secrete)

		if err != nil {
			t.Errorf("CreateToken failed : %s", err.Error())
			continue
		}
		if test.token != token {
			t.Errorf("CreateToken failed to input %v, expected %s , got %s", test.Arg1, test.token, token)
		}
	}
}

func TestParseToken(t *testing.T) {
	var tests = []struct {
		token   string
		secrete string
		resp    TokenClaim
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjEyMzQ1LCJ1c2VybmFtZSI6Im5lemhhIn0.xHGOIRzIylTLWx-ceTa6UMsw4uO-kQk4asfZoT0XKms", secrete, TokenClaim{Username: "nezha", ID: "1", Timestamp: 12345}},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIiLCJ0aW1lc3RhbXAiOjAsInVzZXJuYW1lIjoicGFuZGEifQ.ODaFlPsegh2nguRZtJrXWRXJCZutCKTuKBH6DFSltdg", secrete, TokenClaim{Username: "panda", ID: "2", Timestamp: 0}},
	}

	for _, test := range tests {
		resp, err := ParseToken(test.token, test.secrete)

		if err != nil {
			t.Errorf("ParseToken failed : %s", err.Error())
			continue
		}

		if fmt.Sprint(resp) != fmt.Sprint(test.resp) {
			t.Errorf("CreateToken failed to input %v, expected %v , got %v", test.token, test.resp, resp)
		}
	}
}
