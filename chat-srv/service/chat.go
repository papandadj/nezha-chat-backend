package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"github.com/papandadj/nezha-chat-backend/proto/chat"
)

//Post 发送消息
func (s *Service) Post(ctx context.Context, req *chat.PostReq, resp *chat.PostResp) (err error) {

	queue := sum256(req.UserId)
	data := ChatData{
		Message:  req.Message,
		Date:     strconv.FormatInt(time.Now().UnixNano(), 10),
		Receiver: req.UserId,
		Sender:   req.TokenId,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	logger.With("queue", queue).With("data", req.Message).With("userID", req.TokenId).Debugln("用户发送消息")
	err = s.Dao.Post(queue, dataBytes)

	if err != nil {
		logger.Errorln(err)
		return
	}

	return
}

//GetQueue 发送消息
func (s *Service) GetQueue(ctx context.Context, req *chat.GetQueueReq, resp *chat.GetQueueResp) (err error) {
	sum := sum256(req.TokenId)
	logger.With("tokenID", req.TokenId).With("sum", sum).Debugln("返回用户队列")
	//在rabbit里面创建队列
	err = s.Dao.CreateQueue(sum)
	if err != nil {
		logger.Errorln(err)
		return
	}
	resp.Queue = sum
	return
}

func sum256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

//ChatData 发送到rabbit里面的数据
type ChatData struct {
	Message  string `json:"message"`
	Date     string `json:"date"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}
