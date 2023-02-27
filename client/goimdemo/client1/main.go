package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	goimmessage "websocket_client/common/goim_message"

	"github.com/gorilla/websocket"
)

var addr = "127.0.0.1:8080"

type Msg struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func main() {
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/sendweb",
	}
	token, err := getToken()
	if err != nil {
		fmt.Println("failed login", err)
		return
	}
	wmsg := &goimmessage.Message{
		FormId:   17,
		TargetId: 16,
		Type:     goimmessage.PrivateChat,
		Media:    goimmessage.Text,
		Content:  "你好,很高兴认识你!",
	}
	dialer := websocket.DefaultDialer
	dialer.Subprotocols = append(dialer.Subprotocols, token)
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dialer.Dial", err)
		return
	}
	fmt.Println("建立连接成功")
	defer conn.Close()
	err = conn.WriteJSON(wmsg)
	if err != nil {
		fmt.Println("write:", err)
		return
	}
	rmsg := &goimmessage.Message{}
	for {
		err = conn.ReadJSON(rmsg)
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("接受到%d的信息:%s\n", rmsg.FormId, rmsg.Content)
		err = conn.WriteJSON(wmsg)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
		time.Sleep(time.Second)
	}
}

func getToken() (string, error) {
	res, err := http.PostForm("http://127.0.0.1:8080/login", url.Values{"username": {"demo2"}, "password": {"123123"}})
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(res.Body)
	msg := &Msg{}
	err = json.Unmarshal(body, msg)
	return msg.Token, err
}

func jsonE(m goimmessage.Message) {
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
