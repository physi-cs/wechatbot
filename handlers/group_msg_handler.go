package handlers

import (
	"log"
	"strings"

	"github.com/869413421/wechatbot/glm"
	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

// GroupMessageHandler 群消息处理
type GroupMessageHandler struct {
}

// handle 处理消息
func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewGroupMessageHandler 创建群消息处理器
func NewGroupMessageHandler() MessageHandlerInterface {

	return &GroupMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收群消息
	sender, err := msg.Sender()
	group := openwechat.Group{sender}
	log.Printf("Received Group %v Text Msg : %v", group.NickName, msg.Content)

	// 不是@的不处理
	if !msg.IsAt() {
		return nil
	}

	// 替换掉@文本，然后向GPT发起请求
	replaceText := "@" + sender.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	//reply, err := glm.Completions(requestText)
	reply, err := glm.Completions(group.NickName, requestText)
	if err != nil {
		if tooMuchRound, ok := err.(glm.TooMuchRound); ok {
			log.Printf("reply message: %v \n", err)
			msg.ReplyText(tooMuchRound.GetMessage())
			return err
		}
		log.Printf("glm request error: %v \n", err)
		msg.ReplyText("机器人故障，正在修复，请等待。")
		return err
	}
	if reply == "" {
		return nil
	}

	// 获取@我的用户
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		log.Printf("get sender in group error :%v \n", err)
		return err
	}

	// 回复@我的用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	atText := "@" + groupSender.NickName
	replyText := atText + " " + reply
	_, err = msg.ReplyText(replyText)
	if err != nil {
		log.Printf("response group error: %v \n", err)
	}
	return err
}
