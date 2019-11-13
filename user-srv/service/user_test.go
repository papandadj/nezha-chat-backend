package service

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/papandadj/nezha-chat-backend/common"

	"github.com/papandadj/nezha-chat-backend/proto/user"
	"github.com/papandadj/nezha-chat-backend/user-srv/dao"
)

var (
	testStub    *Stub
	testService *Service
	testUser    *dao.ModelUser
)

type Stub struct {
}

func (s *Stub) UserPost(username, password string) (user *dao.ModelUser, err error) {
	user = testUser
	return
}

func (s *Stub) UserGetByUsername(username string) (user *dao.ModelUser, ok bool, err error) {
	//有记录返回ok成功
	if username == "recorded" {
		ok = true
		return
	}
	return
}

func (s *Stub) UserCheckPassword(username, password string) (user *dao.ModelUser, ok bool, err error) {
	return
}

func TestMain(m *testing.M) {
	//新建测试用户
	now := time.Now()
	testUser = new(dao.ModelUser)
	testUser = new(dao.ModelUser)
	testUser.ID = 1
	testUser.Username = "test"
	testUser.Password = "test"
	testUser.CreateAt = &now
	testUser.UpdateAt = &now

	//新建dao数据
	testStub = new(Stub)
	//新建服务
	testService = New(testStub)

	code := m.Run()
	os.Exit(code)
}

func TestPost(t *testing.T) {
	var tests = []struct {
		Username string
		password string
		Resp     *user.PostResp
	}{
		{"recorded", "888", &user.PostResp{Error: &user.Error{Code: 400, Msg: common.UserHasRegistered}}},
		{"unrecorded", "888", &user.PostResp{Error: nil}},
		{"", "888", &user.PostResp{Error: &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}}},
		{"unrecorded", "", &user.PostResp{Error: &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}}},
		{"", "", &user.PostResp{Error: &user.Error{Code: 400, Msg: common.UsernameOrPasswordIsNull}}},
	}

	for _, test := range tests {
		resp := &user.PostResp{}

		err := testService.Post(context.Background(), &user.PostReq{Username: test.Username, Password: test.password}, resp)

		if err != nil {
			t.Error("failed to connect server  ", err)
			break
		}

		expected := fmt.Sprint(resp)
		got := fmt.Sprint(test.Resp)
		t.Logf("expected=%s, got=%s", expected, got)
		if expected != got {
			t.Errorf("post failed to input %s, expected %v , got %s", test, expected, got)
		}
	}
}
