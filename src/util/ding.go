package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

// MsgText def
type MsgText struct {
	Content string `json:"content"`
}

// MsgAt def
type MsgAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// MsgBody def
type MsgBody struct {
	MsgType string  `json:"msgtype"`
	Text    MsgText `json:"text"`
	At      MsgAt   `json:"at"`
}

// PostErrorf send error to dingding
func PostErrorf(msg string, args ...interface{}) {
	//ip, islocal := localhost()
	//if islocal {
	//	return
	//}
	ip := viper.GetString("dingding.ip")
	messages := strings.Split(fmt.Sprintf(msg, args...), "\n")
	if len(messages) > 10 {
		msg = strings.Join(messages[:3], "\n") + "\n"
		for i, t, flag := 3, 0, false; t < 3 && i < len(messages); i++ {
			if flag && strings.Contains(messages[i], "\t") {
				msg += messages[i] + "\n"
				t++
			}
			if strings.Contains(messages[i], "runtime/panic.go") {
				flag = true
			}
		}
	} else {
		msg = strings.Join(messages, "\n")
	}
	msg = fmt.Sprintf("ip: %s\n", ip) + msg

	url := viper.GetString("dingding.api_url")
	atMobiles := viper.GetStringSlice("dingding.at_phone_err")
	sendDingRobot(msg, atMobiles, url)

}

// sendDingRobot send message to dingding
func sendDingRobot(msg string, atMobiles []string, url string) {

	// msg = fmt.Sprintf(msg, args...)
	body := MsgBody{
		MsgType: "text",
		Text: MsgText{
			Content: msg,
		},
		At: MsgAt{
			AtMobiles: atMobiles,
			IsAtAll:   false,
		},
	}
	bs, err := json.Marshal(body)
	if err != nil {
		glog.Errorf("post msg to dingding error: %v", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("post msg to dingding error: %v", err)
	}
	if resp.StatusCode != 200 {
		glog.Error("post msg to dingding status code not 200, code:", resp.StatusCode)
	}
}
