package main

import (
	"github.com/wizjin/weixin"
	"net/http"
	"strings"
)

//回复jock，两种情况，回复文本或者回复文本加图片
func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	jock := GetJock()
	if jock.Type == "content" {
		w.ReplyText(strings.TrimSpace(jock.Content))
	} else {
		w.ReplyText("好像有一点错误")
	}
}

// 关注事件的处理函数
func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}

func main() {
	client := weixin.New("66523202", "wx954b9b3b03a90e6c", "05189072910ac5127f823169110790ae")
	client.HandleFunc(weixin.MsgTypeText, Echo)
	client.HandleFunc(weixin.MsgTypeShortVideo, Echo)
	client.HandleFunc(weixin.MsgTypeVoice, Echo)
	client.HandleFunc(weixin.MediaTypeImage, Echo)
	client.HandleFunc(weixin.EventSubscribe, Subscribe)
	http.Handle("/", client)
	http.ListenAndServe(":80", nil)
}
