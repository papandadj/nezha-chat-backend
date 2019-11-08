package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/papandadj/nezha-chat-backend/proto/auth"
)

var (
	stub    *Stub
	service *Service
)

type Stub struct {
}

func (s *Stub) AuthCheck(token string) (success bool, err error) { return }
func (s *Stub) AuthSaveToken(username, token string) (err error) { return }
func (s *Stub) AuthDelToken(token string) (err error)            { return }

func TestMain(m *testing.M) {
	//新建dao数据
	stub = new(Stub)
	//新建服务
	service = New(stub)

	service.tokenSecrete = secrete

	code := m.Run()
	os.Exit(code)
}

func TestGetToken(t *testing.T) {
	var tests = []struct {
		username string
		id       string
		Resp     *auth.GetTokenResp
	}{
		{"nezha", "1", &auth.GetTokenResp{Token: ""}},
	}

	for _, test := range tests {
		resp := &auth.GetTokenResp{}

		req := &auth.GetTokenReq{
			Username: test.username,
			Id:       test.id,
		}
		err := service.GetToken(context.Background(), req, resp)

		if err != nil {
			t.Error("failed to connect server  ", err)
			break
		}

		if resp.Token == "" {
			t.Errorf("post failed to input %s, expected not null string  , got '' ", test)
		}
		t.Log(resp.Token)
	}
}

func TestCheck(t *testing.T) {
	var tests = []struct {
		token string
		resp  *auth.CheckResp
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjEyMzQ1LCJ1c2VybmFtZSI6Im5lemhhIn0.xHGOIRzIylTLWx-ceTa6UMsw4uO-kQk4asfZoT0XKms",
			&auth.CheckResp{Username: "nezha", Id: "1"},
		}, {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjEyMzQ1LCJ1c2VybmFtZSI6Im5lemhhIn0.xHGOIRzIylTLWx-ceTa6UMsw4uO-kQk4asfZoT0XKmd",
			&auth.CheckResp{
				Error: &auth.Error{Code: 400, Msg: "signature is invalid"},
			},
		},
	}

	for _, test := range tests {
		resp := &auth.CheckResp{}
		req := &auth.CheckReq{Token: test.token}

		err := service.Check(context.Background(), req, resp)
		if err != nil {
			t.Error("failed to connect server  ", err)
			break
		}

		if fmt.Sprint(resp) != fmt.Sprint(test.resp) {
			t.Errorf("failed to input %s, expected %v string  , got %v ", test.token, test.resp, resp)
		}

	}
}
